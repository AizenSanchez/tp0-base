package communication

import (
	"encoding/binary"
	"fmt"
)

type ProtocolHeader struct {
	messageType uint8
	messageSize uint16
}

func (ProtocolHeader *ProtocolHeader) GetMessageType() uint8 {
	return ProtocolHeader.messageType
}

func (ProtocolHeader *ProtocolHeader) GetMessageSize() uint16 {
	return ProtocolHeader.messageSize
}

func (protocolHeader *ProtocolHeader) SetMessageSize(size uint16) {
	protocolHeader.messageSize = size
}

func (protocolHeader *ProtocolHeader) Serialize() []byte {
	result := make([]byte, 0)
	result = append(result, protocolHeader.messageType)
	sizeBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(sizeBytes, protocolHeader.messageSize)
	result = append(result, sizeBytes...)
	return result
}

func DeserializeHeader(headerBytes []byte) (ProtocolHeader, error) {
	if len(headerBytes) != 3 {
		return ProtocolHeader{}, fmt.Errorf("invalid header size: expected 3 bytes, got %d bytes", len(headerBytes))
	}
	messageType := headerBytes[0]
	messageSize := binary.BigEndian.Uint16(headerBytes[1:3])
	return ProtocolHeader{
		messageType: messageType,
		messageSize: messageSize,
	}, nil
}
