package routers

import (
	"net/http"
	"os"

	"github.com/diegoafg1009/auto-radar-api-gateway/internal/handlers"

	"github.com/diegoafg1009/auto-radar-scraping-microservice/pkg/genproto/autoscraper/v1/autoscraperv1connect"
	"github.com/labstack/echo/v4"

	_ "github.com/joho/godotenv/autoload"
)

func AutoScraper(e *echo.Echo) {
	autoScraperUrl := os.Getenv("AUTO_SCRAPER_URL")
	client := autoscraperv1connect.NewAutoScraperServiceClient(http.DefaultClient, autoScraperUrl)
	handler := handlers.NewAutoScraperHandler(client)
	router := e.Group("/auto-scraper")
	router.GET("/find-by-filter", handler.FindByFilter)
}
