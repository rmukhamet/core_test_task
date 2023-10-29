package webserver

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func (ws *WebServer) getRetailerByID(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.AcceptsCharsets("utf-8")

	retailerID := c.Params("id")

	retailer, err := ws.retailerController.GetRetailerByID(c.Context(), retailerID)
	if err != nil {
		log.Print(fmt.Errorf("failed get retailer by ID %s with error: %w", retailerID, err))
		c.Status(fiber.StatusInternalServerError)

		return err
	}

	response := NewRetailerResponse(retailer)
	c.Status(fiber.StatusOK).JSON(response)

	return nil
}

func (ws *WebServer) getRetailerList(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.AcceptsCharsets("utf-8")

	retailers, err := ws.retailerController.GetRetailerList(c.Context())
	if err != nil {
		log.Print(fmt.Errorf("failed get retailer list: %w", err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	response := NewRetalerListResponse(retailers)
	c.Status(fiber.StatusOK).JSON(response)

	return nil
}

func (ws *WebServer) delete(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.AcceptsCharsets("utf-8")

	retailerID := c.Params("id")

	err := ws.retailerController.DeleteRetailer(c.Context(), retailerID)
	if err != nil {
		log.Print(fmt.Errorf("failed get retailer list: %w", err))
		c.Status(fiber.StatusInternalServerError)

		return err
	}

	c.Status(fiber.StatusOK)

	return nil
}

func (ws *WebServer) getRetailerVersionList(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.AcceptsCharsets("utf-8")

	retailerID := c.Params("id")

	versions, err := ws.retailerController.GetRetailerVersionList(c.Context(), retailerID)
	if err != nil {
		log.Print(fmt.Errorf("failed get retailer version list: %w", err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	response := NewRetalerVersionListResponse(versions)
	c.Status(fiber.StatusOK).JSON(response)

	return nil
}

func (ws *WebServer) getRetailerVersion(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.AcceptsCharsets("utf-8")

	retailerID := c.Params("id")
	versionID, err := c.ParamsInt("version_id")
	if err != nil {
		log.Printf("failed parse version_id with error: %s\n", err.Error())
		c.Status(fiber.StatusBadRequest).SendString("version_id should be number")
		return err
	}

	retailer, err := ws.retailerController.GetRetailerVersion(c.Context(), retailerID, versionID)
	if err != nil {
		log.Print(fmt.Errorf("failed get retailer version: %w", err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	response := NewRetailerResponse(retailer)

	c.Status(fiber.StatusOK).JSON(response)

	return nil
}
func (ws *WebServer) deleteRetailerVersion(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.AcceptsCharsets("utf-8")

	retailerID := c.Params("id")
	versionID, err := c.ParamsInt("version_id")
	if err != nil {
		log.Printf("failed parse version_id with error: %s\n", err.Error())
		c.Status(fiber.StatusBadRequest).SendString("version_id should be number")
		return err
	}

	err = ws.retailerController.DeleteRetailerVersion(c.Context(), retailerID, versionID)
	if err != nil {
		log.Print(fmt.Errorf("failed delete retailer: %s version: %d, error: %w", retailerID, versionID, err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	c.Status(fiber.StatusOK)

	return nil
}

func (ws *WebServer) login(c *fiber.Ctx) error {
	// todo make auth service layer
	creds := make(map[string]string)

	err := c.BodyParser(&creds)
	if err != nil {
		log.Print(fmt.Errorf("failed parse json: %+v, with error: %w", string(c.BodyRaw()), err))
		c.Status(fiber.StatusBadRequest)
		return err
	}

	login, ok := creds["login"]
	if !ok {
		c.Status(fiber.StatusBadRequest)
		return nil
	}

	u := &url.URL{
		Scheme: "http",
		Host:   ws.AuthURL,
		Path:   "/generate",
	}

	q := u.Query()
	q.Set("login", login)
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(c.Context(), http.MethodGet, u.String(), nil)
	if err != nil {
		log.Print(fmt.Errorf("failed create request error: %w", err))
		c.Status(fiber.StatusInternalServerError)

		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(fmt.Errorf("failed request error: %w", err))
		c.Status(fiber.StatusInternalServerError)

		return err
	}

	if resp.StatusCode != http.StatusOK {
		log.Print(fmt.Errorf("failed authorize user: %s with error: %w", login, err).Error())
		c.Status(fiber.StatusUnauthorized)

		return nil
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)

		return err
	}

	c.Status(fiber.StatusOK).SendString(string(b))

	return nil
}
