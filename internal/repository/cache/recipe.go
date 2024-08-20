package cache

import (
	"context"
	"crud/internal/domain"
	"errors"
	"fmt"
	"slices"
	"sync"
)

const (
	maxBatchSize       = 10
	RecipeDumpFileName = "recipes.json"
)

type RecipeCache struct {
	pool map[string]*domain.Recipe
	ids  []string
	mtx  sync.RWMutex
}

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
	c.ids = make([]string, 0, len(c.pool))
	for id, _ := range c.pool {
		c.ids = append(c.ids, id)
	}
	slices.Sort(c.ids)
	return &c, nil
}
func (c *RecipeCache) GetRecipes(index, batchSize int) (*domain.ResponseRecipes, error) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	if index >= len(c.ids) {
		return nil, errors.New("index out of range recipes")
	}
	if batchSize == 0 || batchSize > maxBatchSize {
		return nil, fmt.Errorf("invalid reques parameters: batch size > %d ", maxBatchSize)
	}
	recipes := &domain.ResponseRecipes{
		Recipes: make([]domain.Recipe, 0, batchSize),
	}
	for _, id := range c.ids {
		recipes.Recipes = append(recipes.Recipes, *c.pool[id])
	}
	return recipes, nil

}
func (c *RecipeCache) Get(id string) (*domain.Recipe, error) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	if val, ok := c.pool[id]; ok {
		return val, nil
	}
	return nil, errors.New("recipe not found")
}
func (c *RecipeCache) Count() int {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return len(c.pool)
}
func (c *RecipeCache) Set(id string, recipe *domain.Recipe) error {

	c.mtx.Lock()
	c.pool[id] = recipe
	c.ids = append(c.ids, id)
	slices.Sort(c.ids)
	c.mtx.Unlock()

	return nil
}
func (c *RecipeCache) Delete(id string) error {

	c.mtx.Lock()
	defer c.mtx.Unlock()
	delete(c.pool, id)
	for i, val := range c.ids {
		if val == id {
			c.ids = append(c.ids[:i], c.ids[i+1:]...)
			break
		}
	}
	return nil
}
