package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssibrahimbas/claim-auth.go/src/i18n"
	"github.com/ssibrahimbas/claim-auth.go/src/result"
	"go.mongodb.org/mongo-driver/mongo"
)

func errorHandler(cfg *Config, i *i18n.I18n) func(c *fiber.Ctx, err error) error {
	return func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*result.Result); ok {
			return c.Status(e.Code).JSON(e)
		}
		if e, ok := err.(*result.DataResult); ok {
			return c.Status(e.Code).JSON(e)
		}
		if err == mongo.ErrNoDocuments {
			l, a := i.GetLanguagesInContext(c)
			return c.Status(fiber.StatusNotFound).JSON(result.Error(i.Translate(cfg.NFMsgKey, l, a), fiber.StatusNotFound))
		}
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}
		if cfg.DfMsgKey != "" {
			l, a := i.GetLanguagesInContext(c)
			return c.Status(code).JSON(result.Error(i.Translate(cfg.DfMsgKey, l, a), code))
		}
		err = c.Status(code).JSON(result.Error(err.Error(), code))
		return err
	}
}
