package communication

func (ProtocolHeader *ProtocolHeader) GetMessageType() uint8 {
	return ProtocolHeader.messageType
}

func (ProtocolHeader *ProtocolHeader) GetMessageSize() uint8 {
	return ProtocolHeader.messageSize
}

func (protocolHeader *ProtocolHeader) SetMessageSize(size uint8) {
	protocolHeader.messageSize = size
}
func (ProtocolModel *ProtocolModel) GetHeader() ProtocolHeader {
	return ProtocolModel.header
}

func (ProtocolModel *ProtocolModel) GetBody() ProtocolBody {
	return ProtocolModel.body
}
func (protocolModel *ProtocolModel) Serialize() []byte {
	result := make([]byte, 0)
	bodySerialized := protocolModel.serializeBody()
	protocolModel.header.SetMessageSize(uint8(len(bodySerialized)))
	result = append(result, protocolModel.serializeHeader()...)
	result = append(result, protocolModel.serializeBody()...)
	return result
}

func (protocolModel *ProtocolModel) serializeHeader() []byte {
	return protocolModel.header.Serialize()
}

func (ProtocolHeader *ProtocolHeader) Serialize() []byte {
	headerSerialized := make([]byte, 0)
	headerSerialized = append(headerSerialized, SerializeUint8(ProtocolHeader.GetMessageType())...)
	headerSerialized = append(headerSerialized, SerializeUint8(ProtocolHeader.GetMessageSize())...)
	return headerSerialized
}

func (protocolModel *ProtocolModel) serializeBody() []byte {
	return protocolModel.body.Serialize()
}
func (protocolBody *ProtocolBody) Serialize() []byte {
	return SerializeClientBet(protocolBody.clientBet)
}
