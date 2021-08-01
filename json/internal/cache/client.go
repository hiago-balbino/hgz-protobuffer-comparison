package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/hiago-balbino/hgz-protobuffer-comparison/json/internal/model"
)

const keyPattern = "json#book#{%d}"

// Repository is a struct that contains cache client to execute commands
type Repository struct {
	client *redis.Client
}

// NewCacheRepository is a constructor to create a new instance of CacheRepository
func NewCacheRepository(addr string) Repository {
	return Repository{
		client: redis.NewClient(&redis.Options{
			Addr: addr,
			DB:   1,
		}),
	}
}

// Set is a function to set values into cache
func (c Repository) Set(ctx context.Context, addressBook model.AddressBook) error {
	key := fmt.Sprintf(keyPattern, addressBook.People[0].Id)

	addressBookBytes, err := json.Marshal(addressBook)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, key, addressBookBytes, 0).Err()
}

// Get is a function to get values from cache
func (c Repository) Get(ctx context.Context, id int32) (model.AddressBook, error) {
	key := fmt.Sprintf(keyPattern, id)
	data, err := c.client.Get(ctx, key).Bytes()
	if err != nil {
		return model.AddressBook{}, err
	}

	addressBook := model.AddressBook{}
	err = json.Unmarshal(data, &addressBook)
	if err != nil {
		return model.AddressBook{}, err
	}

	return addressBook, nil
}
