package category

type Service interface {
	CreateCategory(category *Category)
}

type serviceImpl struct{
	Repo CategoryRepository `inject:"caterepo"`
}


func GetCategoryService() Service{
	return &serviceImpl{}
}

func (s *serviceImpl) CreateCategory(category *Category) {
	s.Repo.Save(category)
}