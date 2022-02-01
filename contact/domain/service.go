package domain

type ContactRepository interface {
	Get() []string
	Save(string) string
}

type ContactService struct {
	repository ContactRepository
}

func NewContactService(cr ContactRepository) *ContactService {
	return &ContactService{
		repository: cr,
	}
}

func (cs *ContactService) Get() []string {
	return cs.repository.Get()
}

func (cs *ContactService) Save(s string) string {
	return cs.repository.Save(s)
}
