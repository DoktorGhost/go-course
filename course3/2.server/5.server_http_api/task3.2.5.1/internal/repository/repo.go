package repository

type StorageRepository interface {
	Create(login, password string) error
	Read(login string) (string, error)
	Update(loginOld, passwordNew string) error
	Delete(login string) error
}
