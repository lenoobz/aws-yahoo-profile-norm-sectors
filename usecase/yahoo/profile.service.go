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

// FindYahooProfiles finds all Yahoo profiles
func (s *Service) FindYahooProfiles(ctx context.Context) ([]*entities.YahooProfile, error) {
	s.log.Info(ctx, "finding all Yahoo profiles")
	return s.repo.FindYahooProfiles(ctx)
}
