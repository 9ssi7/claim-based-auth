package current_user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssibrahimbas/claim-auth.go/src/i18n"
	"github.com/ssibrahimbas/claim-auth.go/src/jwt"
	"github.com/ssibrahimbas/claim-auth.go/src/result"
)

type CurrentUser struct {
	ID    string   `json:"uuid"`
	Email string   `json:"email"`
	Roles []string `json:"roles"`
}

type Config struct {
	Jwt    *jwt.Jwt
	I18n   *i18n.I18n
	MsgKey string
}

func New(cnf *Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		t := c.Cookies("token")
		if t == "" {
			return c.Next()
		}
		res, err := cnf.Jwt.VerifyAndParse(t)
		if err != nil || res == nil {
			l, a := cnf.I18n.GetLanguagesInContext(c)
			return result.Error(cnf.I18n.Translate(cnf.MsgKey, l, a), fiber.StatusUnauthorized)
		}
		c.Locals("user", res)
		return c.Next()
	}
}

func Parse(c *fiber.Ctx) *jwt.UserClaim {
	return c.Locals("user").(*jwt.UserClaim)
}
