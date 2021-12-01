package repository

type Repository interface {
	Connect() error
	Close() error
}
