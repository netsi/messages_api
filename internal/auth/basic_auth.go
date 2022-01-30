package auth

import (
	"github.com/gin-gonic/gin"
	"log"
	"messages_api/internal/users/repository"
	"net/http"
)

// BasicAuth middleware that tries to get an admin user from the userRepository, if the user is not found
// or the password does not match the request is aborted with 401 HTTP status.
func BasicAuth(userRepository repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			abortAndReturnUnauthorized(c)

			return
		}

		user, err := userRepository.GetAdmin(ctx, username)
		if err != nil {
			log.Printf("error while getting the admin by username: %s", err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)

			return
		}

		if user == nil {
			log.Println("admin not found")
			abortAndReturnUnauthorized(c)

			return
		}

		if user.Password != password {
			log.Println("invalid credentials attempt")
			abortAndReturnUnauthorized(c)

			return
		}

		c.Next()
	}
}

func abortAndReturnUnauthorized(c *gin.Context) {
	c.Header("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	c.AbortWithStatus(http.StatusUnauthorized)
}
