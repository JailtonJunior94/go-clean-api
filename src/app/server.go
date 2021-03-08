package server

import (
	"fmt"
	"log"

	"github.com/jailtonjunior94/go-clean-api/src/app/config"
)

func Run() {
	app := config.App()

	fmt.Printf("🚀 API is running on http://localhost:%d", 3000)
	log.Fatal(app.Listen(fmt.Sprintf(":%v", 3000)))
}
