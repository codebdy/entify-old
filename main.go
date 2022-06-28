package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"rxdrag.com/entify/authentication"
	"rxdrag.com/entify/config"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/db"
	"rxdrag.com/entify/handler"
	"rxdrag.com/entify/repository"
	"rxdrag.com/entify/resolve"
	"rxdrag.com/entify/schema"
)

const PORT = 4000

func checkParams() {
	dbConfig := config.GetDbConfig()
	if dbConfig.Driver == "" ||
		dbConfig.Host == "" ||
		dbConfig.Database == "" ||
		dbConfig.User == "" ||
		dbConfig.Port == "" ||
		dbConfig.Password == "" {
		panic("Params is not enough, please set")
	}
}

func checkMetaInstall() {
	if !repository.IsEntityExists(consts.META_ENTITY_NAME) {
		repository.Install()
	}
}

func checkMediaInstall() {
	if !repository.IsEntityExists(consts.MEDIA_ENTITY_NAME) {
		resolve.InstallMedia()
	}
}

func main() {
	defer db.Close()
	checkParams()
	checkMetaInstall()
	repository.InitGlobalModel()
	repository.LoadModel()
	if config.Storage() != "" {
		checkMediaInstall()
	}
	if config.AuthUrl() == "" && !repository.IsEntityExists(consts.META_USER) {
		schema.InitAuthInstallSchema()
	} else {
		schema.InitSchema()
	}

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
