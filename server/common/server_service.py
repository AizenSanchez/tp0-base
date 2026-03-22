from .utils import Bet, store_bets


class ServerService:
    def __init__(self):
        pass

    def register_bet(self, bet: Bet) -> Bet:
        store_bets([bet])
        return bet
        