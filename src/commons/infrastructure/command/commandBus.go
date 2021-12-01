package command

import "fmt"

type CommandBus struct {
	handlersMap map[string]CommandHandler
}

func NewCommandBus() CommandBus {
	return CommandBus{handlersMap: make(map[string]CommandHandler)}
}

func (cb *CommandBus) RegisterHandler(c Command, ch CommandHandler) {
	cmdId := c.CommandId()

	_, ok := cb.handlersMap[cmdId]
	if !ok {
		cb.handlersMap[cmdId] = ch
	}
}

func (cb CommandBus) Exec(c Command) error {
	cmdId := c.CommandId()

	ch, ok := cb.handlersMap[cmdId]
	if !ok {
		return fmt.Errorf("there not any CommandHandler associate to Command id %s", cmdId)
	}

	return ch.Handle(c)
}
