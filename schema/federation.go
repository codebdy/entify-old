package schema

import "github.com/graphql-go/graphql"

//union _Entity

var _ServiceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "_Service",
		Fields: graphql.Fields{
			"sdl": &graphql.Field{
				Type:        graphql.String,
				Description: "Service SDL",
			},
		},
		Description: "_Service type of federation schema specification",
	},
)
