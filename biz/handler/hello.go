package handler

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sophist/go-framework/biz/service"
)

func Hello(ctx context.Context, c *app.RequestContext) {
	words, err := service.GenerateWords()
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Println("generate hello from: %")
	c.JSON(consts.StatusOK, utils.H{
		"message": words,
	})
}
