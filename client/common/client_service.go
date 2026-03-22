package common

type ClientService struct {
	clientBet ClientBet
}

func NewClientService(clientBet ClientBet) ClientService {
	return ClientService{
		clientBet: clientBet,
	}
}

func (clientService ClientService) GetClientBet() ClientBet {
	return clientService.clientBet
}

func (clientService *ClientService) CreateClientBet(ClientConfig ClientConfig) (ClientBet, error) {
	clientBet, err := NewClientBet(ClientConfig)
	if err != nil {
		return ClientBet{}, err
	}
	clientService.clientBet = clientBet
	return clientBet, nil
}
