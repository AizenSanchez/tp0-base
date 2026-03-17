package common

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

func (cb ClientBet) GetClient() Client {
	return cb.client
}

func (cb ClientBet) GetNumber() uint64 {
	return cb.number
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
