package communication

import "github.com/7574-sistemas-distribuidos/docker-compose-init/client/common"

type ProtocolBody struct {
	clientBet common.ClientBet
}

func DeserializeBody(bodyBytes []byte) (ProtocolBody, error) {
	return ProtocolBody{}, nil
}

func (protocolBody *ProtocolBody) Serialize() []byte {
	return SerializeClientBet(protocolBody.clientBet)
}
