package cache

import (
	lru "github.com/hashicorp/golang-lru"
)

var CommandCache *lru.ARCCache

func init() {
	var err error
	CommandCache, err = lru.NewARC(1 << 20)
	if err != nil {
		panic(err)
	}
}
