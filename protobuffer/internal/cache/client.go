package cache

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/hiago-balbino/hgz-protobuffer-comparison/protobuffer/internal/model"
	"google.golang.org/protobuf/proto"
)

const keyPattern = "proto#book#{%d}"

// Repository is a struct that contains cache client to execute commands
type Repository struct {
	client *redis.Client
}

// NewCacheRepository is a constructor to create a new instance of CacheRepository
func NewCacheRepository(addr string) Repository {
	return Repository{
		client: redis.NewClient(&redis.Options{
			Addr: addr,
			DB:   0,
		}),
	}
}

// Set is a function to set values into cache
func (c Repository) Set(ctx context.Context, addressBook *model.AddressBook) error {
	key := fmt.Sprintf(keyPattern, addressBook.GetPeople()[0].GetId())

	addressBookBytes, err := proto.Marshal(addressBook)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, key, addressBookBytes, 0).Err()
}

// Get is a function to get values from cache
func (c Repository) Get(ctx context.Context, id int32) (*model.AddressBook, error) {
	key := fmt.Sprintf(keyPattern, id)
	data, err := c.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	addressBook := model.AddressBook{}
	err = proto.Unmarshal(data, &addressBook)
	if err != nil {
		return nil, err
	}

	return &addressBook, nil
}
