package modelB

import (
	"fmt"
	"github.com/asaskevich/EventBus"
)

type Bus struct {
	EventBus EventBus.Bus
}

type ServerB interface {
	Run()
	Send()
}

func calculator(a int, b int) {
	fmt.Printf("modelB: %d\n", a+b)
}

func (bus Bus) Receive() {
	err := bus.EventBus.Subscribe("sendFromA", calculator)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func (bus Bus) Send() {
	bus.EventBus.Publish("sendFromB", 10, 10)
}