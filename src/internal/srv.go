package internal

import (
	"time"

	"github.com/ssibrahimbas/claim-auth.go/src/cipher"
	"github.com/ssibrahimbas/claim-auth.go/src/dto"
	"github.com/ssibrahimbas/claim-auth.go/src/entity"
	"github.com/ssibrahimbas/claim-auth.go/src/i18n"
	"github.com/ssibrahimbas/claim-auth.go/src/jwt"
)

type Srv struct {
	r      *Repo
	i18n   *i18n.I18n
	jwt    *jwt.Jwt
	m      *Mapper
	cipher *cipher.Cipher
}

type SrvParams struct {
	Repo   *Repo
	I18n   *i18n.I18n
	Jwt    *jwt.Jwt
	Cipher *cipher.Cipher
}

func NewSrv(p *SrvParams) *Srv {
	return &Srv{
		r:      p.Repo,
		i18n:   p.I18n,
		jwt:    p.Jwt,
		m:      NewMapper(),
		cipher: p.Cipher,
	}
}

func (s *Srv) Register(d *dto.RegisterRequest) (string, bool) {
	if _, scs := s.r.GetUserByEmail(d.Email); scs {
		return messages.UserAlreadyExist, false
	}
	pw, err := s.cipher.Encrypt(d.Password)
	if err != nil {
		panic(err)
	}
	u := s.m.MapRegisterDtoToUserEntity(d, pw)
	u = s.r.CreateUser(u)
	return s.generateToken(u), true
}

func (s *Srv) Login(d *dto.LoginRequest) (string, bool) {
	u, scs := s.r.GetUserByEmail(d.Email)
	if !scs {
		return messages.UserNotFound, false
	}
	if c := s.checkPassword(d.Password, u); !c {
		return messages.WrongPassword, false
	}
	return s.generateToken(u), true
}

func (s *Srv) Logout(dto *dto.LogOutRequest) string {
	return s.expireToken(dto.Token)
}

func (s *Srv) AddAdminRole(dto *dto.AdminRoleRequest) string {
	s.r.AddAdminRole(dto.UUID)
	return s.expireToken(dto.Token)
}

func (s *Srv) generateToken(u *entity.User) string {
	t, err := s.jwt.Sign(s.m.MapUserToJwtClaim(u, time.Now().Add(time.Hour*24).Unix()))
	if err != nil {
		panic(err)
	}
	return t
}

func (s *Srv) checkPassword(pw string, u *entity.User) bool {
	b, err := s.cipher.Compare(pw, u.Password)
	if err != nil {
		panic(err)
	}
	return b
}

func (s *Srv) expireToken(t string) string {
	tk, err := s.jwt.Expire(t)
	if err != nil {
		panic(err)
	}
	return tk
}
