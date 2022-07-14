package manager_test

import (
	"testing"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/manager"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestPasswordManager_Encrypt(t *testing.T) {
	passwordManager := manager.NewPasswordManager(bcrypt.DefaultCost)
	plainPassword := "12345"
	encryptedPassword, err := passwordManager.Encrypt(plainPassword)

	if err != nil {
		t.Fail()
	}

	assert.Nil(t, bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(plainPassword)))
	assert.NotNil(t, bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte("yagaktau")))
}

func TestPasswordManager_Compare(t *testing.T) {
	passwordManager := manager.NewPasswordManager(bcrypt.DefaultCost)
	plainPassword := []byte("12345")
	encryptedPassword, err := bcrypt.GenerateFromPassword(plainPassword, bcrypt.DefaultCost)

	if err != nil {
		t.Fail()
	}

	err = passwordManager.Compare(string(encryptedPassword), string(plainPassword))

	if err != nil {
		t.Fail()
	}

	otherPlainPassword := []byte("yagaktau")
	err = passwordManager.Compare(string(encryptedPassword), string(otherPlainPassword))

	assert.NotNil(t, err)
}
