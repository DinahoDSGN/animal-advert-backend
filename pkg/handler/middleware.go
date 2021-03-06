package handler

import (
	_ "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strings"
)

const userCtx = "userId"

func (h *Handler) CORSMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "false")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		log.Println("CORS initialized")

		c.Next()
	}
}

func (h *Handler) UserIdentity(c *gin.Context) {

	header := c.GetHeader("Authorization")
	if header == "" {
		newErrorResponse(c, fiber.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, fiber.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, fiber.StatusUnauthorized, err.Error())
		return
	}

	if userId == 0 {
		newErrorResponse(c, fiber.StatusInternalServerError, "invalid access token")
		return
	}
}
