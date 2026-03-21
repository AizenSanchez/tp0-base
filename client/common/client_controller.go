package common

type ClientController struct {
	clientService ClientService
}

func NewClientController(clientService ClientService) ClientController {
	return ClientController{
		clientService: clientService,
	}
}

func (clientController ClientController) GetClientBet() ClientBet {
	return clientController.clientService.GetClientBet()
}

func (clientController ClientController) GetClientDNI() string {
	clientBet := clientController.clientService.GetClientBet()
	client := clientBet.GetClient()
	return client.GetDni()
}

func (clientController ClientController) GetClientBetNumber() string {
	clientBet := clientController.clientService.GetClientBet()
	return clientBet.GetNumber()
}
