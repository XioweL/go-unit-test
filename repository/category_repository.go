package repository

import "github.com/XioweL/go-unit-test.git/helper/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
