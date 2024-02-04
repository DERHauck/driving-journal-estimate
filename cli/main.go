package main

import (
	"driving-journal-estimate/cmd"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	if env, _ := os.LookupEnv("ENVIRONMENT"); env == "BACKEND" {
		router := gin.Default()
		initRoutes(router)
		err := router.Run(":8080")
		if err != nil {
			_ = fmt.Sprintf(err.Error())
		}
		return
	} else {
		cmd.Execute()
	}
}
