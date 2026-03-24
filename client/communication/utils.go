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

func DeserializeClientBet(bytes []byte) (common.ClientBet, error) {
	offset := 0
	name, size := DeserializeString(bytes[offset:])
	offset += size
	lastName, size := DeserializeString(bytes[offset:])
	offset += size
	dni, size := DeserializeString(bytes[offset:])
	offset += size
	birthDate, size := DeserializeString(bytes[offset:])
	offset += size
	number, size := DeserializeString(bytes[offset:])
	offset += size

	clientConfig := common.NewClientConfig(name, lastName, dni, birthDate, number)
	clientBet, err := common.NewClientBet(clientConfig)
	if err != nil {
		return common.ClientBet{}, err
	}
	return clientBet, nil
}

func DeserializeString(bytes []byte) (string, int) {
	strSize := int(bytes[0])
	str := string(bytes[1 : 1+strSize])
	return str, 1 + strSize
}
