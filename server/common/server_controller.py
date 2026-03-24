from .server_service import ServerService

class ServerController:
    def __init__(self, server_service: ServerService):
        self.server_service = server_service
    
    def register_bet(self, bet):
        return self.server_service.register_bet(bet)
    
    def register_batch_bets(self, bets):
        return self.server_service.register_batch_bets(bets)
    
    def ask_for_winners(self, agency_id):
        return self.server_service.ask_for_winners(agency_id)
    
    def get_winners(self, agency_id):
        return self.server_service.get_winners(agency_id)