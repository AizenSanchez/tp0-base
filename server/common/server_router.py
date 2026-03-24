from .server_controller import ServerController
from communication.protocol_frame import ProtocolFrame
import logging
from .utils import Bet
class ServerRouter:
    def __init__(self, server_controller: ServerController):
        self.server_controller = server_controller
    
    def route(self, protocol_frame: ProtocolFrame) -> ProtocolFrame:
        message_type = protocol_frame.GetMessageType()
        if message_type == 1:
            bet = self.server_controller.register_bet(protocol_frame.body.data)
            if not bet:
                logging.error(f"action: apuesta_almacenada | result: failure ")
                return ProtocolFrame.NewProtocolFrameError()
            logging.info(f"action: apuesta_almacenada | result: success | dni: {bet.document} | numero: {bet.number}")
        elif message_type == 4:
            bets = self.server_controller.register_batch_bets(protocol_frame.body.data)
            if not bets:
                logging.error(f"action: apuesta_recibida | result: failure | cantudad: {len(protocol_frame.body.data)}")
                return ProtocolFrame.NewProtocolFrameError()
            logging.info(f"action: apuesta_recibida | result: success | cantidad: {len(bets)}")
        elif message_type == 5: 
            winners = self.server_controller.get_winners(protocol_frame.body.data)
            logging.info(f"action: sorteo | result: success")
            return ProtocolFrame.NewProtocolFrameAskForWinnersSuccess(winners)
        else:
            logging.error(f"action: mensaje_recibido | result: failure | message_type: {message_type}")
            return ProtocolFrame.NewProtocolFrameError()
        return ProtocolFrame.NewProtocolFrameRegisterSuccess()
        
        