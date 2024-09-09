package cache

import (
	"context"
	"errors"
	"sync"
)

type RecipeOwnerCache struct {
	pool map[string]string
	mtx  sync.RWMutex
}

const RecipeOwnerDumpFileName = "recipe_owner.json"

func RecipeOwnerCacheInit(ctx context.Context, wg *sync.WaitGroup) (*RecipeOwnerCache, error) {
	var c RecipeOwnerCache
	c.pool = make(map[string]string)

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		makeDump(RecipeOwnerDumpFileName, c.pool)
	}()

	if err := loadFromDump(RecipeOwnerDumpFileName, &c.pool); err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *RecipeOwnerCache) Get(id string) (string, error) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	if val, ok := c.pool[id]; ok {
		return val, nil
	}
	return "", errors.New("recipe owner not found")
}

func (c *RecipeOwnerCache) Set(recipeId string, ownerId string) error {
	c.mtx.Lock()
	c.pool[recipeId] = ownerId
	c.mtx.Unlock()

	return nil
}

func (c *RecipeOwnerCache) Delete(id string) error {
	c.mtx.Lock()
	delete(c.pool, id)
	c.mtx.Unlock()
	return nil
}
