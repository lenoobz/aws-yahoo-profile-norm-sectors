package entities

import (
	"context"
	"fmt"
	"strings"

	logger "github.com/lenoobz/aws-lambda-logger"
	"github.com/lenoobz/aws-yahoo-profile-norm-sectors/consts"
)

// YahooProfile struct
type YahooProfile struct {
	Ticker string `json:"ticker,omitempty"`
	Sector string `json:"sector,omitempty"`
}

// MapYahooProfileToAssetSectorBreakdown map Yahoo profile to asset sector breakdown
func (f *YahooProfile) MapYahooProfileToAssetSectorBreakdown(ctx context.Context, log logger.ContextLog) *FundBreakdown {
	sectorCode, err := getSectorCode(f.Sector)
	if err != nil {
		log.Warn(ctx, "get sector code failed", "error", err, "ticker", f.Ticker, "sector name", f.Sector)
	}

	sectorBreakdown := &SectorBreakdown{
		SectorCode:  sectorCode,
		SectorName:  f.Sector,
		FundPercent: 100,
	}

	var sectorsBreakdown []*SectorBreakdown
	sectorsBreakdown = append(sectorsBreakdown, sectorBreakdown)

	fundBreakdown := &FundBreakdown{
		Ticker:     f.Ticker,
		AssetClass: consts.EQUITY,
		Sectors:    sectorsBreakdown,
	}

	return fundBreakdown
}

// getSectorCode gets sector code of a given name
func getSectorCode(name string) (string, error) {
	for _, sector := range consts.Sectors {
		if strings.EqualFold(strings.ToUpper(sector.Name), strings.ToUpper(name)) {
			return sector.Code, nil
		}
	}
	return "OTH", fmt.Errorf("cannot find sector code for sector %s", name)
}
