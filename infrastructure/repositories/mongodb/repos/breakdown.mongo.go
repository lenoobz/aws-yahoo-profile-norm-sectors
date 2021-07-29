package repos

import (
	"context"
	"fmt"
	"time"

	logger "github.com/lenoobz/aws-lambda-logger"
	"github.com/lenoobz/aws-yahoo-profile-norm-sectors/config"
	"github.com/lenoobz/aws-yahoo-profile-norm-sectors/consts"
	"github.com/lenoobz/aws-yahoo-profile-norm-sectors/entities"
	"github.com/lenoobz/aws-yahoo-profile-norm-sectors/infrastructure/repositories/mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CountryBreakdownMongo struct
type CountryBreakdownMongo struct {
	db     *mongo.Database
	client *mongo.Client
	log    logger.ContextLog
	conf   *config.MongoConfig
}

// NewSectorBreakdownMongo creates new country breakdown mongo repo
func NewSectorBreakdownMongo(db *mongo.Database, log logger.ContextLog, conf *config.MongoConfig) (*CountryBreakdownMongo, error) {
	if db != nil {
		return &CountryBreakdownMongo{
			db:   db,
			log:  log,
			conf: conf,
		}, nil
	}

	// set context with timeout from the config
	// create new context for the query
	ctx, cancel := createContext(context.Background(), conf.TimeoutMS)
	defer cancel()

	// set mongo client options
	clientOptions := options.Client()

	// set min pool size
	if conf.MinPoolSize > 0 {
		clientOptions.SetMinPoolSize(conf.MinPoolSize)
	}

	// set max pool size
	if conf.MaxPoolSize > 0 {
		clientOptions.SetMaxPoolSize(conf.MaxPoolSize)
	}

	// set max idle time ms
	if conf.MaxIdleTimeMS > 0 {
		clientOptions.SetMaxConnIdleTime(time.Duration(conf.MaxIdleTimeMS) * time.Millisecond)
	}

	// construct a connection string from mongo config object
	cxnString := fmt.Sprintf("mongodb+srv://%s:%s@%s", conf.Username, conf.Password, conf.Host)

	// create mongo client by making new connection
	client, err := mongo.Connect(ctx, clientOptions.ApplyURI(cxnString))
	if err != nil {
		return nil, err
	}

	return &CountryBreakdownMongo{
		db:     client.Database(conf.Dbname),
		client: client,
		log:    log,
		conf:   conf,
	}, nil
}

// Close disconnect from database
func (r *CountryBreakdownMongo) Close() {
	ctx := context.Background()
	r.log.Info(ctx, "close mongo client")

	if r.client == nil {
		return
	}

	if err := r.client.Disconnect(ctx); err != nil {
		r.log.Error(ctx, "disconnect mongo failed", "error", err)
	}
}

///////////////////////////////////////////////////////////////////////////////
// Implement interface
///////////////////////////////////////////////////////////////////////////////

// InsertAssetSectorBreakdown adds new asset sector breakdown
func (r *CountryBreakdownMongo) InsertAssetSectorBreakdown(ctx context.Context, breakdown *entities.FundBreakdown) error {
	// create new context for the query
	ctx, cancel := createContext(ctx, r.conf.TimeoutMS)
	defer cancel()

	sectorBreakdownModel, err := models.NewFundBreakdownModel(ctx, r.log, breakdown, r.conf.SchemaVersion)
	if err != nil {
		r.log.Error(ctx, "create model failed", "error", err)
		return err
	}

	// what collection we are going to use
	colname, ok := r.conf.Colnames[consts.ASSET_SECTORS_COLLECTION]
	if !ok {
		r.log.Error(ctx, "cannot find collection name")
	}
	col := r.db.Collection(colname)

	filter := bson.D{{
		Key:   "ticker",
		Value: sectorBreakdownModel.Ticker,
	}}

	update := bson.D{
		{
			Key:   "$set",
			Value: sectorBreakdownModel,
		},
		{
			Key: "$setOnInsert",
			Value: bson.D{{
				Key:   "createdAt",
				Value: time.Now().UTC().Unix(),
			}},
		},
	}

	opts := options.Update().SetUpsert(true)

	_, err = col.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		r.log.Error(ctx, "update one failed", "error", err)
		return err
	}

	return nil
}
