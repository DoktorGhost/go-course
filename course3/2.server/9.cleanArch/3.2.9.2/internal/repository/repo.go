package repository

type StorageRepository interface {
	Create(login, password string) error
	Read(login string) (string, error)
}
