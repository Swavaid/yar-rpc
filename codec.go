package dis_sys

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// 定义数据格式和编解码

type RPCData struct {
	// 访问的函数
	Name string
	// 访问时传的参数
	Args []interface{}
}

// 编码   转成二进制
func encode(data RPCData) ([]byte, error) {
	var buf bytes.Buffer
	// 获取编码器
	bufEnc := gob.NewEncoder(&buf)


	// 对数据编码
	err:=bufEnc.Encode(data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return buf.Bytes(), nil
}

// 解码    转回原数据
func decode(b []byte) (RPCData, error) {
	buf := bytes.NewBuffer(b)
	// 获取解码器
	bufDec := gob.NewDecoder(buf)
	var data RPCData


	// 对数据解码
	err:=bufDec.Decode(&data)
	if err != nil {
		return data, nil
	}
	return data, nil
}
