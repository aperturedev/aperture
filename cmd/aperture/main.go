package main

import (
	"github.com/aperturedev/aperture/internal/app"
	"github.com/aperturedev/aperture/internal/httpecho"
	"github.com/aperturedev/aperture/internal/hyper"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	// nc, err := nats.Connect("nats://nats:4222")
	// must(err)
	//
	// js, err := jetstream.New(nc)
	// must(err)
	//
	// ctx := context.Background()
	//
	// defer nc.Close()
	//
	// must(natsio.Init(ctx, js))

	e := echo.New()

	e.HTTPErrorHandler = hyper.NewErrorHandler()
	e.Renderer = hyper.NewRenderer("views/dashboard")

	e.Static("/public", "public/dashboard")
	e.File("/favicon.ico", "public/dashboard/favicon.ico")

	e.Use(middleware.Logger())

	httpecho.RegisterDashboardServer(e, app.NewProjectionService())

	log.Fatal(e.Start(":8080"))
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
