package common

type ClientController struct {
	clientService ClientService
}

func NewClientController(clientService ClientService) ClientController {
	return ClientController{
		clientService: clientService,
	}
}

func (clientController *ClientController) CreateClientBet(ClientConfig ClientConfig) (ClientBet, error) {
	clientBet, err := clientController.clientService.CreateClientBet(ClientConfig)
	if err != nil {
		return ClientBet{}, err
	}
	return clientBet, nil
}
