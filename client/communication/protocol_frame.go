package communication

type ProtocolFrame struct {
	header ProtocolHeader
	body   []byte
}

func NewProtocolFrameRequestRegisterBet(bodyBytes []byte) ProtocolFrame {
	return ProtocolFrame{
		header: ProtocolHeader{
			messageType: 1,
			messageSize: 0,
		},
		body: bodyBytes,
	}
}

func NewProtocolFrameRequestRegisterBets(bodyBytes []byte) ProtocolFrame {
	return ProtocolFrame{
		header: ProtocolHeader{
			messageType: 4,
			messageSize: 0,
		},
		body: bodyBytes,
	}
}

func (protocolFrame *ProtocolFrame) GetHeader() ProtocolHeader {
	return protocolFrame.header
}

func (protocolFrame *ProtocolFrame) GetBody() []byte {
	return protocolFrame.body
}
func (protocolFrame *ProtocolFrame) Serialize() []byte {
	result := make([]byte, 0)
	bodySerialized := protocolFrame.serializeBody()
	protocolFrame.header.SetMessageSize(uint16(len(bodySerialized)))
	result = append(result, protocolFrame.serializeHeader()...)
	result = append(result, bodySerialized...)
	return result
}

func (protocolFrame *ProtocolFrame) GetMessageType() uint8 {
	return protocolFrame.header.GetMessageType()
}

func (protocolFrame *ProtocolFrame) serializeHeader() []byte {
	return protocolFrame.header.Serialize()
}

func (protocolFrame *ProtocolFrame) serializeBody() []byte {
	return protocolFrame.body
}
