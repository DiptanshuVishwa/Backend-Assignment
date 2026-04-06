package src
import ( "strings"
"github.com/gofiber/fiber/v2"
"github.com/golang-jwt/jwt/v5"
)
func AuthMiddleware(c *fiber.Ctx)error{
	authHeader := c.Get("Authorization")
	if authHeader == ""{
		return c.Status(401).JSON(fiber.Map{
			"error":"Missing token",
		})
	}
	tokenString := strings.Split(authHeader, " ")[1]
	token,err := jwt.Parse(tokenString, func(t *jwt.Token)(interface{},error){
		return jwtSecret, nil
	})
	if err != nil || !token.Valid{
		return c.Status(401).JSON(fiber.Map{
			"error":"Invalid token",
		})
	}
	claims := token.Claims.(jwt.MapClaims)
	c.Locals("username",claims["username"])
	c.Locals("role",claims["role"])
	return c.Next()
}
func RoleMiddleware(allowedRoles ...string) fiber.Handler{
	return func(c *fiber.Ctx)error{
		userRole:=c.Locals("role").(string)
		for _, role := range allowedRoles{
			if role == userRole{
				return c.Next()
			}
		}
		return c.Status(403).JSON(fiber.Map{
			"error":"Access denied",
		})
	}
}