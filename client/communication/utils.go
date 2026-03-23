package communication

import "github.com/7574-sistemas-distribuidos/docker-compose-init/client/common"

func SerializeString(str string) []byte {
	strSerialized := []byte(str)
	result := make([]byte, 0)
	result = append(result, byte(len(strSerialized)))
	result = append(result, strSerialized...)
	return result
}

func SerializeClientBet(clientBet common.ClientBet) []byte {
	clientBetSerialized := make([]byte, 0)
	clientBetSerialized = append(clientBetSerialized, SerializeClient(clientBet.GetClient())...)
	clientBetSerialized = append(clientBetSerialized, SerializeString(clientBet.GetNumber())...)

	result := make([]byte, 0)
	result = append(result, byte(len(clientBetSerialized)))
	result = append(result, clientBetSerialized...)
	return result
}

func SerializeClient(client common.Client) []byte {
	clientSerialized := make([]byte, 0)
	clientSerialized = append(clientSerialized, SerializeString(client.GetName())...)
	clientSerialized = append(clientSerialized, SerializeString(client.GetLastName())...)
	clientSerialized = append(clientSerialized, SerializeString(client.GetDni())...)
	clientSerialized = append(clientSerialized, SerializeString(client.GetBirthDate())...)

	result := make([]byte, 0)
	result = append(result, byte(len(clientSerialized)))
	result = append(result, clientSerialized...)
	return result
}
