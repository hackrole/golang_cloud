package msgqueue

type EventListner struct {
  Listen(eventNames ...string) (<-chan Event, <-chan error, error)
}

