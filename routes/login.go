package routes

import (
	"example/server/utils"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func SignInRoutes(app fiber.Router) {
	app.Use(func(c *fiber.Ctx) error {

		log.Println("this is the signup route")
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		now := utils.GetCurrentDate()
		return c.SendString(fmt.Sprintf("Hello world server from fiber\nDate: %d", now))
	})

	app.Post("/upload", func(c *fiber.Ctx) error {
		form, err := c.MultipartForm()

		if err != nil {
			return err
		}

		files := form.File["file"]

		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			fileName := fmt.Sprintf("%d-%s", utils.GetCurrentDate(), strings.ReplaceAll(file.Filename, " ", "-"))

			// Save the files to disk:
			err := c.SaveFile(file, fmt.Sprintf("./%s", fileName))
			if err != nil {
				return err
			}
		}

		return c.JSON(utils.SendSucessResponse("successfully uploaded file"))

	})

}
