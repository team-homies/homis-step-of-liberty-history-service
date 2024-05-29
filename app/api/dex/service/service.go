package service

import (
	"main/app/api/dex/resource"
	"main/database/repository"
)

type DexService interface {
	GetQuote() (res *resource.GetQuoteResponse, err error)
	GetTags() (res []resource.GetTagsResponse, err error)
}

func NewDexService() DexService {
	return &dexService{
		DexService: &dexService{},
	}
}

type dexService struct {
	DexService
}

func (d *dexService) GetTags() (res []resource.GetTagsResponse, err error) {
	res = []resource.GetTagsResponse{}
	tags, err := repository.NewRepository().GetTags()
	if err != nil {
		return
	}

	for _, tag := range tags {
		res = append(res, resource.GetTagsResponse{
			Id:   tag.ID,
			Name: tag.Name,
		})
	}
	return res, err
}

func (d *dexService) GetQuote() (res *resource.GetQuoteResponse, err error) {
	quote, err := repository.NewRepository().GetQuote()
	if err != nil {
		return
	}
	res = &resource.GetQuoteResponse{
		Content:  quote.Content,
		ImageUrl: quote.ImageUrl,
	}
	return
}
