package main

import (
	"fmt"
	"net/url"
)

func main() {
	key := "测试编码文件" + `{"name":"juicefs"}` + string('\u001F')

	u := url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   fmt.Sprintf("/v1/%s", url.PathEscape(key)),
	}

	fmt.Println(u.String())
}
