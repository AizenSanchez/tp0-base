from .protocol_frame import ProtocolFrame
from .protocol_header import ProtocolHeader
import logging
import socket
import sys
import signal

HEADER_SIZE = 2

class ServerConnection:
    def __init__(self, port, listen_backlog):
        signal.signal(signal.SIGTERM, self._graceful_shutdown)
        self.socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.socket.bind('', port)
        self.socket.listen(listen_backlog)
        self.clients_sockets = {}

    def send_to_and_close(self, protocol_frame: ProtocolFrame, client_socket: socket.socket):
        bytes_to_send = protocol_frame.serialize()
        bytes_sent = client_socket.send(bytes_to_send)
        if bytes_sent != len(bytes_to_send):
            logging.error(f"action: send_message | result: failure | reason: short write | expected_size: {len(bytes_to_send)} | sent_size: {bytes_sent}")
        client_socket.close()

    
    def receive(self) -> tuple[ProtocolFrame, socket.socket]:
        client_socket, addr = self.socket.accept()
        self.clients_sockets[addr] = client_socket
        logging.info(f"action: accept_connection | result: success | ip: {addr[0]}")
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
    
    def _graceful_shutdown(self, signum, frame):
        self.socket.close()
        logging.info("action: close_server_socket | result: success")
        for client_addr, client_sock in self.clients_sockets.items():
            client_sock.close()                
            logging.info(f"action: close_client_socket | result: success | ip: {client_addr[0]}")
        logging.info("action: shutdown_server | result: success")
        sys.exit(0)