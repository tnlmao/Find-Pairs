package handler

import (
	"encoding/json"
	"findindices/model"
	"findindices/service/driver"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Handler(svc driver.FindIndices) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := ctx.Request.Body
		fmt.Println(request)
		model := &model.Request{}
		decoder := json.NewDecoder(request)
		err := decoder.Decode(&model)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(model)
		response := svc.FindIndicesFunc(model.Array, model.Target)

		ctx.JSON(response.Code, response)
	}
}
