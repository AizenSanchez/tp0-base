
import datetime

from common.utils import Bet

import logging

logging.basicConfig(level=logging.INFO)

class ProtocolBody:
    def __init__(self, bet: Bet):
        self.bet = bet
    
    def ProtocolBodyFromBytes(bytes):
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
        return ProtocolBody(Bet('0', name, lastname, dni, birthdate, number))

    def deserialize_client(bytes):
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
    
    def serialize(self) -> bytes:
        return bytes()