package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssibrahimbas/claim-auth.go/src/i18n"
	"github.com/ssibrahimbas/claim-auth.go/src/result"
	"github.com/ssibrahimbas/claim-auth.go/src/validator"
)

func (h *Client) ParseBody(c *fiber.Ctx, v *validator.Validator, i *i18n.I18n, d interface{}) {
	l, a := i.GetLanguagesInContext(c)
	if err := c.BodyParser(d); err != nil {
		panic(result.Error(i.Translate("error_invalid_request_body", l, a), fiber.StatusBadRequest))
	}
	h.validateStruct(d, v, i, l, a)
}

func (h *Client) ParseQuery(c *fiber.Ctx, v *validator.Validator, i *i18n.I18n, d interface{}) {
	l, a := i.GetLanguagesInContext(c)
	if err := c.QueryParser(d); err != nil {
		panic(result.Error(i.Translate("error_invalid_request_query", l, a), fiber.StatusBadRequest))
	}
	h.validateStruct(d, v, i, l, a)
}

func (h *Client) ParseParams(c *fiber.Ctx, v *validator.Validator, i *i18n.I18n, d interface{}) {
	l, a := i.GetLanguagesInContext(c)
	if err := c.ParamsParser(d); err != nil {
		panic(result.Error(i.Translate("error_invalid_request_params", l, a), fiber.StatusBadRequest))
	}
	h.validateStruct(d, v, i, l, a)
}

func (h *Client) validateStruct(d interface{}, v *validator.Validator, i *i18n.I18n, l, a string) {
	if errors := v.ValidateStruct(d, l, a); len(errors) > 0 {
		panic(result.ErrorData(i.Translate("error_validation_failed", l, a), errors, fiber.StatusBadRequest))
	}
}
