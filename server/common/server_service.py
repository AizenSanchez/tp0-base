from .utils import Bet, store_bets, load_bets, has_won

import threading
import logging

class ServerService:
    def __init__(self, bet_file_lock: threading.Lock, agencies_barrier: threading.Barrier):
        self.bet_file_lock = bet_file_lock
        self.agencies_barrier = agencies_barrier
        pass

    def register_bet(self, bet: Bet) -> Bet:
        self.bet_file_lock.acquire()
        store_bets([bet])
        self.bet_file_lock.release()
        return bet
    
    def register_batch_bets(self, bets: list[Bet]) -> list[Bet]:
        self.bet_file_lock.acquire()
        store_bets(bets)
        self.bet_file_lock.release()
        return bets
    
    def get_winners(self, agency_id: int) -> list[Bet]:
        winners = []
        logging.info(f"action: ask_for_winners | result: waiting | agency_id: {agency_id}")
        self.agencies_barrier.wait()
        logging.info(f"action: ask_for_winners | result: proceeding | agency_id: {agency_id}")
        self.bet_file_lock.acquire()
        for bet in load_bets():
            if bet.agency == agency_id and has_won(bet):
                winners.append(bet)
        self.bet_file_lock.release()

        return winners