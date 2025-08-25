package auction

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/joaolima7/leilao-goexpert/internal/entity/auction_entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestAutomaticAuctionClosing(t *testing.T) {
	originalInterval := os.Getenv("AUCTION_INTERVAL")
	os.Setenv("AUCTION_INTERVAL", "2s")
	defer os.Setenv("AUCTION_INTERVAL", originalInterval)

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	database := client.Database("test_auctions")
	defer database.Collection("auctions").Drop(ctx)

	repo := NewAuctionRepository(database)

	auction, err := auction_entity.CreateAuction(
		"Test Product",
		"Test Category",
		"Test Description for automatic closing",
		auction_entity.New,
	)
	assert.Nil(t, err)

	err = repo.CreateAuction(ctx, auction)
	assert.Nil(t, err)

	time.Sleep(3 * time.Second)

	var updatedAuction AuctionEntityMongo
	filter := bson.M{"_id": auction.Id}
	err = repo.Collection.FindOne(ctx, filter).Decode(&updatedAuction)
	assert.Nil(t, err)

	assert.Equal(t, auction_entity.Completed, updatedAuction.Status)
}
