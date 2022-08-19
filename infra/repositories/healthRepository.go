package repositories

import (
	"github.com/karim-w/kddd/domain/health"
	"github.com/karim-w/kddd/infra/cache"
)

type healthRepositoryImpl struct {
	cache *cache.Memcache
	// db    *sqlx.DB --> for a SQL Database
}

func NewHealthRepository(
	c *cache.Memcache,
) health.HealthRepository {
	return &healthRepositoryImpl{
		cache: c,
	}
}

func (r *healthRepositoryImpl) GetHealth() health.Health {
	if v, ok := r.cache.Get("health"); ok {
		return v.(health.Health)
	} else {
		return health.Health{
			Status: "OK",
		}
	}
}
