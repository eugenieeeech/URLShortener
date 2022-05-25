package server

import (
	"URLShortner/model"
	"URLShortner/utils"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func redirect(c *fiber.Ctx) error {
	shortUrl := c.Params("redirect")
	short, err := model.FindByURLshortUrl(shortUrl)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not find short in DB " + err.Error(),
		})
	}
	// grab any stats you want...
	short.Clicked += 1
	err = model.UpdateURLshort(short)
	if err != nil {
		fmt.Printf("error updating: %v\n", err)
	}

	return c.Redirect(short.Redirect, fiber.StatusTemporaryRedirect)
}

func getAllShorts(c *fiber.Ctx) error {
	urlshorts, err := model.GetAllurlshorts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all shorten links " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(urlshorts)
}

func getShort(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not parse id " + err.Error(),
		})
	}

	short, err := model.Geturlshorts(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not retreive goly from db " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(short)
}

func createShort(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var short model.URLshort
	err := c.BodyParser(&short)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing JSON " + err.Error(),
		})
	}

	if short.Random {
		short.URLshort = utils.RandomURL(8)
	}

	err = model.CreateURLshort(short)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not create short in db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(short)

}

func updateShort(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var short model.URLshort

	err := c.BodyParser(&short)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not parse json " + err.Error(),
		})
	}

	err = model.UpdateURLshort(short)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not update short link in DB " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(short)
}

func deleteShort(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not parse id from url " + err.Error(),
		})
	}

	err = model.DeleteURLshort(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not delete from db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "short deleted.",
	})
}

func SetupAndListen() {

	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/r/:redirect", redirect)

	router.Get("/short", getAllShorts)
	router.Get("/short/:id", getShort)
	router.Post("/short", createShort)
	router.Patch("/short", updateShort)
	router.Delete("/short/:id", deleteShort)

	router.Listen(":3000")

}
