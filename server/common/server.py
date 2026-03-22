from communication.connection import ServerConnection
from common.server_router import ServerRouter
class Server:
    def __init__(self, port, listen_backlog):
        self.connection = ServerConnection(port, listen_backlog)
        self.server_router = ServerRouter()

    def run(self):

        while True:
            protocol_frame, client_socket = self.connection.receive()
            response_frame = self.server_router.route(protocol_frame)
            self.connection.send_to_and_close(response_frame, client_socket)

