package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func main() {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// Get command-line arguments
	projectId, ok := os.LookupEnv("PROJECT_ID")
	if !ok {
		panic("PROJECT_ID is not set")
	}

	itr := client.Buckets(ctx, projectId)
	for {
		b, e := itr.Next()
		if e == iterator.Done {
			break
		}
		if e != nil {
			panic(e)
		}
		fmt.Println(b.Name)
	}
}
