package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"rxdrag.com/entify/authentication"
	"rxdrag.com/entify/handler"
	"rxdrag.com/entify/resolve"
	"rxdrag.com/entify/schema"
)

const PORT = 4000

func main() {
	h := handler.New(&handler.Config{
		SchemaResolveFn: schema.ResolveSchema,
		Pretty:          true,
		GraphiQLConfig:  &handler.GraphiQLConfig{},
	})

	http.Handle("/graphql",
		authentication.CorsMiddleware(
			authentication.AuthMiddleware(
				resolve.LoadersMiddleware(h),
			),
		),
	)
	http.HandleFunc("/subscriptions", handler.NewFunc(schema.ResolveSchema))
	fmt.Println(fmt.Sprintf("Running a GraphQL API server at http://localhost:%d/graphql", PORT))
	fmt.Println(fmt.Sprintf("Subscriptions endpoint is http://localhost:%d/subscriptions", PORT))
	err2 := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	if err2 != nil {
		fmt.Printf("启动失败:%s", err2)
	}
}
