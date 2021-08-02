package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	logger "github.com/lenoobz/aws-lambda-logger"
	"github.com/lenoobz/aws-yahoo-profile-norm-sectors/config"
	"github.com/lenoobz/aws-yahoo-profile-norm-sectors/infrastructure/repositories/mongodb/repos"
	"github.com/lenoobz/aws-yahoo-profile-norm-sectors/usecase/breakdown"
	"github.com/lenoobz/aws-yahoo-profile-norm-sectors/usecase/yahoo"
)

type YahooProfileSectorRequest struct {
	Tickers []string `json:"tickers"`
}

func main() {
	lambda.Start(lambdaHandler)
}

func lambdaHandler(ctx context.Context, req YahooProfileSectorRequest) {
	log.Println("lambda handler is called")

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
	if err := breakdownService.AddAssetSectorBreakdownByTickers(ctx, req.Tickers); err != nil {
		log.Fatal("add asset sector breakdown failed")
	}
}
