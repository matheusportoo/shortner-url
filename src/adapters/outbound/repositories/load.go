package repositories_impl

import (
	"github.com/matheusportoo/shortner-url/src/adapters/outbound/redis_service"
	"github.com/matheusportoo/shortner-url/src/domain/ports/repositories"
)

type Methods struct {
	URL repositories.URL
}

func Load() *Methods {
	return &Methods{
		URL: NewURLRepositoryImpl(redis_service.RdsConnection),
	}
}
