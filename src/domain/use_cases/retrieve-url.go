package use_cases

import (
	"context"

	repositories_impl "github.com/matheusportoo/shortner-url/src/adapters/outbound/repositories"
)

type RetrieURLParamsDTO struct {
	Id string
}

type RetrieURLResponseDTO struct {
	Id   string
	Link string
}

func RetrieURL(params RetrieURLParamsDTO) (error, RetrieURLResponseDTO) {
	ctx := context.Background()
	repositories := repositories_impl.Load()

	err, url := repositories.URL.Retrieve(ctx, params.Id)
	if err != nil {
		return err, RetrieURLResponseDTO{}
	}

	response := RetrieURLResponseDTO{
		Id:   url.Id,
		Link: url.Link,
	}

	return nil, response
}
