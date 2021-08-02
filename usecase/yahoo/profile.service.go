package yahoo

import (
	"context"

	logger "github.com/lenoobz/aws-lambda-logger"
	"github.com/lenoobz/aws-yahoo-profile-norm-sectors/entities"
)

// Service sector
type Service struct {
	repo Repo
	log  logger.ContextLog
}

// NewService create new service
func NewService(repo Repo, log logger.ContextLog) *Service {
	return &Service{
		repo: repo,
		log:  log,
	}
}

// FindYahooProfilesByTickers finds all Yahoo profiles for given tickers
func (s *Service) FindYahooProfilesByTickers(ctx context.Context, tickers []string) ([]*entities.YahooProfile, error) {
	s.log.Info(ctx, "finding all Yahoo profiles by tickers", "tickers", tickers)
	return s.repo.FindYahooProfilesByTickers(ctx, tickers)
}

// FindYahooProfiles finds all Yahoo profiles
func (s *Service) FindYahooProfiles(ctx context.Context) ([]*entities.YahooProfile, error) {
	s.log.Info(ctx, "finding all Yahoo profiles")
	return s.repo.FindYahooProfiles(ctx)
}
