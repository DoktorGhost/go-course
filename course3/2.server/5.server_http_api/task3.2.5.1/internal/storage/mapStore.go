package storage

import (
	"errors"
)

type StorageMap struct {
	store map[string][]byte
}

func NewStorageMap() *StorageMap {
	result := &StorageMap{
		store: make(map[string][]byte),
	}
	return result
}

func (sm *StorageMap) Create(login, password string) error {
	if _, ok := sm.store[login]; !ok {
		sm.store[login] = []byte(password)
		return nil
	}
	return errors.New("Username already exists")
}

func (sm *StorageMap) Update(loginOld, passwordNew string) error {
	if _, ok := sm.store[loginOld]; !ok {
		return errors.New("Login not found")
	}
	sm.store[loginOld] = []byte(passwordNew)
	return nil
}

func (sm *StorageMap) Delete(login string) error {
	if _, ok := sm.store[login]; !ok {
		return errors.New("Login not found")
	}

	delete(sm.store, login)
	return nil
}

func (sm *StorageMap) Read(login string) (string, error) {
	if _, ok := sm.store[login]; !ok {
		return "", errors.New("Login not found")
	}
	return string(sm.store[login]), nil
}
