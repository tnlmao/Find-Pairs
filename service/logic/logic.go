package logic

import (
	"findindices/model"
	"findindices/service"
	"findindices/service/driver"
)

type FindIndicesService struct{}

func NewFindIndicesService() driver.FindIndices {
	return &FindIndicesService{}
}

// FindIndicesFunc implements driver.FindIndices.
func (f *FindIndicesService) FindIndicesFunc(array []int, target int) model.Response {
	if len(array) == 0 {
		return service.ServiceResponse(model.Response{
			Code: 500,
			Msg:  "Array Empty",
		})
	}
	resultSlice := [][]int{}
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			innerSlice := []int{}

			if array[i]+array[j] == target {
				innerSlice = append(innerSlice, i, j)
				resultSlice = append(resultSlice, innerSlice)
			}

		}
	}
	if len(resultSlice) == 0 {
		return model.Response{
			Code: 200,
			Msg:  "No Pairs Add upto The Target Number",
			Resp: resultSlice,
		}
	}
	return model.Response{
		Code: 200,
		Msg:  "Success",
		Resp: resultSlice,
	}
}
