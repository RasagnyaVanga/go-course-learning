package mock_practice

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestCARSTRUCT_Start(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	MockInterface := NewMockCAR(controller) //instantiate car and assign it to interface
	var ctx *gin.Context
	MockInterface.EXPECT().EngineCheck(ctx, MockInterface).Return(true) //mocking engine check
	result := CARSTRUCT{}.Start(ctx, "carkey", MockInterface)
	assert.Equal(t, true, result)
}
