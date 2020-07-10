package employee

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func FindAll(c *gin.Context) {
	fmt.Println("Starting the application...")
	response, err := http.Get(viper.GetString("employee.url"))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		c.Data(http.StatusOK, gin.MIMEJSON, data)
	}
}
