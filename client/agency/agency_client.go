package agency

import (
	"fmt"

	"github.com/7574-sistemas-distribuidos/docker-compose-init/client/communication"
	"github.com/op/go-logging"
)

const (
	FILE_PATH = "/data/agency-%s.csv"
)

var log = logging.MustGetLogger("log")

type AgencyClient struct {
	id            string
	batch         BatchBuilder
	serverAddress string
	stopChannel   chan struct{}
}

func NewAgencyClient(id string, batchSize int, serverAddress string, stopChannel chan struct{}) (AgencyClient, error) {
	batch, err := NewBatchBuilder(fmt.Sprintf(FILE_PATH, id), batchSize)
	if err != nil {
		return AgencyClient{}, err
	}
	return AgencyClient{
		id:            id,
		batch:         batch,
		serverAddress: serverAddress,
		stopChannel:   stopChannel,
	}, nil
}

func (agencyClient *AgencyClient) RegisterBets() error {
	for {
		select {
		case <-agencyClient.stopChannel:
			log.Infof("action: stop_registering_bets | result: success")
			return nil
		default:
		}
		batch, batchSize, err := agencyClient.batch.ReadBatch()
		if err != nil {
			return err
		}
		if len(batch) == 0 {
			break
		}
		serverConnection, err := communication.NewConnection(agencyClient.serverAddress, agencyClient.id)
		if err != nil {
			return err
		}
		protocolFrame := communication.NewProtocolFrameRequestRegisterBets(batch)
		err = serverConnection.SendMessage(protocolFrame)
		if err != nil {
			return err
		}
		protocolFrameResponse, err := serverConnection.ReceiveMessage()
		if err != nil {
			return err
		}

		if protocolFrameResponse.GetMessageType() == 2 {
			log.Infof("action: apuesta_enviada | result: success | cantidad: %v",
				batchSize,
			)
		} else if protocolFrameResponse.GetMessageType() == 3 {
			log.Infof("action: apuesta_enviada | result: fail | cantidad: %v",
				batchSize)
		}

		serverConnection.Close()
	}
	return nil
}
