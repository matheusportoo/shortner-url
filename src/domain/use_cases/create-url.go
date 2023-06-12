package use_cases

import (
	"context"

	"github.com/matheusportoo/shortner-url/src/adapters/outbound/id_service"
	repositories_impl "github.com/matheusportoo/shortner-url/src/adapters/outbound/repositories"
	"github.com/matheusportoo/shortner-url/src/domain/entities"
)

type SaveURLParamsDTO struct {
	Link string `json:"link"`
}

type SaveURLResponseDTO struct {
	Link string `json:"link"`
	Id   string `json:"id"`
}

func CreateURL(params SaveURLParamsDTO) (error, SaveURLResponseDTO) {
	ctx := context.Background()
	repositories := repositories_impl.Load()

	err, id := id_service.Create(8)
	if err != nil {
		return err, SaveURLResponseDTO{}
	}

	url := entities.URL{
		Link:      params.Link,
		Id:        id,
		ExpiresIn: "24h",
	}

	err = repositories.URL.Save(ctx, url)
	if err != nil {
		return err, SaveURLResponseDTO{}
	}

	return nil, SaveURLResponseDTO{
		Link: params.Link,
		Id:   id,
	}
}
