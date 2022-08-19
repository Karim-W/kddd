package health

type HealthService struct {
	repo HealthRepository
}

func NewHealthService(
	repo HealthRepository,
) *HealthService {
	return &HealthService{
		repo: repo,
	}
}

func (s *HealthService) GetHealth() Health {
	return s.repo.GetHealth()
}
