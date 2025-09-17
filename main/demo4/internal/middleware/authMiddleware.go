package middleware

import (
	"demo4/pkg/auth"
	error2 "demo4/pkg/error"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWT认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			error2.ThrowErr(c, error2.ErrInvalidCredentials, "请求头中没有找到token")
			return
		}

		// 解析Bearer Token
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			error2.ThrowErr(c, error2.ErrInvalidCredentials, "token格式有问题")
			return
		}

		// 验证Token
		claims, err := auth.ParseToken(parts[1])
		if err != nil {
			error2.ThrowErr(c, error2.ErrInvalidCredentials, "token无效： "+err.Error())
			return
		}

		// 将用户信息存入上下文
		c.Set("userID", claims.UserID) // newCtx := c.WithValue("userID", 123)这会创建新的上下文，原 c 中的数据不会被修改
		c.Set("userRole", claims.Role)
		//当中间件执行到 c.Next() 时，会暂停当前中间件的执行，先去执行后续的中间件和最终的业务处理器。
		//等后续所有逻辑执行完毕后，会回到 c.Next() 之后的代码，继续执行当前中间件剩余的逻辑（这种特性可以用来实现 “后置处理”，如记录响应时间、清理资源等）
		//如果中间件中没有调用 c.Next()，那么请求处理流程会在当前中间件中中断，后续的中间件和业务处理器都不会被执行。
		c.Next() // 关键：让请求继续传递到下一个中间件或处理器
		// 3. 下一个中间件或处理器完成之后回来接着（可选）后续处理（如记录日志、统计等）
		// 会在业务处理器执行完后再回到这里
	}
}

// 角色权限中间件
func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("userRole")
		if !exists || role != requiredRole {
			err := error2.ErrUnauthorized
			c.Error(err)
			c.Abort()
			return
		}
		c.Next()
	}
}
