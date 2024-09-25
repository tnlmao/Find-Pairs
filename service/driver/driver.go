package driver

import "findindices/model"

type FindIndices interface {
	FindIndicesFunc(array []int, target int) model.Response
}
