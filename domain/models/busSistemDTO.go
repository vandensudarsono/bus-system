package models

import errorcode "github.com/vandensudarsono/bus-system/internal/errorCode"

type (
	Response struct {
		Data   interface{} `json:"data"`
		Meta   interface{} `json:"meta"`
		Status Status      `json:"status"`
	}

	Status struct {
		Code    errorcode.ErrorCode `json:"code"`
		Message string              `json:"message"`
	}
)
