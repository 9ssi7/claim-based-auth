package required_auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssibrahimbas/claim-auth.go/src/i18n"
	"github.com/ssibrahimbas/claim-auth.go/src/jwt"
	"github.com/ssibrahimbas/claim-auth.go/src/result"
)

type Config struct {
	I18n   *i18n.I18n
	MsgKey string
}

func New(cnf *Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		u := c.Locals("user")
		if u == nil || u.(*jwt.UserClaim).IsExpired() {
			l, a := cnf.I18n.GetLanguagesInContext(c)
			return result.Error(cnf.I18n.Translate(cnf.MsgKey, l, a), fiber.StatusUnauthorized)
		}
		return c.Next()
	}
}
