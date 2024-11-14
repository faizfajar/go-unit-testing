package services

import (
	"errors"
	"unit-testing/entity"
	"unit-testing/repository"
)

type CategoryServices struct {
	Repository repository.CategoryRepository
}

func (service CategoryServices) Get(id string) (*entity.Category, error) {

	category := service.Repository.FindById(id)

	if category == nil {
		return nil, errors.New("category not found")
	} else {
		return category, nil
	}

}

func (service CategoryServices) GetName(id string) (*entity.Category, error) {

	category := service.Repository.FindNameById(id)

	if category == nil {
		return nil, errors.New("category not found")
	} else {
		return category, nil
	}

}
