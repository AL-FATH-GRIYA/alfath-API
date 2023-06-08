package repository

type repositoryPool struct {
	User UserRepository
}

func InitRepositoryInstance() *repositoryPool {
	return &repositoryPool{
		User: NewUserRepository(),
	}
}
