package http

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/ssibrahimbas/claim-auth.go/src/i18n"
)

type Client struct {
	App *fiber.App
}

type Config struct {
	NFMsgKey string
	DfMsgKey string
}

var defaultConfig = Config{
	// not found message key
	NFMsgKey: "not_found",
	// default message key
	DfMsgKey: "",
}

func New(i18n *i18n.I18n, config ...Config) *Client {
	var cfg = defaultConfig
	if len(config) > 0 {
		cfg = config[0]
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler(&cfg, i18n),
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})
	app.Use(recover.New())
	return &Client{
		App: app,
	}
}

func (h *Client) Listen(p string) error {
	return h.App.Listen(p)
}
