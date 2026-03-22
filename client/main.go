package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/op/go-logging"
	"github.com/spf13/viper"

	"github.com/7574-sistemas-distribuidos/docker-compose-init/client/common"
	"github.com/7574-sistemas-distribuidos/docker-compose-init/client/communication"
)

var log = logging.MustGetLogger("log")

// InitConfig Function that uses viper library to parse configuration parameters.
// Viper is configured to read variables from both environment variables and the
// config file ./config.yaml. Environment variables takes precedence over parameters
// defined in the configuration file. If some of the variables cannot be parsed,
// an error is returned
func InitConfig() (*viper.Viper, error) {
	v := viper.New()

	// Configure viper to read env variables with the CLI_ prefix
	v.AutomaticEnv()
	v.SetEnvPrefix("cli")
	// Use a replacer to replace env variables underscores with points. This let us
	// use nested configurations in the config file and at the same time define
	// env variables for the nested configurations
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Add env variables supported
	v.BindEnv("id")
	v.BindEnv("server", "address")
	v.BindEnv("log", "level")
	v.BindEnv("nombre", "NOMBRE")
	v.BindEnv("apellido", "APELLIDO")
	v.BindEnv("dni", "DNI")
	v.BindEnv("nacimiento", "NACIMIENTO")
	v.BindEnv("numero", "NUMERO")
	// Try to read configuration from config file. If config file
	// does not exists then ReadInConfig will fail but configuration
	// can be loaded from the environment variables so we shouldn't
	// return an error in that case
	v.SetConfigFile("./data_client/config.yaml")
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("Configuration could not be read from config file. Using env variables instead")
	}

	return v, nil
}

// InitLogger Receives the log level to be set in go-logging as a string. This method
// parses the string and set the level to the logger. If the level string is not
// valid an error is returned
func InitLogger(logLevel string) error {
	baseBackend := logging.NewLogBackend(os.Stdout, "", 0)
	format := logging.MustStringFormatter(
		`%{time:2006-01-02 15:04:05} %{level:.5s}     %{message}`,
	)
	backendFormatter := logging.NewBackendFormatter(baseBackend, format)

	backendLeveled := logging.AddModuleLevel(backendFormatter)
	logLevelCode, err := logging.LogLevel(logLevel)
	if err != nil {
		return err
	}
	backendLeveled.SetLevel(logLevelCode, "")

	// Set the backends to be used.
	logging.SetBackend(backendLeveled)
	return nil
}

// PrintConfig Print all the configuration parameters of the program.
// For debugging purposes only
func PrintConfig(v *viper.Viper) {
	log.Infof("action: config | result: success | client_id: %s | server_address: %s | client_name: %s | client_lastname: %s | client_dni: %s | client_birthdate: %s | client_bet_number: %s | log_level: %s",
		v.GetString("id"),
		v.GetString("server.address"),
		v.GetString("nombre"),
		v.GetString("apellido"),
		v.GetString("dni"),
		v.GetString("nacimiento"),
		v.GetString("numero"),
		v.GetString("log.level"),
	)
}

func main() {
	v, err := InitConfig()
	if err != nil {
		log.Criticalf("%s", err)
		os.Exit(1)
	}

	if err := InitLogger(v.GetString("log.level")); err != nil {
		log.Criticalf("%s", err)
		os.Exit(1)
	}

	// Print program config with debugging purposes
	PrintConfig(v)

	clientConfig := common.NewClientConfig(
		v.GetString("nombre"),
		v.GetString("apellido"),
		v.GetString("dni"),
		v.GetString("nacimiento"),
		v.GetString("numero"),
	)

	clientService := common.NewClientService(common.ClientBet{})
	_, err = clientService.CreateClientBet(clientConfig)
	if err != nil {
		log.Criticalf("%s", err)
		os.Exit(1)
	}

	clientController := common.NewClientController(clientService)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM)
	stop := make(chan struct{})
	go func() {
		<-sigChan
		close(stop)
	}()

	connection, err := communication.NewConnection(v.GetString("server.address"), v.GetString("id"))
	if err != nil {
		log.Criticalf("%s", err)
		os.Exit(1)
	}

	registerBet(clientController, connection, v.GetString("id"))
	connection.Close()
	os.Exit(0)
}

func registerBet(clientController common.ClientController, connection *communication.Connection, clientID string) {
	protocolFrame := communication.NewProtocolFrameRequestRegisterBet(clientController.GetClientBet())
	if err := connection.SendMessage(protocolFrame); err != nil {
		log.Criticalf("%s", err)
		return
	}
	protocolFrameResponse, err := connection.ReceiveMessage()
	if err != nil {
		log.Criticalf("%s", err)
		return
	}

	if protocolFrameResponse.GetMessageType() == 2 {
		log.Infof("action: apuesta_enviada | result: success | dni: %v| numero: %v",
			clientController.GetClientDNI(),
			clientController.GetClientBetNumber(),
		)
	} else if protocolFrameResponse.GetMessageType() == 3 {
		log.Infof("action: apuesta_enviada | result: fail | dni: %v| numero: %v",
			clientController.GetClientDNI(),
			clientController.GetClientBetNumber())
	}

}
