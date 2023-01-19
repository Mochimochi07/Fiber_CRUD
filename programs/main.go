package main

import (
	"github.com/gofiber/fiber"
)

type Song struct {
	Title  string `json:"title"`
	Album  string `json:"album"`
	Length int    `json:"length"`
}

var songs []Song

func init() {
	songs = []Song{
		{Title: "Dreams Don't Turn to Dust", Album: "Maybe I'm Dreaming", Length: 216},
		{Title: "The Technicolor Phase", Album: "Maybe I'm Dreaming", Length: 210},
		{Title: "Rainbow Veins", Album: "Maybe I'm Dreaming", Length: 237},
	}
}

func main() {
	app := fiber.New()

	app.Post("/songs", func(c *fiber.Ctx) {
		var song Song
		if err := c.BodyParser(&song); err != nil {
			c.Status(400).Send(err)
			return
		}
		songs = append(songs, song)
		c.JSON(song)
	})

	
	app.Get("/songs", func(c *fiber.Ctx) {
		c.JSON(songs)
	})

	app.Get("/songs/:title", func(c *fiber.Ctx) {
		title := c.Params("title")
		for _, song := range songs {
			if song.Title == title {
				c.JSON(song)
				return
			}
		}
		c.Status(404).Send("song not found")
	})

	
	app.Put("/songs/:title", func(c *fiber.Ctx) {
		title := c.Params("title")
		var newSong Song
		if err := c.BodyParser(&newSong); err != nil {
			c.Status(400).Send(err)
			return
		}
		for i, song := range songs {
			if song.Title == title {
				songs[i] = newSong
				c.JSON(newSong)
				return
			}
		}
		c.Status(404).Send("song not found")
	})


	app.Delete("/songs/:title", func(c *fiber.Ctx) {
		title := c.Params("title")
		for i, song := range songs {
			if song.Title == title {
				songs = append(songs[:i], songs[i+1:]...)
				c.Send("song deleted")
				return
			}
		}
		c.Status(404).Send("song not found")
	})

	app.Listen(3000)
}
