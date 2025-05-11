package handlers

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/srohatgi01/tpp-helper/data"
	service "github.com/srohatgi01/tpp-helper/service"
)

func EstimateHandler(c *fiber.Ctx) error {
	// Parse form values
	widthStr := c.FormValue("width")
	heightStr := c.FormValue("height")
	paper := c.FormValue("paper")
	lamination := c.FormValue("lamination")

	// convert string to int
	width, err1 := strconv.Atoi(widthStr)
	height, err2 := strconv.Atoi(heightStr)

	if err1 != nil || err2 != nil {
		return renderError(c, errors.New("invalid input"))
	}

	photoService := service.PhotoService{
		Width:      width,
		Height:     height,
		Paper:      paper,
		Lamination: lamination == "on",
	}

	total, err := photoService.CalculateTotal()
	if err != nil {
		return renderError(c, err)
	}

	return c.Render("index", fiber.Map{
		"Width":      width,
		"Height":     height,
		"Paper":      paper,
		"Papers":     data.GetPaperList(),
		"Lamination": lamination,
		"Result": fiber.Map{
			"Total":     total,
			"Area":      photoService.GetArea(),
			"Paper":     data.GetPaperByName(paper),
			"Breakdown": photoService.GetBreakdown(),
		},
	})
}

func renderError(c *fiber.Ctx, err error) error {
	return c.Render("index", fiber.Map{
		"Result": fiber.Map{"Error": err.Error()},
	})
}
