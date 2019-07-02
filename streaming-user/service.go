package user

// Service methods.
type Service interface{}

// ServiceImpl service dependecies.
type ServiceImpl struct {
	repository Repository
}

// NewServiceImpl service constructor.
func NewServiceImpl(repository Repository) *ServiceImpl {
	return &ServiceImpl{
		repository: repository,
	}
}
