package main

import (
	"context"
	"log"

	"github.com/google/uuid"
	corid "github.com/lenoobz/aws-lambda-corid"
	logger "github.com/lenoobz/aws-lambda-logger"
	"github.com/lenoobz/aws-yahoo-profile-norm-sectors/config"
	"github.com/lenoobz/aws-yahoo-profile-norm-sectors/infrastructure/repositories/mongodb/repos"
	"github.com/lenoobz/aws-yahoo-profile-norm-sectors/usecase/breakdown"
	"github.com/lenoobz/aws-yahoo-profile-norm-sectors/usecase/yahoo"
)

func main() {
	appConf := config.AppConf

	// create new logger
	zap, err := logger.NewZapLogger()
	if err != nil {
		log.Fatal("create app logger failed")
	}
	defer zap.Close()

	// create new repository
	yahooProfileRepo, err := repos.NewYahooProfileMongo(nil, zap, &appConf.Mongo)
	if err != nil {
		log.Fatalf("create Yahoo profile mongo failed: %v\n", err)
	}
	defer yahooProfileRepo.Close()

	// create new repository
	sectorBreakdownRepo, err := repos.NewSectorBreakdownMongo(nil, zap, &appConf.Mongo)
	if err != nil {
		log.Fatal("create sector breakdown mongo failed")
	}
	defer sectorBreakdownRepo.Close()

	// create new service
	yahooService := yahoo.NewService(yahooProfileRepo, zap)
	breakdownService := breakdown.NewService(sectorBreakdownRepo, *yahooService, zap)

	// try correlation context
	id, _ := uuid.NewRandom()
	ctx := corid.NewContext(context.Background(), id)

	// try correlation context
	if err := breakdownService.AddAssetSectorBreakdown(ctx); err != nil {
		log.Fatal("add asset sector breakdown failed")
	}
}
