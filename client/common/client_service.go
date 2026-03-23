package common

type ClientService struct {
}

func NewClientService() ClientService {
	return ClientService{}
}

func (clientService *ClientService) CreateClientBet(ClientConfig ClientConfig) (ClientBet, error) {
	clientBet, err := NewClientBet(ClientConfig)
	if err != nil {
		return ClientBet{}, err
	}
	return clientBet, nil
}
