from .protocol_header import ProtocolHeader
from .protocol_body import ProtocolBody


class ProtocolFrame:
    def __init__(self, header: ProtocolHeader, body: ProtocolBody):
        self.header = header
        self.body = body
    
    def ProtocolFrameFromBytes(bytes):
        header = ProtocolHeader.ProtocolHeaderFromBytes(bytes[0:2])
        body = ProtocolBody.ProtocolBodyFromBytes(bytes[2:2+header.GetMessageSize()], header.GetMessageType())
        return ProtocolFrame(header, body)
    
    def serialize(self) -> bytes:
        header_bytes =self.header.serialize()
        body_bytes = self.body.serialize()
        return header_bytes + body_bytes
    
    def NewProtocolFrameError():
        return ProtocolFrame(ProtocolHeader(3, 0), ProtocolBody(None))

    def NewProtocolFrameRegisterSuccess():
        return ProtocolFrame(ProtocolHeader(2, 0), ProtocolBody(None))
    
    def GetMessageType(self):
        return self.header.GetMessageType()