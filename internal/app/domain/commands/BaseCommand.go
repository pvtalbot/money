package commands

type CommandHandler[C any] interface {
	Handle(cmd C) error
}
