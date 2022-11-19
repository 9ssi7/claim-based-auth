package internal

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ssibrahimbas/claim-auth.go/src/dto"
	"github.com/ssibrahimbas/claim-auth.go/src/middlewares/current_user"
	"github.com/ssibrahimbas/claim-auth.go/src/result"
)

func (h *Handler) Register(ctx *fiber.Ctx) error {
	d := &dto.RegisterRequest{}
	h.parseBody(ctx, d)
	l, a := h.i18n.GetLanguagesInContext(ctx)
	d.IP = ctx.IP()
	t, s := h.s.Register(d)
	if !s {
		return result.Error(h.i18n.Translate(t, l, a), fiber.StatusConflict)
	}
	h.setTokenToCookie(ctx, t)
	return result.Success(h.i18n.Translate(messages.RegisterSuccess, l, a), fiber.StatusCreated)
}

func (h *Handler) Login(ctx *fiber.Ctx) error {
	d := &dto.LoginRequest{}
	h.parseBody(ctx, d)
	l, a := h.i18n.GetLanguagesInContext(ctx)
	t, s := h.s.Login(d)
	if !s {
		return result.Error(h.i18n.Translate(t, l, a), fiber.StatusBadRequest)
	}
	h.setTokenToCookie(ctx, t)
	return result.Success(h.i18n.Translate(messages.LoginSuccess, l, a), fiber.StatusOK)
}

func (h *Handler) Logout(ctx *fiber.Ctx) error {
	u := current_user.Parse(ctx)
	t := h.s.Logout(&dto.LogOutRequest{
		UUID:  u.UUID,
		Token: ctx.Cookies(handlerKeys.Token),
	})
	h.setTokenToCookie(ctx, t)
	l, a := h.i18n.GetLanguagesInContext(ctx)
	return result.Success(h.i18n.Translate(messages.LogoutSuccess, l, a), fiber.StatusOK)
}

func (h *Handler) AddRole(ctx *fiber.Ctx) error {
	u := current_user.Parse(ctx)
	t := h.s.AddAdminRole(&dto.AdminRoleRequest{
		UUID:  u.UUID,
		Token: ctx.Cookies(handlerKeys.Token),
	})
	h.setTokenToCookie(ctx, t)
	l, a := h.i18n.GetLanguagesInContext(ctx)
	return result.Success(h.i18n.Translate(messages.AddRoleSuccess, l, a), fiber.StatusOK)
}

func (h *Handler) Test(ctx *fiber.Ctx) error {
	return result.Success("test", fiber.StatusOK)
}

func (h *Handler) parseBody(c *fiber.Ctx, d interface{}) {
	h.h.ParseBody(c, h.v, h.i18n, d)
}

func (h *Handler) setTokenToCookie(c *fiber.Ctx, t string) {
	c.Cookie(&fiber.Cookie{
		Name:     handlerKeys.Token,
		Value:    t,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: h.c.JWT.HTTPOnly,
		Secure:   h.c.JWT.Secure,
		Domain:   h.c.JWT.Domain,
	})
}
