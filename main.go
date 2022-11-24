// Program to illustrate fmt.Print()

package main

// import fmt package
import (
	"github.com/codetagon/crud_golang_reactjs/bootstrap"
	"github.com/codetagon/crud_golang_reactjs/repository"
	"github.com/gofiber/fiber/v2"
)

type Repository repository.Repository

func main() {
	app := fiber.New()
	bootstrap.InitializeApp(app)
}
