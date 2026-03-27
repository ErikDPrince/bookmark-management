package service

import (
	"context"

	"github.com/ErikDPrince/bookmark-management/internal/repository"
)

const (
	codeLength = 7
)

type ShortenURL interface {
	ShortenURL(ctx context.Context, url string) (string, error)
}

type shortenURLService struct {
	urlStorage repository.URLStorage
	codeGen    CodeGenerator
}

func NewShortenURLService(urlStorage repository.URLStorage, codeGen CodeGenerator) ShortenURL {
	return &shortenURLService{
		urlStorage: urlStorage,
		codeGen:    codeGen,
	}
}

func (s *shortenURLService) ShortenURL(ctx context.Context, url string) (string, error) {
	// gen code
	code, err := s.codeGen.GenerateCode(codeLength)
	if err != nil {
		return "", err
	}
	// call repo add code - url
	err = s.urlStorage.StoreURL(ctx, code, url)
	if err != nil {
		return "", err
	}
	// return code
	return code, nil
}
