package claim_guard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssibrahimbas/claim-auth.go/src/i18n"
	"github.com/ssibrahimbas/claim-auth.go/src/jwt"
	"github.com/ssibrahimbas/claim-auth.go/src/middlewares/current_user"
	"github.com/ssibrahimbas/claim-auth.go/src/result"
)

type Config struct {
	Claims []string
	I18n   *i18n.I18n
	MsgKey string
}

func New(cnf *Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		u := current_user.Parse(ctx)
		if checkClaims(u, cnf.Claims) {
			return ctx.Next()
		}
		l, a := cnf.I18n.GetLanguagesInContext(ctx)
		return result.Error(cnf.I18n.Translate(cnf.MsgKey, l, a), fiber.StatusForbidden)
	}
}

func checkClaim(u *jwt.UserClaim, c string) bool {
	for _, r := range u.Roles {
		if r == c {
			return true
		}
	}
	return false
}

func checkClaims(u *jwt.UserClaim, cs []string) bool {
	for _, c := range cs {
		if checkClaim(u, c) {
			return true
		}
	}
	return false
}
