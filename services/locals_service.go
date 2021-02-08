package services

import (
	"encoding/json"
	"github.com/benbarron/golang-auth-server/database"
	"github.com/gofiber/fiber/v2"
)

type LocalsStorage struct {

}

func NewLocalsStorage() *LocalsStorage {
	return &LocalsStorage{}
}

func (s *LocalsStorage) GetUser(ctx *fiber.Ctx) database.User {
	userJson := ctx.Locals("user").([]byte)
	var user database.User
	json.Unmarshal(userJson, &user)
	return user
}

func (s *LocalsStorage) SetUser(ctx *fiber.Ctx, u database.User) {
	userJson, _ := json.Marshal(u)
	ctx.Locals("user", userJson)
}
