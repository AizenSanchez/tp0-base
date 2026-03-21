package common

type Client struct {
	name      string
	lastName  string
	dni       string
	birthDate string
}

func NewClient(name string, lastName string, dni string, birthDate string) (Client, error) {
	if err := ValidateDNI(dni); err != nil {
		return Client{}, err
	}

	if err := ValidateBirthDate(birthDate); err != nil {
		return Client{}, err
	}

	return Client{
		name:      name,
		lastName:  lastName,
		dni:       dni,
		birthDate: birthDate,
	}, nil
}

func (c Client) GetName() string {
	return c.name
}

func (c Client) GetLastName() string {
	return c.lastName
}

func (c Client) GetDni() string {
	return c.dni
}

func (c Client) GetBirthDate() string {
	return c.birthDate
}

type ClientBet struct {
	client Client
	number string
}

func NewClientBet(clientConfig ClientConfig) (ClientBet, error) {
	client, err := NewClient(clientConfig.name, clientConfig.lastName, clientConfig.dni, clientConfig.birthDate)
	if err != nil {
		return ClientBet{}, err
	}

	if err := ValidateBetNumber(clientConfig.number); err != nil {
		return ClientBet{}, err
	}

	return ClientBet{
		client: client,
		number: clientConfig.number,
	}, nil
}

func (cb ClientBet) GetClient() Client {
	return cb.client
}

func (cb ClientBet) GetNumber() string {
	return cb.number
}

type ClientConfig struct {
	name      string
	lastName  string
	dni       string
	birthDate string
	number    string
}

func NewClientConfig(name string, lastName string, dni string, birthDate string, number string) ClientConfig {
	return ClientConfig{
		name:      name,
		lastName:  lastName,
		dni:       dni,
		birthDate: birthDate,
		number:    number,
	}
}
