package common

import "strconv"

type BirthDate struct {
	day   uint8
	month uint8
	year  uint64
}

func (bd BirthDate) GetDay() uint8 {
	return bd.day
}

func (bd BirthDate) GetMonth() uint8 {
	return bd.month
}

func (bd BirthDate) GetYear() uint64 {
	return bd.year
}

type Client struct {
	name      string
	lastName  string
	dni       uint64
	birthDate BirthDate
}

func NewClient(name string, lastName string, dni string, birthDate string) (Client, error) {
	dniUint, err := strconv.ParseUint(dni, 10, 64)
	if err != nil {
		return Client{}, err
	}

	birthDateObj, err := BirthDateFromString(birthDate)
	if err != nil {
		return Client{}, err
	}

	return Client{
		name:      name,
		lastName:  lastName,
		dni:       dniUint,
		birthDate: birthDateObj,
	}, nil
}

func (c Client) GetName() string {
	return c.name
}

func (c Client) GetLastName() string {
	return c.lastName
}

func (c Client) GetDni() uint64 {
	return c.dni
}

func (c Client) GetBirthDate() BirthDate {
	return c.birthDate
}

type ClientBet struct {
	client Client
	number uint64
}

func NewClientBet(clientConfig ClientConfig) (ClientBet, error) {
	client, err := NewClient(clientConfig.name, clientConfig.lastName, clientConfig.dni, clientConfig.birthDate)
	if err != nil {
		return ClientBet{}, err
	}

	number, err := strconv.ParseUint(clientConfig.number, 10, 64)
	if err != nil {
		return ClientBet{}, err
	}

	return ClientBet{
		client: client,
		number: number,
	}, nil
}

func (cb ClientBet) GetClient() Client {
	return cb.client
}

func (cb ClientBet) GetNumber() uint64 {
	return cb.number
}

type ClientConfig struct {
	name      string
	lastName  string
	dni       string
	birthDate string
	number    string
}
