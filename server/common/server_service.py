from .utils import Bet, store_bets, load_bets, has_won


class ServerService:
    def __init__(self, total_agencies: int):
        self.total_agencies = total_agencies
        self.agencies_asked = {}
        pass

    def register_bet(self, bet: Bet) -> Bet:
        store_bets([bet])
        return bet
    
    def register_batch_bets(self, bets: list[Bet]) -> list[Bet]:
        store_bets(bets)
        return bets
    
    def ask_for_winners(self, agency_id: int) -> bool:
        self.agencies_asked[agency_id] = True
        return len(self.agencies_asked) == self.total_agencies
    
    def get_winners(self, agency_id: int) -> list[Bet]:
        winners = []
        self.agencies_asked[agency_id] = True
        if len(self.agencies_asked) == self.total_agencies:
            for bet in load_bets():
                if bet.agency == agency_id and has_won(bet):
                    winners.append(bet)
        return winners