package services

import (
	"testing"
	"unit-testing/entity"
	"unit-testing/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryServices = CategoryServices{Repository: categoryRepository}

func TestCategoryServicesGetNotFound(t *testing.T) {

	// program untuk mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryServices.Get("1")

	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryServicesGetSuccess(t *testing.T) {

	datas := entity.Category{
		Id:   "1",
		Name: "laptop",
	}

	// program untuk mock
	categoryRepository.Mock.On("FindNameById", "2").Return(datas)

	category, err := categoryServices.GetName("2")

	assert.Nil(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, datas.Id, category.Id)
	assert.Equal(t, datas.Name, category.Name)
}
