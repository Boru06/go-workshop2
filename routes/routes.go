package routes


import(
	 c "workshop2/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func Rountes(app *fiber.App) {
	//5.0 basicAuth
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "21022566",
		},
	}))
    
	api:= app.Group("/api")
	v1:=api.Group("/v1")
	v3:=api.Group("/v3")

	v1.Post("/fact/:fact",c.Factorial)

	v3.Post("/ball",c.TaxID)
	v1.Post("/register",c.Register)
	app.Get("/", c.Hello)



}
