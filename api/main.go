package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

const Filepath = "/data/times_called"

func EnsureFileExists() {
	_, err := os.Stat(Filepath)
	if os.IsNotExist(err) {
		_, err := os.Stat("/data")
		if os.IsNotExist(err) {
			err = os.Mkdir("/data", 0o644)
			if err != nil {
				log.Fatal(err)
			}
		}
		_, err = os.Create(Filepath)
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile(Filepath, []byte("0"), 0o644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// GetTimesCalled reads the number of times the API was called from a file
// increments it by one and writes to the file.
//
// There might be a race condition between different workers on different pods,
// but the example is just to get a feel for the differences between stateful
// sets and deployments.
func GetTimesCalled() int {
	contents, err := os.ReadFile(Filepath)
	if err != nil {
		log.Fatal(err)
	}

	num, err := strconv.Atoi(string(contents))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := os.WriteFile(
			Filepath,
			[]byte(strconv.Itoa(num+1)),
			0o644,
		)
		if err != nil {
			log.Fatal(err)
		}
	}()

	return num
}

func GetRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg":          "hello world!",
			"times_called": GetTimesCalled(),
		})
	})
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "ok",
		})
	})
	return r
}

func main() {
	r := GetRouter()
	EnsureFileExists()
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
