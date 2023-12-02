package main

import (
	"fmt"

	"github.com/muhammedarifp/user/config"
	"github.com/muhammedarifp/user/di"

	_ "github.com/muhammedarifp/user/docs"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	cfg := config.LoadConfig()
	server, err := di.InitializeAPI(cfg)
	//cronejobs.InitCronJobs()

	if err != nil {
		fmt.Println("1. ", err.Error())
	}
	if err := server.Start(cfg.APP_PORT); err != nil {
		fmt.Println("2. ", err.Error())
	}
}
