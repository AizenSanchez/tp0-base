package agency

import (
	"fmt"
	"strconv"

	"github.com/7574-sistemas-distribuidos/docker-compose-init/client/common"
	"github.com/7574-sistemas-distribuidos/docker-compose-init/client/communication"
	"github.com/op/go-logging"
)

const (
	FILE_PATH = "/data/agency-%s.csv"
)

var log = logging.MustGetLogger("log")

type AgencyClient struct {
	id            int
	batch         BatchBuilder
	serverAddress string
	stopChannel   chan struct{}
}

func NewAgencyClient(id string, batchSize int, serverAddress string, stopChannel chan struct{}) (AgencyClient, error) {
	batch, err := NewBatchBuilder(fmt.Sprintf(FILE_PATH, id), batchSize)
	if err != nil {
		return AgencyClient{}, err
	}

	agencyId, err := strconv.Atoi(id)
	if err != nil {
		return AgencyClient{}, err
	}
	return AgencyClient{
		id:            agencyId,
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
		bodyBytes := make([]byte, 0)
		bodyBytes = append(bodyBytes, byte(agencyClient.id))
		bodyBytes = append(bodyBytes, batch...)
		protocolFrame := communication.NewProtocolFrameRequestRegisterBets(bodyBytes)
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

func (agencyClient *AgencyClient) AskForWinners() {
	for {
		serverConnection, err := communication.NewConnection(agencyClient.serverAddress, agencyClient.id)
		if err != nil {
			log.Errorf("action: ask_for_winners | result: fail | error: %v", err)
			return
		}
		protocolFrame := communication.NewProtocolFrameRequestAskForWinners(agencyClient.id)
		err = serverConnection.SendMessage(protocolFrame)
		if err != nil {
			log.Errorf("action: ask_for_winners | result: fail | error: %v", err)
			return
		}
		protocolFrameResponse, err := serverConnection.ReceiveMessage()
		if err != nil {
			log.Errorf("action: ask_for_winners | result: fail | error: %v", err)
			return
		}
		winners, results_ready, err := protocolFrameResponse.GetWinners()
		if err != nil {
			log.Errorf("action: ask_for_winners | result: fail | error: %v", err)
			return
		}
		if !results_ready {
			log.Infof("action: ask_for_winners | result: wait | agency_id: %v", agencyClient.id)
			serverConnection.Close()
			continue
		}
		log.Infof("action: ask_for_winners | result: success | agency_id: %v | winners_count: %v", agencyClient.id, len(winners))
		serverConnection.Close()
		agencyClient.showWinners(winners)
		break
	}
}

func (agencyClient *AgencyClient) showWinners(winners []common.ClientBet) {
	for _, winner := range winners {
		log.Infof("action: winner | result: success | client_name: %v | client_last_name: %v | client_dni: %v | client_birth_date: %v | bet_number: %v",
			winner.GetClient().GetName(),
			winner.GetClient().GetLastName(),
			winner.GetClient().GetDni(),
			winner.GetClient().GetBirthDate(),
			winner.GetNumber(),
		)
	}
}
