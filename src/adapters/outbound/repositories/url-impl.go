package repositories_impl

import (
	"context"
	"encoding/json"

	"github.com/matheusportoo/shortner-url/src/domain/entities"
	"github.com/redis/go-redis/v9"
)

type URLRepositoryImpl struct {
	db *redis.Client
}

func NewURLRepositoryImpl(redis *redis.Client) *URLRepositoryImpl {
	return &URLRepositoryImpl{db: redis}
}

func (u *URLRepositoryImpl) Retrieve(ctx context.Context, id string) (error, entities.URL) {
	value, err := u.db.Get(ctx, id).Result()
	if err != nil {
		return err, entities.URL{}
	}

	var url entities.URL

	json.Unmarshal([]byte(value), &url)

	return nil, url
}

func (u *URLRepositoryImpl) Save(ctx context.Context, url entities.URL) error {
	value, err := json.Marshal(url)
	if err != nil {
		return err
	}

	err = u.db.Set(ctx, url.Id, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
