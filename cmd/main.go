package main

import (
	"log"
	"net/http"
	"os/exec"

	"github.com/labstack/echo/v4"
)

func main() {
	// AWS credentials config

	// Bootstrap echo server
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "dev-env-home")
	})

	e.GET("/v1/clusters/list", func(c echo.Context) error {
		cmd, err := exec.Command("vcluster", "list").Output()
		if err != nil {
			log.Fatal(err)
			return c.String(http.StatusUnprocessableEntity, "Our worker monkeys had trouble listing your clusters")
		}
		log.Println(string(cmd))
		return c.String(http.StatusOK, "?")
	})

	e.POST("/v1/clusters/new", func(c echo.Context) error {
		name := c.FormValue("vcluster-name")
		log.Printf("Creating vcluster %s", name)
		cmd, err := exec.Command("vcluster", "create", name).Output()
		if err != nil {
			log.Printf("%v", string(cmd))
			log.Fatal(err)
			return c.String(http.StatusUnprocessableEntity, "Our worker monkeys were unable to create your dev enviroment")
		}
		log.Println(string(cmd))

		return c.String(http.StatusOK, "Cluster created!")
	})

	e.Logger.Fatal(e.Start(":2525"))
}
