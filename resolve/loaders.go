package resolve

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"rxdrag.com/entity-engine/model/graph"
)

type Loaders struct {
	loaders map[string]*dataloader.Loader
}

func CreateDataLoaders() *Loaders {
	return &Loaders{
		loaders: make(map[string]*dataloader.Loader, 1),
	}
}

func (l *Loaders) GetLoader(entity *graph.Entity) *dataloader.Loader {
	if l.loaders[entity.Name()] == nil {
		l.loaders[entity.Name()] = dataloader.NewBatchedLoader(QueryBatchFn(entity))
	}
	return l.loaders[entity.Name()]
}

func QueryBatchFn(entity *graph.Entity) dataloader.BatchFunc {
	return func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result
		// handleError := func(err error) []*dataloader.Result {
		// 	var results []*dataloader.Result
		// 	var result dataloader.Result
		// 	result.Error = err
		// 	results = append(results, &result)
		// 	return results
		// }

		return results
	}
}
