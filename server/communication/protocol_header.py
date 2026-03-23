
class ProtocolHeader:
    def __init__(self, message_type, message_size):
        self.message_type = message_type
        self.message_size = message_size
    
    def GetMessageType(self):
        return self.message_type
    
    def GetMessageSize(self):
        return self.message_size
    
    def ProtocolHeaderFromBytes(bytes):
        message_type = bytes[0]
        message_size = int.from_bytes(bytes[1:3], byteorder='big')
        return ProtocolHeader(message_type, message_size)
    
    def serialize(self) -> bytes:
        return bytes([self.message_type]) + self.message_size.to_bytes(2, byteorder='big')