package application

import (
	"secret/domain/entity"
	"secret/domain/repository"
)

type SecretApp struct {
	theSecret repository.SecretRepository
}

var _SecretApplication = &SecretApp{}

type SecretApplication interface {
	CreateSecret(entity.SecretStruct) (interface{}, error)
	UpdateSecret(string, entity.SecretStruct) (interface{}, error)
	DeleteSecret(string, entity.SecretStruct) (interface{}, error)
	GetSecretByRef(string, entity.SecretStruct) (interface{}, error)
	GetAllSecret(entity.SecretStruct, []entity.SecretStruct) (interface{}, error)
	GetServiceSecretList(string, entity.SecretStruct) (interface{}, error)
}

func (u *SecretApp) CreateSecret(c entity.SecretStruct) (interface{}, error) {
	return u.theSecret.CreateSecret(c)
}

func (u *SecretApp) UpdateSecret(ref string, c entity.SecretStruct) (interface{}, error) {
	return u.theSecret.UpdateSecret(ref, c)
}


func (u *SecretApp) DeleteSecret(ref string, c entity.SecretStruct) (interface{}, error) {
	return u.theSecret.DeleteSecret(ref, c)
}

func (u *SecretApp) GetSecretByRef(ref string, c entity.SecretStruct) (interface{}, error) {
	return u.theSecret.GetSecretByRef(ref, c)
}


func (u *SecretApp) GetAllSecret(c entity.SecretStruct, C []entity.SecretStruct) (interface{}, error) {
	return u.theSecret.GetAllSecret(c, C)
}


func (u *SecretApp) GetServiceSecretList(ref string, c entity.SecretStruct) (interface{}, error) {
	return u.theSecret.GetServiceSecretList(ref, c)
}