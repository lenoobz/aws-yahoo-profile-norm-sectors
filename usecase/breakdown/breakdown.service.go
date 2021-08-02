package breakdown

import (
	"context"

	logger "github.com/lenoobz/aws-lambda-logger"
	"github.com/lenoobz/aws-yahoo-profile-norm-sectors/usecase/yahoo"
)

// Service sector
type Service struct {
	breakdownRepo Repo
	yahooService  yahoo.Service
	log           logger.ContextLog
}

// NewService create new service
func NewService(breakdownRepo Repo, yahooService yahoo.Service, log logger.ContextLog) *Service {
	return &Service{
		breakdownRepo: breakdownRepo,
		yahooService:  yahooService,
		log:           log,
	}
}

// AddAssetSectorBreakdownByTickers adds new asset sector breakdown for given tickers
func (s *Service) AddAssetSectorBreakdownByTickers(ctx context.Context, tickers []string) error {
	s.log.Info(ctx, "adding new asset sector breakdown by tickers", "tickers", tickers)

	profiles, err := s.yahooService.FindYahooProfilesByTickers(ctx, tickers)
	if err != nil {
		s.log.Error(ctx, "find all Yahoo profiles by tickers failed", "error", err)
	}

	for _, profile := range profiles {
		sectorBreakdown := profile.MapYahooProfileToAssetSectorBreakdown(ctx, s.log)

		if err := s.breakdownRepo.InsertAssetSectorBreakdown(ctx, sectorBreakdown); err != nil {
			s.log.Error(ctx, "insert asset sector breakdown failed", "error", err)
			return err
		}
	}

	return nil
}

// AddAssetSectorBreakdown adds new asset sector breakdown
func (s *Service) AddAssetSectorBreakdown(ctx context.Context) error {
	s.log.Info(ctx, "adding new asset sector breakdown")

	profiles, err := s.yahooService.FindYahooProfiles(ctx)
	if err != nil {
		s.log.Error(ctx, "find all Yahoo profiles failed", "error", err)
	}

	for _, profile := range profiles {
		sectorBreakdown := profile.MapYahooProfileToAssetSectorBreakdown(ctx, s.log)

		if err := s.breakdownRepo.InsertAssetSectorBreakdown(ctx, sectorBreakdown); err != nil {
			s.log.Error(ctx, "insert asset sector breakdown failed", "error", err)
			return err
		}
	}

	return nil
}
