package main

import (
	"fmt"
	"os/exec"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/mntky/convA2J/pkg"
)


type ConvData struct {
	Token           string
    Team_id         string
    Team_domain     string
    Service_id      string
    Channel_id      string
    Channel_name	string
    Timestamp       string
    User_id         string
    User_name       string
}

func main() {
	g := gin.Default()
    g.POST("/", func(c *gin.Context) {
            buf := make([]byte, 2048)
            n, _ := c.Request.Body.Read(buf)
            b := string(buf[0:n])

            decode(b)
    })
    g.Run(":8080")
}

func decode(data string) {
	var convdata *ConvData = &ConvData{}
    resp := conv.Convert(data)
    err := json.Unmarshal([]byte(resp), convdata)
    if err != nil {
            fmt.Println(err)
    }

    fmt.Println(convdata.User_name)
	out, err := exec.Command("sh", "./send.sh", convdata.User_name).Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
