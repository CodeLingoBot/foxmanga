package category

import (
	"github.com/go-xorm/xorm"
)

type CategoryRepository interface {
	Save(cate *Category)
	Update(cate *Category)
	FindById(id string) *Category
}

func CreateRepository() CategoryRepository {
	return &categoryRepository{}
}

type categoryRepository struct {
	DB *xorm.Engine `inject:"database"`
}

func (c *categoryRepository) Save(cate *Category) {
	c.DB.Insert(cate)
}

func (c *categoryRepository) Update(cate *Category) {
	c.DB.Update(cate)
}

func (c *categoryRepository) FindById(id string) *Category {
	cate := &Category{ID: id}
	c.DB.Get(cate)
	return cate
}
