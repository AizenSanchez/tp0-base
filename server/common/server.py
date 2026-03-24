from communication.connection import ServerConnection
from common.server_router import ServerRouter
from common.server_controller import ServerController
from common.server_service import ServerService
class Server:
    def __init__(self, port, listen_backlog, total_agencies):
        self.connection = ServerConnection(port, listen_backlog)
        server_service = ServerService()
        server_controller = ServerController(server_service)
        self.server_router = ServerRouter(server_controller)

    def run(self):

        while True:
            protocol_frame, client_socket = self.connection.receive()
            response_frame = self.server_router.route(protocol_frame)
            self.connection.send_to_and_close(response_frame, client_socket)

