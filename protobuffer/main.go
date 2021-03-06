package main

import (
	"context"
	"fmt"

	"github.com/hiago-balbino/hgz-protobuffer-comparison/protobuffer/internal/cache"
	"github.com/hiago-balbino/hgz-protobuffer-comparison/protobuffer/internal/model"
)

func main() {
	ctx := context.Background()
	cacheRepository := cache.NewCacheRepository("localhost:6379")

	insert(ctx, cacheRepository)
	fetch(ctx, cacheRepository)
}

func insert(ctx context.Context, cacheRepository cache.Repository) {
	for id := int32(1); id <= 10000; id++ {
		addressBook := model.CreateAddressBook(id)

		err := cacheRepository.Set(ctx, addressBook)
		if err != nil {
			fmt.Printf("error to insert data into cache: %s", err.Error())
			continue
		}
	}

	fmt.Println("finished insertion")
}

func fetch(ctx context.Context, cacheRepository cache.Repository) {
	for id := int32(1); id <= 10000; id++ {
		addressBook, err := cacheRepository.Get(ctx, id)
		if err != nil {
			fmt.Printf("error to fetch data from cache: %s", err.Error())
			continue
		}

		fmt.Println(addressBook)
	}
	fmt.Println("finished fetch")
}
