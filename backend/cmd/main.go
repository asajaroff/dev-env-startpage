package main

import (
	"log"
	"net/http"
	"os/exec"

	"github.com/labstack/echo/v4"
)

func main() {
	// AWS credentials config

	// Checks before starting the web server
	path, err := exec.LookPath("helm")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(path)

	path, err = exec.LookPath("vcluster")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(path)

	// Bootstrap echo server
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "dev-env-home")
	})

	// List vclusters
	e.GET("/v1/clusters/list", func(c echo.Context) error {
		cmd, err := exec.Command("vcluster", "list", "--output=json").Output()
		if err != nil {
			log.Fatal(err)
			return c.String(http.StatusFailedDependency, "There was a problem fetching the clusters")
		}
		log.Println(string(cmd))
		return c.String(http.StatusOK, string(cmd))
	})

	// Create vclusters
	e.POST("/v1/clusters/new", func(c echo.Context) error {
		name := c.FormValue("vcluster-name")
		log.Printf("Creating vcluster %s", name)
		cmd, err := exec.Command("vcluster", "create", name, "--connect=false").Output()
		if err != nil {
			log.Printf("%v", string(cmd))
			log.Fatal(err)
			return c.String(http.StatusUnprocessableEntity, "Our worker monkeys were unable to create your dev enviroment")
		}
		log.Println(string(cmd))

		return c.String(http.StatusOK, "Cluster created!")
	})

	// Return Kubeconfig

	e.Logger.Fatal(e.Start(":2525"))
}
