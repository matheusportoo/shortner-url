package repositories

import (
	"context"

	"github.com/matheusportoo/shortner-url/src/domain/entities"
)

type URL interface {
	Retrieve(ctx context.Context, id string) (error, entities.URL)
	Save(ctx context.Context, url entities.URL) error
}
