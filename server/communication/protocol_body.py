
import datetime

from common.utils import Bet

import logging

logging.basicConfig(level=logging.INFO)

class ProtocolBody:
    def __init__(self, data):
        self.data = data
    
    def ProtocolBodyFromBytes(message_type,bytes):
        if message_type == 1:
            agency_id = int.from_bytes(bytes[0:1], byteorder='big')
            bet = ProtocolBody.deserialize_client_bet(bytes, agency_id)
            return ProtocolBody(bet)
        if message_type == 4:
            bets = ProtocolBody.deserialize_batch_bets(bytes)
            return ProtocolBody(bets)   
        if message_type == 5:
            agency_id = int.from_bytes(bytes[0:1], byteorder='big')
            return ProtocolBody(agency_id)
        logging.error(f"action: deserialize_protocol_body | result: failure | message_type: {message_type}")

    def deserialize_client_bet(bytes, agency_id) -> Bet:
        offset = 0
        client_bet_size = int.from_bytes(bytes[offset:offset+1], byteorder='big')
        offset += 1
        client_size = int.from_bytes(bytes[offset:offset+1], byteorder='big')
        offset += 1
        client_bytes = bytes[offset:offset+client_size]
        name, lastname, dni, birthdate = ProtocolBody.deserialize_client(client_bytes)
        offset += client_size
        number_size = int.from_bytes(bytes[offset:offset+1], byteorder='big')
        offset += 1
        number = bytes[offset:offset+number_size].decode('utf-8')
        logging.info(f"action: deserialize_protocol_body | result: success | name: {name} | lastname: {lastname} | dni: {dni} | birthdate: {birthdate} | number: {number}")
        
        return Bet(str(agency_id), name, lastname, dni, birthdate, number)

    def deserialize_client(bytes) -> tuple:
        offset =0
        name_size = int.from_bytes(bytes[offset:offset+1], byteorder='big')
        offset += 1
        name = bytes[offset:offset+name_size].decode('utf-8')
        offset += name_size
        lastname_size = int.from_bytes(bytes[offset:offset+1], byteorder='big')
        offset += 1
        lastname = bytes[offset:offset+lastname_size].decode('utf-8')
        offset += lastname_size
        dni_size = int.from_bytes(bytes[offset:offset+1], byteorder='big')
        offset += 1
        dni = bytes[offset:offset+dni_size].decode('utf-8')
        offset += dni_size
        birthdate_size = int.from_bytes(bytes[offset:offset+1], byteorder='big')
        offset += 1
        birthdate = bytes[offset:offset+birthdate_size].decode('utf-8')
        offset += birthdate_size
        return name, lastname, dni, birthdate
    
    def serialize(self, message_type: int) -> bytes:
        if message_type == 6:
            return self._serialize_winners()
        return bytes()
    
    def _serialize_winners(self) -> bytes:
        bytesToSend = len(self.data).to_bytes(1, byteorder='big')
        for bet in self.data:
            client_bet_bytes = self._serialize_client_bet(bet)
            bytesToSend += len(client_bet_bytes).to_bytes(1, byteorder='big') + client_bet_bytes
        return bytesToSend
    
    def _serialize_client_bet(self, bet: Bet) -> bytes:
        bytesToSend = self._serialize_string(bet.first_name)
        bytesToSend += self._serialize_string(bet.last_name)
        bytesToSend += self._serialize_string(bet.document)
        bytesToSend += self._serialize_string(bet.birthdate.strftime("%Y-%m-%d"))
        bytesToSend += self._serialize_string(str(bet.number))
        return bytesToSend

    def _serialize_string(self, string: str) -> bytes:
        string_bytes = string.encode('utf-8')
        return len(string_bytes).to_bytes(1, byteorder='big') + string_bytes
    
    def deserialize_batch_bets(bytes) -> list[Bet]:
        offset = 0
        agency_id = int.from_bytes(bytes[offset:offset+1], byteorder='big')
        offset += 1
        batch_size = int.from_bytes(bytes[offset:offset+1], byteorder='big')
        logging.info(f"action: deserialize_batch_size | result: success | batch_size: {batch_size}")
        offset += 1
        bets = []
        for _ in range(batch_size):
            client_bet_size = int.from_bytes(bytes[offset:offset+1], byteorder='big')
            client_bet = ProtocolBody.deserialize_client_bet(bytes[offset:offset+client_bet_size+1], agency_id)
            offset += 1
            offset += client_bet_size
            bets.append(client_bet)
        logging.info(f"action: deserialize_protocol_body | result: success | batch_size: {len(bets)}")
        return bets