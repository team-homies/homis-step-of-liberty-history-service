package service

import (
	"main/app/api/dex/resource"
	"main/database/repository"
)

type DexService interface {
	GetQuote() (res *resource.GetQuoteResponse, err error)
}

func NewDexService() DexService {
	return &dexService{
		DexService: &dexService{},
	}
}

type dexService struct {
	DexService
}

func (d *dexService) GetQuote() (res *resource.GetQuoteResponse, err error) {
	var Id uint

	quote, err := repository.NewRepository().FindQuote(Id)
	if err != nil {
		return
	}

	res = &resource.GetQuoteResponse{
		Content:  quote.Content,
		ImageUrl: quote.ImageUrl,
	}
	return
}


