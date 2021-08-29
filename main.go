package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kurojs/ipinfo/ent"
	"github.com/kurojs/ipinfo/internal/services"
	"github.com/kurojs/ipinfo/internal/storages"
)

func main() {
	app := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: true,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://agile-anchorage-23778.herokuapp.com",
	}))

	fmt.Println("Creating Geo Service...")
	geoService, err := services.NewGeoService()
	if err != nil {
		log.Fatalf("Create Geo service failed, %s", err)
	}
	defer func() {
		geoService.Close()
	}()

	fmt.Println("Creating Storage Service...")
	storage, err := storages.NewStorage("ec2-184-72-228-128.compute-1.amazonaws.com", "5432", "yjmfpttbtkqcdl", "d474s08g9k5mk0", "dc2c5ac844b10bee1f72fbf23d600bfe4864503d5c29619e6730de735b0e425d")
	if err != nil {
		log.Fatalf("Create Storage service failed, %s", err)
	}
	defer func() {
		storage.Close()
	}()

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("✋ Hello there, IP Info is at your service")
	})

	app.Get("api/info", func(c *fiber.Ctx) error {
		geoInfo, err := geoService.GetIPInfo(c.IPs()[0])
		if err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}

		storage.CreateHistory(c.Context(), &ent.History{
			IP:        geoInfo.IP,
			City:      geoInfo.City,
			Region:    geoInfo.Region,
			LoginTime: c.Context().Time(),
		})

		return c.JSON(geoInfo)
	})

	// Return visited IP historical
	app.Get("api/history", func(c *fiber.Ctx) error {
		histories, err := storage.GetHistories(c.Context())
		if err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}

		return c.JSON(histories)
	})

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
