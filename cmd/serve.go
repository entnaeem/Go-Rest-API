package cmd

import (
	"fmt"
	"os"
	"sysagent/config"
	"sysagent/infra/db"
	"sysagent/repo"
	"sysagent/rest"
	"sysagent/rest/handlers/product"
	"sysagent/rest/handlers/user"
	middleware "sysagent/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	middlewares := middleware.NewMiddlewares(cnf)

	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
