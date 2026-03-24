from .protocol_frame import ProtocolFrame
from .protocol_header import ProtocolHeader
import logging
import socket
import sys
import signal

HEADER_SIZE = 3

class ServerConnection:
    def __init__(self, port, listen_backlog):
        self.socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.socket.bind(('', port))
        self.socket.listen(listen_backlog)
        self.clients_sockets = {}

    def send_to(self, protocol_frame: ProtocolFrame, client_socket: socket.socket) -> socket.socket:
        bytes_to_send = protocol_frame.serialize()
        bytes_sent = client_socket.send(bytes_to_send)
        if bytes_sent != len(bytes_to_send):
            logging.error(f"action: send_message | result: failure | reason: short write | expected_size: {len(bytes_to_send)} | sent_size: {bytes_sent}")
        return client_socket

    def accept_client(self)-> socket.socket:
        client_socket, addr = self.socket.accept()
        self.clients_sockets[addr] = client_socket
        logging.info(f"action: accept_connection | result: success | ip: {addr[0]}")
        return client_socket
    
    def receive_from(self, client_socket: socket.socket) -> tuple[ProtocolFrame, socket.socket]:
        header_bytes = client_socket.recv(HEADER_SIZE)
        if len(header_bytes) < HEADER_SIZE:
            logging.error(f"action: receive_message | result: failure | reason: short read for header | expected_size: {HEADER_SIZE} | received_size: {len(header_bytes)}")
            return ProtocolFrame.NewProtocolFrameError(), client_socket
        header = ProtocolHeader.ProtocolHeaderFromBytes(header_bytes)
        body_bytes = client_socket.recv(header.GetMessageSize())
        if len(body_bytes) < header.GetMessageSize():
            logging.error(f"action: receive_message | result: failure | reason: short read for body | expected_size: {header.GetMessageSize()} | received_size: {len(body_bytes)}")
            return ProtocolFrame.NewProtocolFrameError(), client_socket
        return ProtocolFrame.ProtocolFrameFromBytes(header_bytes + body_bytes), client_socket
    
    def close(self):
        self.socket.close()
        for client_socket in self.clients_sockets.values():
            client_socket.close()
    