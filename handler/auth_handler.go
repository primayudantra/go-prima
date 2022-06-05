package handler

import (
	"encoding/json"
	"fmt"
	"go-prima/repository"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) RegistrationHandler(c *fiber.Ctx) (err error) {
	ctx := c.Context()
	body := c.Body()

	incomingRequest := &RegistrationRequest{}
	err = json.Unmarshal(body, incomingRequest)

	if err != nil {
		return err
	}

	isUserExist, _ := s.store.GetUser(ctx, incomingRequest.Email)

	if isUserExist.Email != "" {
		return fiber.NewError(fiber.ErrBadRequest.Code, "something wrong")
	}

	// Encrypt password
	encryptedPassword := encrypt(incomingRequest.Password)

	arg := repository.RegistrationParams{
		Email:        incomingRequest.Email,
		PasswordHash: encryptedPassword,
	}

	err = s.store.Registration(ctx, arg)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return c.JSON(incomingRequest)
}

func (s *Server) LoginHandler(c *fiber.Ctx) (err error) {
	ctx := c.Context()
	body := c.Body()

	incomingRequest := &LoginRequest{}
	err = json.Unmarshal(body, incomingRequest)

	if err != nil {
		fmt.Println(err)
	}

	user, _ := s.store.GetUser(ctx, incomingRequest.Email)

	if user.PasswordHash != encrypt(incomingRequest.Password) {
		return fiber.NewError(fiber.ErrBadRequest.Code, "credential is not valid")
	}

	token := CreateAccessToken()

	return c.JSON(&LoginResponse{
		Token: token,
	})
}

type RegistrationRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
