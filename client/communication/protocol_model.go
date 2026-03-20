package communication

import "github.com/7574-sistemas-distribuidos/docker-compose-init/client/common"

type ProtocolHeader struct {
	messageType uint8
	messageSize uint8
}

type ProtocolBody struct {
	clientBet common.ClientBet
}

type ProtocolModel struct {
	header ProtocolHeader
	body   ProtocolBody
}

func NewProtocolModelRequest(clientBet common.ClientBet) ProtocolModel {
	return ProtocolModel{
		header: ProtocolHeader{
			messageType: 1,
			messageSize: 0,
		},
		body: ProtocolBody{
			clientBet: clientBet,
		},
	}
}
