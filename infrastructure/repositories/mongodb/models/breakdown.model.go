package models

import (
	"context"
	"time"

	logger "github.com/lenoobz/aws-lambda-logger"
	"github.com/lenoobz/aws-yahoo-profile-norm-sectors/consts"
	"github.com/lenoobz/aws-yahoo-profile-norm-sectors/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FundBreakdownModel struct
type FundBreakdownModel struct {
	ID         *primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt  int64               `bson:"createdAt,omitempty"`
	ModifiedAt int64               `bson:"modifiedAt,omitempty"`
	Enabled    bool                `bson:"enabled"`
	Deleted    bool                `bson:"deleted"`
	Schema     string              `bson:"schema,omitempty"`
	Source     string              `bson:"source,omitempty"`
	Ticker     string              `bson:"ticker,omitempty"`
	AssetClass string              `bson:"assetClass,omitempty"`
	Sectors    []*BreakdownModel   `bson:"sectors,omitempty"`
}

// BreakdownModel struct
type BreakdownModel struct {
	SectorCode  string  `bson:"sectorCode,omitempty"`
	SectorName  string  `bson:"sectorName,omitempty"`
	FundPercent float64 `bson:"fundPercent,omitempty"`
}

// NewFundBreakdownModel create new fund exposure model
func NewFundBreakdownModel(ctx context.Context, log logger.ContextLog, e *entities.FundBreakdown, schemaVersion string) (*FundBreakdownModel, error) {
	var breakdownModel []*BreakdownModel

	for _, sector := range e.Sectors {
		breakdownModel = append(breakdownModel, &BreakdownModel{
			SectorCode:  sector.SectorCode,
			SectorName:  sector.SectorName,
			FundPercent: sector.FundPercent,
		})
	}

	return &FundBreakdownModel{
		ModifiedAt: time.Now().UTC().Unix(),
		Enabled:    true,
		Deleted:    false,
		Schema:     schemaVersion,
		Source:     consts.DATA_SOURCE,
		Ticker:     e.Ticker,
		AssetClass: e.AssetClass,
		Sectors:    breakdownModel,
	}, nil
}
