package src

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx)error{
		return c.SendString("Backend is running!")
	})
	app.Post("/login", LoginController)
	app.Get("/dashboard", AuthMiddleware, func(c *fiber.Ctx) error {
		username := c.Locals("username")
		role := c.Locals("role")
		return c.JSON(fiber.Map{
			"message":  "Welcome",
			"username": username,
			"role":     role,
		})
	})
	app.Get("/admin", AuthMiddleware, RoleMiddleware("admin", "superadmin"), func(c *fiber.Ctx) error {
		return c.SendString("Welcome Admin")
	})
	app.Get("/teacher", AuthMiddleware, RoleMiddleware("teacher", "superadmin"), func(c *fiber.Ctx) error {
		return c.SendString("Welcome Teacher")
	})
	app.Get("/student", AuthMiddleware, RoleMiddleware("student", "superadmin"), func(c *fiber.Ctx) error {
		return c.SendString("Welcome Student")
	})
}
