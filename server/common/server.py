from communication.connection import ServerConnection
from common.server_router import ServerRouter
from common.server_controller import ServerController
from common.server_service import ServerService
import threading
import signal
import logging
import sys

class Server:
    def __init__(self, port, listen_backlog, total_agencies):
        signal.signal(signal.SIGTERM, self._graceful_shutdown)
        bet_file_lock = threading.Lock()
        agencies_barrier = threading.Barrier(total_agencies)
        self.connection = ServerConnection(port, listen_backlog)
        server_service = ServerService(bet_file_lock, agencies_barrier)
        server_controller = ServerController(server_service)
        self.server_router = ServerRouter(server_controller)
        self.thread_clients: list[threading.Thread] = []

    def run(self):

        while True:
            client_socket = self.connection.accept_client()
            client_thread = threading.Thread(target=self._handle_client, args=(client_socket,))
            client_thread.start()
            self.thread_clients.append(client_thread)

    def _handle_client(self, client_socket):
        while client_socket:
            protocol_frame, client_socket = self.connection.receive_from(client_socket)
            if not protocol_frame or not client_socket:
                break
            response_frame = self.server_router.route(protocol_frame)
            if not self.connection.send_to(response_frame, client_socket):
                break

        if client_socket:
            client_socket.close()

    def _graceful_shutdown(self, signum, frame):
        self.connection.close()
        logging.info("action: close_server_socket | result: success")
        for client_addr, client_sock in self.connection.clients_sockets.items():
            client_sock.close()
            logging.info(f"action: close_client_socket | result: success | ip: {client_addr[0]}")
        logging.info("action: shutdown_server | result: success")

        for thread in self.thread_clients:
            thread.join()

        sys.exit(0)