package main

import "context"

func main() {
	ctx := context.Background()
	go GetFavoriteActionConsumer(ctx)

	run := true

	for run {

	}
}
