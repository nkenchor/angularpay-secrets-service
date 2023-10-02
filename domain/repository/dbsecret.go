package repository

import "secret/domain/entity"

type SecretRepository interface {
	CreateSecret(entity.SecretStruct) (interface{}, error)
	UpdateSecret(string, entity.SecretStruct) (interface{}, error)
	DeleteSecret(string, entity.SecretStruct) (interface{}, error)
	GetSecretByRef(string, entity.SecretStruct) (interface{}, error)
	GetAllSecret(entity.SecretStruct, []entity.SecretStruct) (interface{}, error)
	GetServiceSecretList(string, entity.SecretStruct) (interface{}, error)
}