package events

type EventDispatcher struct {
	handlers map[string][]EventHandler
}

type EventHandler interface {
	Handle(event interface{}) error
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (d *EventDispatcher) Register(eventType string, handler EventHandler) {
	d.handlers[eventType] = append(d.handlers[eventType], handler)
}

func (d *EventDispatcher) Dispatch(eventType string, event interface{}) error {
	handlers, exists := d.handlers[eventType]
	if !exists {
		return nil
	}

	for _, handler := range handlers {
		if err := handler.Handle(event); err != nil {
			return err
		}
	}
	return nil
}