package repository

type repositoryPool struct {
	User  UserRepository
	Brand BrandRepository
}

func InitRepositoryInstance() *repositoryPool {
	return &repositoryPool{
		User:  NewUserRepository(),
		Brand: NewBrandRepository(),
	}
}
