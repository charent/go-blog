package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetSession(c *gin.Context, key string, value string) {
	session := sessions.Default(c)
	session.Set(key, value)
	session.Save()
}

func GetSession(c *gin.Context, key string) string {
	session := sessions.Default(c)
	value := session.Get(key)
	return value.(string)
}
