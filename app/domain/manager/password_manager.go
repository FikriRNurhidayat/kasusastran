package manager

import "golang.org/x/crypto/bcrypt"

type PasswordManager interface {
	Encrypt(password string) (string, error)
	Compare(encryptedPassword string, password string) error
}

type PasswordManagerImpl struct {
	cost int
}

// Compare implements PasswordManager
func (*PasswordManagerImpl) Compare(encryptedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
}

func (m *PasswordManagerImpl) Encrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), m.cost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func NewPasswordManager(cost int) PasswordManager {
	return &PasswordManagerImpl{
		cost: cost,
	}
}
