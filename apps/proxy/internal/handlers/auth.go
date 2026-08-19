package handlers

import (
	"proxy-dashboard/internal/database"

	"github.com/gofiber/fiber/v2"
)


type AuthHandler struct {
	DB *database.Queries
}

func NewAuthHandler(db *database.Queries) *AuthHandler {
	return &AuthHandler{DB: db}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	type Req struct {
		Email 	string `json:"email"`
		Password string `json:"password"`
	}

	var payload Req
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Format salah"})
	}

	user, err := h.DB.CreateUser(c.Context(), database.CreateUserParams{
		Email: payload.Email,
		Password: payload.Password,
	})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error" : "Gagal simpan ke DB" + err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "user berhasil dibuat!",
		"data": user,
	})
}