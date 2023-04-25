package teststore_test

import (
	"testing"

	"github.com/GTech1256/http-rest-api/internal/app/model"
	"github.com/GTech1256/http-rest-api/internal/app/store"
	"github.com/GTech1256/http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser()

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	s := teststore.New()
	u1 := model.TestUser()
	_, err1 := s.User().Find(u1.ID)
	assert.EqualError(t, err1, store.ErrRecordNotFound.Error())

	s.User().Create(u1)
	userFound, err2 := s.User().Find(u1.ID)
	assert.NoError(t, err2)
	assert.NotNil(t, userFound)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	u1 := model.TestUser()
	_, err := s.User().FindByEmail(u1.Email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u1)
	userFound, err := s.User().FindByEmail(u1.Email)
	assert.NoError(t, err)
	assert.NotNil(t, userFound)
}
