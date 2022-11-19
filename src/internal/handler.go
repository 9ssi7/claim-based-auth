package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssibrahimbas/claim-auth.go/src/config"
	"github.com/ssibrahimbas/claim-auth.go/src/http"
	"github.com/ssibrahimbas/claim-auth.go/src/i18n"
	"github.com/ssibrahimbas/claim-auth.go/src/jwt"
	"github.com/ssibrahimbas/claim-auth.go/src/middlewares/claim_guard"
	"github.com/ssibrahimbas/claim-auth.go/src/middlewares/current_user"
	"github.com/ssibrahimbas/claim-auth.go/src/middlewares/required_auth"
	"github.com/ssibrahimbas/claim-auth.go/src/validator"
)

type Handler struct {
	s    *Srv
	c    *config.App
	h    *http.Client
	v    *validator.Validator
	i18n *i18n.I18n
	jwt  *jwt.Jwt
}

type HandlerParams struct {
	Srv   *Srv
	Http  *http.Client
	I18n  *i18n.I18n
	Jwt   *jwt.Jwt
	Valid *validator.Validator
	Cnf   *config.App
}

func NewHandler(p *HandlerParams) *Handler {
	return &Handler{
		s:    p.Srv,
		h:    p.Http,
		i18n: p.I18n,
		jwt:  p.Jwt,
		v:    p.Valid,
		c:    p.Cnf,
	}
}

func (h *Handler) InitAllVersions() {
	h.initV1()
}

func (h *Handler) initV1() {
	v1 := h.h.App.Group("/api/auth/v1")
	v1.Use(h.i18n.I18nMiddleware)
	v1.Post("/login", h.Login)
	v1.Post("/register", h.Register)
	v1.Put("/logout", h.currentUser(), h.requiredAuth(), h.Logout)
	v1.Put("/add-role", h.currentUser(), h.requiredAuth(), h.claimGuard("user"), h.AddRole)
	v1.All("/test", h.currentUser(), h.requiredAuth(), h.claimGuard("admin"), h.Test)
}

func (h *Handler) currentUser() fiber.Handler {
	return current_user.New(&current_user.Config{
		I18n:   h.i18n,
		Jwt:    h.jwt,
		MsgKey: messages.UnAuthorized,
	})
}

func (h *Handler) requiredAuth() fiber.Handler {
	return required_auth.New(&required_auth.Config{
		I18n:   h.i18n,
		MsgKey: messages.UnAuthenticated,
	})
}

func (h *Handler) claimGuard(r ...string) fiber.Handler {
	return claim_guard.New(&claim_guard.Config{
		Claims: r,
		I18n:   h.i18n,
		MsgKey: messages.UnAuthorized,
	})
}
