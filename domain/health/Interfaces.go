package health

type HealthRepository interface {
	GetHealth() Health
}
