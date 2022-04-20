package resolve

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"rxdrag.com/entity-engine/model/graph"
	"rxdrag.com/entity-engine/utils"
)

type Loaders struct {
	loaders map[string]*dataloader.Loader
}

func CreateDataLoaders() *Loaders {
	return &Loaders{
		loaders: make(map[string]*dataloader.Loader, 1),
	}
}

func (l *Loaders) GetLoader(association *graph.Association) *dataloader.Loader {
	if l.loaders[association.Path()] == nil {
		l.loaders[association.Path()] = dataloader.NewBatchedLoader(QueryBatchFn(association))
	}
	return l.loaders[association.Path()]
}

func QueryBatchFn(association *graph.Association) dataloader.BatchFunc {
	return func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		defer utils.PrintErrorStack()
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
