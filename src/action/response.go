package action

import (
	"encoding/json"

	"github.com/dxvgef/tsing"
)

// 响应单键值的数据
type RespData struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data,omitempty"`
}

// 响应多键值的数据
type RespMapData struct {
	Error string                 `json:"error"`
	Data  map[string]interface{} `json:"data,omitempty"`
}

// 构建一个RespMapData
func makeRespMapData() (respData RespMapData) {
	respData.Data = make(map[string]interface{})
	return
}

// 输出JSON数据给客户端
func JSON(ctx *tsing.Context, status int, data interface{}) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ctx.ResponseWriter.WriteHeader(status)
	_, err = ctx.ResponseWriter.Write(dataBytes)
	return err
}

// 输出String给客户端
func String(ctx *tsing.Context, status int, data string) error {
	// 设置header信息
	ctx.ResponseWriter.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	ctx.ResponseWriter.WriteHeader(status)
	_, err := ctx.ResponseWriter.Write([]byte(data))
	return err
}

// 输出HTTP状态码，无返回数据
func Status(ctx *tsing.Context, status int) error {
	ctx.ResponseWriter.WriteHeader(status)
	return nil
}
