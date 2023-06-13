package repository

type repositoryPool struct {
	User   UserRepository
	Brand  BrandRepository
	Branch BranchRepository
}

func InitRepositoryInstance() *repositoryPool {
	return &repositoryPool{
		User:   NewUserRepository(),
		Brand:  NewBrandRepository(),
		Branch: NewBranchRepository(),
	}
}
