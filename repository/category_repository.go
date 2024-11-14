package repository

import "unit-testing/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
	FindNameById(id string) *entity.Category
}
