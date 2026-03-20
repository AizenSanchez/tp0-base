package communication

import "github.com/7574-sistemas-distribuidos/docker-compose-init/client/common"

func SerializeString(str string) []byte {
	strSerialized := []byte(str)
	result := make([]byte, 0)
	result = append(result, []byte{byte(len(strSerialized))}...)
	result = append(result, strSerialized...)
	return result
}

func SerializeUint8(num uint8) []byte {
	uint8Serialized := []byte{byte(num)}
	result := make([]byte, 0)
	result = append(result, []byte{byte(len(uint8Serialized))}...)
	result = append(result, uint8Serialized...)
	return result
}

func SerializeUint64(num uint64) []byte {
	uint64Serialized := []byte{byte(num)}
	result := make([]byte, 0)
	result = append(result, []byte{byte(len(uint64Serialized))}...)
	result = append(result, uint64Serialized...)
	return result
}

func SerializeClientBet(clientBet common.ClientBet) []byte {
	clientBetSerialized := make([]byte, 0)
	clientBetSerialized = append(clientBetSerialized, SerializeClient(clientBet.GetClient())...)
	clientBetSerialized = append(clientBetSerialized, SerializeUint64(uint64(clientBet.GetNumber()))...)

	result := make([]byte, 0)
	result = append(result, []byte{byte(len(clientBetSerialized))}...)
	result = append(result, clientBetSerialized...)
	return result
}

func SerializeClient(client common.Client) []byte {
	clientSerialized := make([]byte, 0)
	clientSerialized = append(clientSerialized, SerializeString(client.GetName())...)
	clientSerialized = append(clientSerialized, SerializeString(client.GetLastName())...)
	clientSerialized = append(clientSerialized, SerializeUint64(client.GetDni())...)
	clientSerialized = append(clientSerialized, SerializeBirthDate(client.GetBirthDate())...)

	result := make([]byte, 0)
	result = append(result, []byte{byte(len(clientSerialized))}...)
	result = append(result, clientSerialized...)
	return result
}

func SerializeBirthDate(birthDate common.BirthDate) []byte {
	dateSerialized := make([]byte, 0)
	dateSerialized = append(dateSerialized, SerializeUint8(birthDate.GetDay())...)
	dateSerialized = append(dateSerialized, SerializeUint8(birthDate.GetMonth())...)
	dateSerialized = append(dateSerialized, SerializeUint64(uint64(birthDate.GetYear()))...)

	result := make([]byte, 0)
	result = append(result, []byte{byte(len(dateSerialized))}...)
	result = append(result, dateSerialized...)
	return result
}
