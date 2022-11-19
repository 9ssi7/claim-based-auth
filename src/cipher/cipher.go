package cipher

import "golang.org/x/crypto/bcrypt"

type Cipher struct{}

func New() *Cipher {
	return &Cipher{}
}

func (c *Cipher) Encrypt(t string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(t), 9)
}

func (c *Cipher) Compare(t string, h []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(h, []byte(t))
	return err == nil, err
}
