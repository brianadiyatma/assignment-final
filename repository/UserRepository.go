package repository

import "assignment-final/model/domain"

type UserRepository interface {
	Create(domain.User) (domain.User, error)
}
