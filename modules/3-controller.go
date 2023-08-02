package modules

import (
	"github.com/purawaktra/raung1-go/utils"
	"github.com/segmentio/kafka-go"
)

type Raung1Controller struct {
	uc Raung1Usecase
}

func CreateRaung1Controller(uc Raung1Usecase) Raung1Controller {
	return Raung1Controller{
		uc: uc,
	}
}

func (ctrl Raung1Controller) InitKafkaController(messages <-chan kafka.Message, done chan<- bool) {
	// start goroutine for blockers loop
	utils.Debug("InitKafkaController", "start go routine for kafka controller")
	go func(messages <-chan kafka.Message, done chan<- bool) {
		// create blocker for go routine
		for message := range messages {
			ctrl.AccountCreation(message)
		}

		// create success termination by send value to channel
		done <- true
		return
	}(messages, done)

	// create return
	return
}

func (ctrl Raung1Controller) AccountCreation(message kafka.Message) {

}
