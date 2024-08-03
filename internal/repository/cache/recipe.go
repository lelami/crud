package cache

import (
	"context"
	"crud/internal/domain"
	"errors"
	"sync"
)

type RecipeCache struct {
	pool map[string]*domain.Recipe
	mtx  sync.RWMutex
}

const RecipeDumpFileName = "recipes.json"

func RecipeCacheInit(ctx context.Context, wg *sync.WaitGroup) (*RecipeCache, error) {
	var c RecipeCache
	c.pool = make(map[string]*domain.Recipe)

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		makeDump(RecipeDumpFileName, c.pool)
	}()

	if err := loadFromDump(RecipeDumpFileName, &c.pool); err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *RecipeCache) Get(id string) (*domain.Recipe, error) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	if val, ok := c.pool[id]; ok {
		return val, nil
	}
	return nil, errors.New("recipe not found")
}

func (c *RecipeCache) Set(id string, recipe *domain.Recipe) error {

	c.mtx.Lock()
	c.pool[id] = recipe
	c.mtx.Unlock()

	return nil
}
func (c *RecipeCache) Delete(id string) error {

	c.mtx.Lock()
	delete(c.pool, id)
	c.mtx.Unlock()

	return nil
}
