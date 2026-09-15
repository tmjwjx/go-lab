package main

import (
	"eventBus/modelA"
	"eventBus/modelB"
	"github.com/asaskevich/EventBus"
)

func main() {
	bus := EventBus.New()
	//bus2 := EventBus.New()

	a := modelA.Bus{EventBus: bus}
	b := modelB.Bus{EventBus: bus}

	//b.Receive()
	//a.Send()
	//
	//a.Receive()
	//b.Send()

	_ = a.EventBus.Subscribe("test", func() {
		println("test1")
	})
	_ = b.EventBus.Subscribe("test", func() {
		println("test2")
	})
	_ = bus.Subscribe("test", func() {
		println("test3")
	})
	b.EventBus.Publish("test")
}