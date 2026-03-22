import utils
from .utils import Bet
class ServerService:
    def __init__(self):
        pass

    def register_bet(self, bet: utils.Bet)-> Bet:
        utils.store_bets([bet])
        