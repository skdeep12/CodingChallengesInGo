package broker

type Topic interface {
	AddProducer() chan<- any
	AddConsumer(name string) <-chan any
	RemoveConsumer(name string)
	Start()
}

type topic struct {
	name      string
	producer  chan any
	consumers map[string]chan any
}

func (t *topic) Start() {
	go func() {
		for msg := range t.producer {
			for _, consumer := range t.consumers {
				consumer <- msg
			}
		}
	}()
}

// AddProducer returns a handle to produce messages on a topic
func (t *topic) AddProducer() chan<- any {
	return t.producer
}

// AddConsumer adds a consumer to the topic
func (t *topic) AddConsumer(name string) <-chan any {
	if consumer, ok := t.consumers[name]; !ok {
		consumer = make(chan any, 100)
		t.consumers[name] = consumer
		return consumer
	} else {
		return consumer
	}
}

func (t *topic) RemoveConsumer(name string) {
	if consumer, ok := t.consumers[name]; ok {
		close(consumer)
		delete(t.consumers, name)
	}
}

func NewTopic(name string, size int) Topic {
	t := &topic{
		name:      name,
		producer:  make(chan any, size),
		consumers: make(map[string]chan any),
	}
	t.Start()
	return t
}
