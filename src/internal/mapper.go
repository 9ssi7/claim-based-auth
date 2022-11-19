package internal

import (
	"time"

	"github.com/ssibrahimbas/claim-auth.go/src/dto"
	"github.com/ssibrahimbas/claim-auth.go/src/entity"
	"github.com/ssibrahimbas/claim-auth.go/src/jwt"
)

type Mapper struct{}

func NewMapper() *Mapper {
	return &Mapper{}
}

func (m *Mapper) MapRegisterDtoToUserEntity(d *dto.RegisterRequest, pw []byte) *entity.User {
	t := time.Now()
	return &entity.User{
		ID:        "",
		Email:     d.Email,
		Password:  pw,
		CreatedAt: t,
		UpdatedAt: t,
		Roles:     []string{"user"},
	}
}

func (m *Mapper) MapUserToJwtClaim(u *entity.User, exp int64) *jwt.UserClaim {
	return &jwt.UserClaim{
		UUID:      u.ID,
		Email:     u.Email,
		Roles:     u.Roles,
		ExpiresIn: exp,
	}
}
