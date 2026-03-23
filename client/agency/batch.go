package agency

import (
	"encoding/csv"
	"os"

	"github.com/7574-sistemas-distribuidos/docker-compose-init/client/common"
	"github.com/7574-sistemas-distribuidos/docker-compose-init/client/communication"
)

const (
	MAX_BUFFER_SIZE = 8000
)

type BatchBuilder struct {
	csvReader        *csv.Reader
	batchSize        int
	maxBatchSize     int
	buffer           []byte
	maxBufferSize    int
	clientController common.ClientController
}

func NewBatchBuilder(path string, batchSize int) (BatchBuilder, error) {
	file, err := os.Open(path)
	if err != nil {
		return BatchBuilder{}, err
	}

	csvReader := csv.NewReader(file)
	clientController := common.NewClientController(common.NewClientService())
	return BatchBuilder{
		csvReader:        csvReader,
		batchSize:        0,
		maxBatchSize:     batchSize,
		buffer:           make([]byte, 0),
		maxBufferSize:    MAX_BUFFER_SIZE,
		clientController: clientController,
	}, nil
}

func (batch *BatchBuilder) ReadBatch() ([]byte, int, error) {
	for batch.batchSize < batch.maxBatchSize {
		record, err := batch.csvReader.Read()
		if err != nil {
			return batch.buffer, batch.batchSize, err
		}
		if record == nil {
			buffer := batch.buffer
			batch.buffer = make([]byte, 0)
			batchSize := batch.batchSize
			batch.batchSize = 0
			return buffer, batchSize, nil
		}

		clientConfig := common.NewClientConfigFromList(record)
		clientBet, err := batch.clientController.CreateClientBet(clientConfig)
		if err != nil {
			return batch.buffer, batch.batchSize, err
		}
		clientBetBytes := communication.SerializeClientBet(clientBet)
		if len(batch.buffer)+len(clientBetBytes) > batch.maxBufferSize {
			buffer := make([]byte, 0)
			buffer = append(buffer, byte(batch.batchSize))
			batch.buffer = clientBetBytes
			batchSize := batch.batchSize
			batch.batchSize = 1

			return buffer, batchSize, nil
		}
		batch.buffer = append(batch.buffer, clientBetBytes...)
		batch.batchSize++
	}

	buffer := batch.buffer
	batch.buffer = make([]byte, 0)
	batchSize := batch.batchSize
	batch.batchSize = 0

	return buffer, batchSize, nil

}
