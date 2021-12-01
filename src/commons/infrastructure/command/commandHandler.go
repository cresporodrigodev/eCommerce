package command

type CommandHandler interface {
	Handle(Command) error
}
