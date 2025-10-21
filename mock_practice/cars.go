package mock_practice

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type CAR interface {
	Start(*gin.Context, string, CAR) bool
	EngineCheck(*gin.Context, CAR) bool
} //mockgen generates a mock implementation for this interface  //generates a struct that implements CAR //MockCAR
type CARSTRUCT struct {
}

func (car CARSTRUCT) Start(context *gin.Context, key string, carInterface CAR) bool {
	if carInterface.EngineCheck(context, carInterface) {
		if key == "carkey" {
			fmt.Println("Let's start the car!")
			return true
		} else {
			fmt.Println("You don't have the right key!")
			return false
		}
	} else {
		return false
	}
}
func (car CARSTRUCT) EngineCheck(context *gin.Context, carInterface CAR) bool {
	return true
}
