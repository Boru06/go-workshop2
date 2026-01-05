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
    v4:=api.Group("/v4")

	 
	v1.Post("/fact/:fact",c.Factorial)

	v3.Post("/ball",c.TaxID)
	v1.Post("/register",c.Register)
	app.Get("/", c.Hello)



   //CRUD dogs
   dog := v1.Group("/dog")
   dog.Get("", c.GetDogs)
   dog.Get("/filter", c.GetDog)
   dog.Get("/json", c.GetDogsJson)
   dog.Post("/", c.AddDog)
   dog.Put("/:id", c.UpdateDog)
   dog.Delete("/:id", c.RemoveDog)
   dog.Get("/deleted",c.GetRemoveDog)
   dog.Get("/range",c.GetDogsRange)

   // company
   cpn:=v4.Group("/company")
   cpn.Get("/",c.GetcpnAll)
   cpn.Get("/filter",c.Getcpn)
   cpn.Post("/",c.Addcpn)
   cpn.Put("/:id",c.Updatecpn)
   cpn.Delete("/:id",c.Removecpn)
  

}

   

