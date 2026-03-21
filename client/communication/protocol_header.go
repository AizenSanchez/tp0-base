package communication

import "fmt"

type ProtocolHeader struct {
	messageType uint8
	messageSize uint8
}

func (ProtocolHeader *ProtocolHeader) GetMessageType() uint8 {
	return ProtocolHeader.messageType
}

func (ProtocolHeader *ProtocolHeader) GetMessageSize() uint8 {
	return ProtocolHeader.messageSize
}

func (protocolHeader *ProtocolHeader) SetMessageSize(size uint8) {
	protocolHeader.messageSize = size
}

func (protocolHeader *ProtocolHeader) Serialize() []byte {
	headerSerialized := make([]byte, 0)
	headerSerialized = append(headerSerialized, SerializeUint8(protocolHeader.GetMessageType())...)
	headerSerialized = append(headerSerialized, SerializeUint8(protocolHeader.GetMessageSize())...)
	return headerSerialized
}

func DeserializeHeader(headerBytes []byte) (ProtocolHeader, error) {
	if len(headerBytes) != 2 {
		return ProtocolHeader{}, fmt.Errorf("invalid header size: expected 2 bytes, got %d bytes", len(headerBytes))
	}
	messageType := headerBytes[0]
	messageSize := headerBytes[1]
	return ProtocolHeader{
		messageType: messageType,
		messageSize: messageSize,
	}, nil
}
