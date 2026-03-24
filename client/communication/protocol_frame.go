package communication

import (
	"errors"

	"github.com/7574-sistemas-distribuidos/docker-compose-init/client/common"
)

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
func NewProtocolFrameRequestAskForWinners(agencyId int) ProtocolFrame {
	return ProtocolFrame{
		header: ProtocolHeader{
			messageType: 5,
			messageSize: 0,
		},
		body: []byte{byte(agencyId)},
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

func (protocolFrame *ProtocolFrame) GetWinners() ([]common.ClientBet, bool, error) {

	if protocolFrame.GetMessageType() == 3 {
		return []common.ClientBet{}, false, errors.New("message type indicates failure")
	} else if protocolFrame.GetMessageType() == 6 {
		result, err := protocolFrame.deserializeBodyToWinners()
		if err != nil {
			return []common.ClientBet{}, false, err
		}
		return result, true, nil
	} else if protocolFrame.GetMessageType() == 7 {
		return []common.ClientBet{}, false, nil
	} else {
		return []common.ClientBet{}, false, errors.New("unexpected message type")
	}
}

func (protocolFrame *ProtocolFrame) deserializeBodyToWinners() ([]common.ClientBet, error) {
	clientBetsNumber := protocolFrame.body[0]
	clientBets := make([]common.ClientBet, 0)
	bodyBytes := protocolFrame.body[1:]
	for i := 0; i < int(clientBetsNumber); i++ {
		clientBetSize := bodyBytes[0]
		bodyBytes = bodyBytes[1:]
		clientBetBytes := bodyBytes[:clientBetSize]
		clientBet, err := DeserializeClientBet(clientBetBytes)
		if err != nil {
			return []common.ClientBet{}, err
		}
		clientBets = append(clientBets, clientBet)
		bodyBytes = bodyBytes[clientBetSize:]
	}
	return clientBets, nil
}
