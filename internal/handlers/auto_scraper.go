package handlers

import (
	"net/http"

	"connectrpc.com/connect"
	autoscraperv1 "github.com/diegoafg1009/auto-radar-scraping-microservice/pkg/genproto/autoscraper/v1"
	"github.com/diegoafg1009/auto-radar-scraping-microservice/pkg/genproto/autoscraper/v1/autoscraperv1connect"
	"github.com/labstack/echo/v4"
)

type AutoScraperHandler struct {
	autoScraperService autoscraperv1connect.AutoScraperServiceClient
}

func NewAutoScraperHandler(autoScraperService autoscraperv1connect.AutoScraperServiceClient) *AutoScraperHandler {
	return &AutoScraperHandler{
		autoScraperService: autoScraperService,
	}
}

func (h *AutoScraperHandler) FindByFilter(c echo.Context) error {
	var req autoscraperv1.FindByFilterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp, err := h.autoScraperService.FindByFilter(c.Request().Context(), connect.NewRequest(&req))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp.Msg)
}
