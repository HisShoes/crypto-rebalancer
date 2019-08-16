package mongodb

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/ksuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/hisshoes/crypto-rebalancer/pkg/portfolio"
)

// Storage - track portfolios and assets in memory
type Storage struct {
	db         *mongo.Client
	portfolios *mongo.Collection
	assets     *mongo.Collection
}

// NewRepository returns an object implementing portfolio repo interface storing in mongodb
func NewRepository(url string) Storage {
	// connect to the db and create a client
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal("Failed to create mongodb client")
	}

	// setup the collections required for the repository implementation
	portfolios := client.Database("cryptos").Collection("portfolios")
	assets := client.Database("cryptos").Collection("assets")

	return Storage{
		db:         client,
		portfolios: portfolios,
		assets:     assets,
	}
}

// GetAssetPrice get the price related to an asset
func (s Storage) GetAssetPrice(n string) (float64, error) {
	ctx, c := context.WithTimeout(context.Background(), 30*time.Second)
	defer c()

	var result portfolio.Asset
	filter := bson.M{"name": n}

	err := s.assets.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return 0, portfolio.ErrMissing
	}

	return result.Price, nil
}

// Portfolio return a portfolio relating to an id
func (s Storage) Portfolio(id string) (portfolio.Portfolio, error) {

	return portfolio.Portfolio{}, portfolio.ErrMissing
}

// ListPortfolios return all the portfolios
func (s Storage) ListPortfolios() ([]portfolio.Portfolio, error) {
	return []portfolio.Portfolio{}, nil
}

// CreatePortfolio create a portfolio and append to the slice
func (s Storage) CreatePortfolio(p portfolio.Portfolio) (string, error) {
	// setup non-user set values
	p.ID = generateID()
	p.UpdateTime = time.Now()

	ctx, c := context.WithTimeout(context.Background(), 30*time.Second)
	defer c()

	_, err := s.portfolios.InsertOne(ctx, p)
	if err != nil {
		return "", err
	}

	return p.ID, nil
}

// UpdatePortfolio update a specific portfolio
func (s Storage) UpdatePortfolio(p portfolio.Portfolio) error {
	return nil
}

func generateID() string {
	return ksuid.New().String()
}
