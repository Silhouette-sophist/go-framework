package service

import (
	"fmt"
	"github.com/sophist/go-framework/biz/util"
)

// GenerateWords 生成信息
func GenerateWords() (string, error) {
	return fmt.Sprintf("generate at %s", util.FormatNow()), nil
}
