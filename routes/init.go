package routes

import (
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"pictureframe/views"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return views.Main().Render(c.Request().Context(), c.Response().Writer)
	})
	e.GET("/random.txt/image", func(c echo.Context) error {
		var folder = "./static/images/random.txt"

		// Read all files in the folder
		files, err := os.ReadDir(folder)
		if err != nil {
			println("lol")
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to read directory",
			})
		}

		// Filter for image files (optional, based on extensions)
		var images []string
		for _, file := range files {
			if !file.IsDir() {
				images = append(images, file.Name())
			}
		}

		if len(images) == 0 {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "No images found",
			})
		}

		// Pick a random image
		randomImage := images[rand.Intn(len(images))]

		// Serve the image
		imagePath := filepath.Join(folder, randomImage)
		return c.File(imagePath)

	})

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
		if code == http.StatusNotFound {
			c.String(http.StatusNotFound, "404 Not Found")
			return
		}
		e.DefaultHTTPErrorHandler(err, c)
	}
}
