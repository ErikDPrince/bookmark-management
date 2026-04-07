package service

import (
	"context"
	"errors"

	"github.com/ErikDPrince/bookmark-management/internal/repository"
	"github.com/redis/go-redis/v9"
)

const (
	shortenCodeLen = 7
)

type ShortenURL interface {
	ShortenURL(ctx context.Context, url string) (string, error)
	GetURL(ctx context.Context, code string) (string, error)
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
	code, err := s.codeGen.GenerateCode(shortenCodeLen)
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

// GetURL looks up the code in Redis and returns the stored URL.

var ErrCodeNotExist = errors.New("code not exists")

func (s *shortenURLService) GetURL(ctx context.Context, code string) (string, error) {
	// call repo get url
	url, err := s.urlStorage.GetURL(ctx, code)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", ErrCodeNotExist
		}
		return "", err
	}
	// return url
	return url, nil
}
