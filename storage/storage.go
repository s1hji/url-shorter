package storage

import "errors"

var ErrorNotFound = errors.New("ссылка не существует")
var ErrorSaved = errors.New("ссылка уже существует")

type Storage interface {
	Save(origin string, short string) error
	GetOriginLink(origin string) (string, error)
	GetShortLink(short string) (string, error)
}
