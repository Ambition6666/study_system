package middleware

import "github.com/gin-gonic/gin"

func If_admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetInt("role")
		if role != 2 {
			c.JSON(401, "权限不足")
			c.Abort()
		} else {
			c.Next()
		}
	}
}
