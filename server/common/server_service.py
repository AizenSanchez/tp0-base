from .utils import Bet, store_bets


class ServerService:
    def __init__(self):
        pass

    def register_bet(self, bet: Bet) -> Bet:
        store_bets([bet])
        return bet
    
    def register_batch_bets(self, bets: list[Bet]) -> list[Bet]:
        store_bets(bets)
        return bets
        