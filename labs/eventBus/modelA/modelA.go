package modelA

import (
	"fmt"
	"github.com/asaskevich/EventBus"
)

type Bus struct {
	EventBus EventBus.Bus
}

type ServerA interface {
	Run()
	Send()
}

func calculator(a int, b int) {
	fmt.Printf("modelA: %d\n", a+b)
}

func (bus Bus) Receive() {
	err := bus.EventBus.Subscribe("sendFromB", calculator)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func (bus Bus) Send() {
	bus.EventBus.Publish("sendFromA", 100, 100)
}