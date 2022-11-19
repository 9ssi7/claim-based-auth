package main

import (
	"github.com/ssibrahimbas/claim-auth.go/src/app"
)

func main() {
	app.New().Init().Serve()
}
