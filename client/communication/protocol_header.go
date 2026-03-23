package communication

import (
	"encoding/binary"
	"fmt"
)

const HEADER_SIZE = 3

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
	if len(headerBytes) != HEADER_SIZE {
		return ProtocolHeader{}, fmt.Errorf("invalid header size: expected %d bytes, got %d bytes", HEADER_SIZE, len(headerBytes))
	}
	messageType := headerBytes[0]
	messageSize := binary.BigEndian.Uint16(headerBytes[1:HEADER_SIZE])
	return ProtocolHeader{
		messageType: messageType,
		messageSize: messageSize,
	}, nil
}
