package webserver

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rmukhamet/core_test_task/internal/apperrors"
)

func (ws *WebServer) create(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.AcceptsCharsets("utf-8")

	retailerRequest := new(RetailerCreateRequest)

	err := c.BodyParser(retailerRequest)
	if err != nil {
		log.Print(fmt.Errorf("failed parse json: %+v, with error: %w", string(c.BodyRaw()), err))
		c.Status(fiber.StatusBadRequest)
	}

	retailer := retailerRequest.ToDTO()

	// TODO Get Actor

	err = retailer.Validate()
	if err != nil {
		log.Print(fmt.Errorf("retailer %v not valid: %w", retailer, err))
		c.Status(fiber.StatusBadRequest).SendString(apperrors.ErrorRetailerNotValid.Error())
	}

	err = ws.retailerController.Create(c.Context(), retailer)
	if err != nil {
		log.Print(fmt.Errorf("failed queue with error: %w", err))
		c.Status(fiber.StatusInternalServerError)
	}

	c.Status(fiber.StatusCreated)

	return nil
}

func (ws *WebServer) update(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.AcceptsCharsets("utf-8")

	retailerRequest := new(RetailerUpdateRequest)

	err := c.BodyParser(retailerRequest)
	if err != nil {
		log.Print(fmt.Errorf("failed parse json: %+v, with error: %w", string(c.BodyRaw()), err))
		c.Status(fiber.StatusBadRequest)
	}

	retailer := retailerRequest.ToDTO()

	// TODO Get Actor from jwt

	err = retailer.Validate()
	if err != nil {
		log.Print(fmt.Errorf("retailer %v not valid: %w", retailer, err))
		c.Status(fiber.StatusBadRequest).SendString(apperrors.ErrorRetailerNotValid.Error())
	}

	err = ws.retailerController.Create(c.Context(), retailer)
	if err != nil {
		log.Print(fmt.Errorf("failed queue with error: %w", err))
		c.Status(fiber.StatusInternalServerError)
	}

	c.Status(fiber.StatusCreated)

	return nil
}
