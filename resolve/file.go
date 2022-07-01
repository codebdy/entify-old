package resolve

import (
	"github.com/graphql-go/graphql"
	"rxdrag.com/entify/config"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/storage"
	"rxdrag.com/entify/utils"
)

func FileUrlResolve(p graphql.ResolveParams) (interface{}, error) {
	defer utils.PrintErrorStack()
	if p.Source != nil {
		fileInfo := p.Source.(storage.FileInfo)
		if config.Storage() == consts.LOCAL {
			return consts.UPLOAD_PRIFIX + "/" + fileInfo.Path, nil
		} else {
			return fileInfo.Path, nil
		}
	}
	return nil, nil
}
