package communication

import "github.com/7574-sistemas-distribuidos/docker-compose-init/client/common"

type ProtocolFrame struct {
	header ProtocolHeader
	body   ProtocolBody
}

func NewProtocolFrameRequestRegisterBet(clientBet common.ClientBet) ProtocolFrame {
	return ProtocolFrame{
		header: ProtocolHeader{
			messageType: 1,
			messageSize: 0,
		},
		body: ProtocolBody{
			clientBet: clientBet,
		},
	}
}

func (protocolFrame *ProtocolFrame) GetHeader() ProtocolHeader {
	return protocolFrame.header
}

func (protocolFrame *ProtocolFrame) GetBody() ProtocolBody {
	return protocolFrame.body
}
func (protocolFrame *ProtocolFrame) Serialize() []byte {
	result := make([]byte, 0)
	bodySerialized := protocolFrame.serializeBody()
	protocolFrame.header.SetMessageSize(uint8(len(bodySerialized)))
	result = append(result, protocolFrame.serializeHeader()...)
	result = append(result, protocolFrame.serializeBody()...)
	return result
}

func (protocolFrame *ProtocolFrame) GetMessageType() uint8 {
	return protocolFrame.header.GetMessageType()
}

func (protocolFrame *ProtocolFrame) serializeHeader() []byte {
	return protocolFrame.header.Serialize()
}

func (protocolFrame *ProtocolFrame) serializeBody() []byte {
	return protocolFrame.body.Serialize()
}
