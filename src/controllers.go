package src
import ( "time"
	"github.com/gofiber/fiber/v2"
		"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("canigetthisinternship")

var users = map[string]struct{
	Password string
	Role string
}{
	"vikasJi" : {"123","superadmin"},
	"diptanshu" : {"123","admin"},
	"ashwin" : {"123","teacher"},
	"student1" : {"123","student"},
}

func LoginController(c *fiber.Ctx)error{
	var req LoginRequest
	err := c.BodyParser(&req)
	if err != nil{
		return c.Status(400).JSON(fiber.Map{
			"error" : "Invalid request body",
		})
	}
	user,exists := users[req.Username]
	if !exists || user.Password != req.Password{
		return c.Status(401).JSON(fiber.Map{
			"error" : "Invalid credentials",
		})
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username" : req.Username,
		"role" : user.Role,
		"exp" : time.Now().Add(time.Hour*24).Unix(),
	})

	tokenString,err := token.SignedString(jwtSecret)
	if err != nil{
		return c.Status(500).JSON(fiber.Map{
			"error":"Could not generate token",
		})
	}
	return c.JSON(fiber.Map{
		"message":"Login successfull","token": tokenString,
	})
}