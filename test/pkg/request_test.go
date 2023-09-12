package pkg

import (
	"fmt"
	"go-grpc/pkg/request"
	"testing"
)

func TestTree(t *testing.T) {
	url := "http://www.baidu.com"
	data := ""
	headers := make(map[string]interface{})
	res, err := request.Get(url, data, headers)
	if err == nil {
		fmt.Println(len(res))
	}
}
