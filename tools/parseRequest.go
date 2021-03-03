package tools

import (
	"errors"
	"github.com/Irbi/citrouser/constants"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetIntFromParams(ctx *gin.Context, paramName string) (val uint, err error) {
	param, ok := ctx.Params.Get(paramName)
	if !ok {
		err = errors.New(constants.ERR_INVALID_REQUEST)
		return
	}

	valInt, err := strconv.Atoi(param)
	if err != nil {
		return
	}

	val = uint(valInt)

	return
}

func GetStringFromParams(ctx *gin.Context, paramName string) (val string, err error) {
	val, ok := ctx.Params.Get(paramName)
	if !ok {
		err = errors.New(constants.ERR_INVALID_REQUEST)
		return
	}

	return
}
