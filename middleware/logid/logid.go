package logid

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sherlockhua/koala/logs"
)

const (
	LOGID_Header_Name = "koala-logid"
)

func LogIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 生成一个唯一的UUID作为logid
		id := uuid.New().String()
		now := time.Now()
		logid := fmt.Sprintf("%04d%02d%02d%02d%02d%02d%s",
			now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(),
			id,
		)
		// 将logid设置到请求的context中
		ctx := context.WithValue(c.Request.Context(), logs.LOGID, logid)
		c.Request = c.Request.WithContext(ctx)

		// 继续执行后续的中间件或请求处理函数
		c.Next()
		c.Writer.Header().Set(LOGID_Header_Name, logid)
	}
}
