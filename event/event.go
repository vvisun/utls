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

// 获取事件总线实例
func EventInstance() *eventBusWrapper {
	once.Do(func() {
		ebr = &eventBusWrapper{
			eb: evtbus.New(),
		}
	})
	return ebr
}

// 监听topic事件,同步
func (e *eventBusWrapper) Subscribe(topic string, fn interface{}) error {
	return e.eb.Subscribe(topic, fn)
}

// 监听topic事件,异步
func (e *eventBusWrapper) SubscribeAsync(topic string, fn interface{}, transactional bool) error {
	return e.eb.SubscribeAsync(topic, fn, transactional)
}

// 发布topic事件
func (e *eventBusWrapper) Publish(topic string, args ...interface{}) {
	e.eb.Publish(topic, args...)
}

// 注销topic事件的fn监听
func (e *eventBusWrapper) Unsubscribe(topic string, fn interface{}) {
	e.eb.Unsubscribe(topic, fn)
}
