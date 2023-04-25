package model_test

import (
	"testing"

	"github.com/GTech1256/http-rest-api/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser()

	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "Valid",
			u: func() *model.User {
				return model.TestUser()
			},
			isValid: true,
		},
		{
			name: "with Encrypted Password",
			u: func() *model.User {
				u := model.TestUser()
				u.Password = ""
				u.EncryptedPassword = "EncryptedPassword"

				return u
			},
			isValid: true,
		},
		{
			name: "Empty Email",
			u: func() *model.User {
				u := model.TestUser()
				u.Email = ""

				return u

			},
			isValid: false,
		},
		{
			name: "Invalid Email",
			u: func() *model.User {
				u := model.TestUser()
				u.Email = "invalid"

				return u

			},
			isValid: false,
		},
		{
			name: "Empty Password",
			u: func() *model.User {
				u := model.TestUser()
				u.Password = ""

				return u

			},
			isValid: false,
		},
		{
			name: "Short Password",
			u: func() *model.User {
				u := model.TestUser()
				u.Email = "short"

				return u

			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}
