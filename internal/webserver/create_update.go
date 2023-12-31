package webserver

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rmukhamet/core_test_task/internal/apperrors"
)

func (ws *WebServer) create(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.AcceptsCharsets("utf-8")

	retailerRequest := new(RetailerCreateRequest)

	err := c.BodyParser(retailerRequest)
	if err != nil {
		log.Print(fmt.Errorf("failed parse json: %+v, with error: %w", string(c.BodyRaw()), err).Error())
		c.Status(fiber.StatusBadRequest)

		return err
	}

	retailer := retailerRequest.ToDTO()

	// TODO Get Actor

	err = retailer.Validate()
	if err != nil {
		log.Print(fmt.Errorf("retailer %v not valid: %w", retailer, err).Error())
		c.Status(fiber.StatusBadRequest).SendString(apperrors.ErrorRetailerNotValid.Error())

		return err
	}

	err = ws.retailerController.Create(c.Context(), retailer)
	if err != nil {
		log.Print(fmt.Errorf("failed create retailer error: %w", err))
		c.Status(fiber.StatusInternalServerError)

		return err
	}

	c.Status(fiber.StatusAccepted)

	time.Sleep(2 * time.Second) // receive response from storage about created success
	log.Print("change status to 201")
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

		return err
	}

	retailer := retailerRequest.ToDTO()

	// TODO Get Actor from jwt

	err = retailer.Validate()
	if err != nil {
		log.Print(fmt.Errorf("retailer %v not valid: %w", retailer, err))
		c.Status(fiber.StatusBadRequest).SendString(apperrors.ErrorRetailerNotValid.Error())

		return err
	}

	err = ws.retailerController.Update(c.Context(), retailer)
	if err != nil {
		log.Print(fmt.Errorf("failed update retailer with error: %w", err).Error())
		c.Status(fiber.StatusInternalServerError)

		return err
	}

	c.Status(fiber.StatusAccepted)

	return nil
}
