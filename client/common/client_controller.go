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

func (clientController ClientController) GetClientDNI() uint64 {
	clientBet := clientController.clientService.GetClientBet()
	client := clientBet.GetClient()
	return client.GetDni()
}

func (clientController ClientController) GetClientBetNumber() uint64 {
	clientBet := clientController.clientService.GetClientBet()
	return clientBet.GetNumber()
}
