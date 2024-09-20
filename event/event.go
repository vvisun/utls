package event

import (
	"sync"

	evtbus "github.com/asaskevich/EventBus"
)

type eventBusWrapper struct {
	eb evtbus.Bus
}

var ebr *eventBusWrapper
var once sync.Once

func EventInstance() *eventBusWrapper {
	once.Do(func() {
		ebr = &eventBusWrapper{
			eb: evtbus.New(),
		}
	})
	return ebr
}

func (e *eventBusWrapper) Subscribe(topic string, fn interface{}) error {
	return e.eb.Subscribe(topic, fn)
}

func (e *eventBusWrapper) SubscribeAsync(topic string, fn interface{}, transactional bool) error {
	return e.eb.SubscribeAsync(topic, fn, transactional)
}

func (e *eventBusWrapper) Publish(topic string, args ...interface{}) {
	e.eb.Publish(topic, args...)
}
