package communication

import (
	"net"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("log")

type Connection struct {
	conn net.Conn
}

func NewConnection(serverAddress string, clientId string) (Connection, error) {
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		log.Criticalf(
			"action: connect | result: fail | client_id: %v | server_address: %v | error: %v",
			clientId,
			serverAddress,
			err,
		)
		return Connection{}, err
	}
	return Connection{
		conn: conn,
	}, nil
}

func (connection *Connection) SendMessage(protocolFrame ProtocolFrame) error {
	messageSerialized := protocolFrame.Serialize()
	bytesSent, err := connection.conn.Write(messageSerialized)
	if err != nil {
		log.Errorf(
			"action: send_message | result: fail | error: %v",
			err,
		)
		return err
	}
	if bytesSent != len(messageSerialized) {
		log.Errorf(
			"action: send_message | result: fail | error: bytes sent %v does not match message size %v",
			bytesSent,
			len(messageSerialized),
		)
		return err
	}
	log.Infof(
		"action: send_message | result: success | bytes_sent: %v",
		bytesSent,
	)
	return nil
}

func (connection *Connection) ReceiveMessage() (ProtocolFrame, error) {
	protocolHeader, err := readHeader(connection.conn)
	if err != nil {
		return ProtocolFrame{}, err
	}
	protocolBody, err := readBody(connection.conn, protocolHeader.GetMessageSize())
	if err != nil {
		return ProtocolFrame{}, err
	}

	protocolFrame := ProtocolFrame{
		header: protocolHeader,
		body:   protocolBody,
	}

	return protocolFrame, nil
}

func readHeader(conn net.Conn) (ProtocolHeader, error) {
	headerBytes := make([]byte, 2)
	bytesRead, err := conn.Read(headerBytes)
	if err != nil {
		log.Errorf(
			"action: read_header | result: fail | error: %v",
			err,
		)
		return ProtocolHeader{}, err
	}
	if bytesRead != len(headerBytes) {
		log.Errorf(
			"action: read_header | result: fail | error: bytes read %v does not match header size %v",
			bytesRead,
			len(headerBytes),
		)
		return ProtocolHeader{}, err
	}

	protocolHeader, err := DeserializeHeader(headerBytes)
	if err != nil {
		log.Errorf(
			"action: read_header | result: fail | error: %v",
			err,
		)
		return ProtocolHeader{}, err
	}
	log.Infof(
		"action: read_header | result: success | bytes_read: %v",
		bytesRead,
	)
	return protocolHeader, nil
}

func readBody(conn net.Conn, bodySize uint8) ([]byte, error) {
	bodyBytes := make([]byte, bodySize)
	bytesRead, err := conn.Read(bodyBytes)
	if err != nil {
		log.Errorf(
			"action: read_body | result: fail | error: %v",
			err,
		)
		return []byte{}, err
	}
	if bytesRead != len(bodyBytes) {
		log.Errorf(
			"action: read_body | result: fail | error: bytes read %v does not match body size %v",
			bytesRead,
			len(bodyBytes),
		)
		return []byte{}, err
	}
	log.Infof(
		"action: read_body | result: success | bytes_read: %v",
		bytesRead,
	)
	return bodyBytes, nil
}

func (connection *Connection) Close() error {
	err := connection.conn.Close()
	if err != nil {
		log.Errorf(
			"action: close_connection | result: fail | error: %v",
			err,
		)
		return err
	}
	log.Infof(
		"action: close_connection | result: success",
	)
	return nil
}
