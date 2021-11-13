package dis_sys

import (
	"net"
	"reflect"
)

// 声明客户端

type Client struct {
	conn net.Conn
}
// 创建客户端对象

func NewClient(conn net.Conn) *Client {
	return &Client{conn: conn}
}
// 实现通用的RPC客户端
// 绑定RPC使用的方法
// 传入访问的函数名

// 函数具体实现在Server端, Client只有函数原型
// 使用MakeFunc() 完成原型到函数的调用
// fPtr指向函数原型

func (c *Client) CallRPC(rpcName string, fPtr interface{}){
	// 通过反射，获取fPtr未初始化的函数原型,fn是一个函数原型
	fn := reflect.ValueOf(fPtr).Elem()

	// 另一个函数，是对第一个函数参数操作
	f := func(args []reflect.Value) []reflect.Value {
		// 处理输入的参数
		inArgs := make([]interface{}, 0, len(args))
		for _, arg := range args{
			inArgs = append(inArgs, arg.Interface())
		}
		// 创建连接
		cliSession := NewSession(c.conn)
		// 编码数据
		reqRPC := RPCData{Name: rpcName, Args: inArgs}
		b, err := encode(reqRPC)
		if err != nil {
			panic(nil)
		}
		// 写出数据
		err = cliSession.Write(b)
		if err != nil {
			panic(nil)
		}
		// 读响应数据
		respBytes, err := cliSession.Read()
		if err != nil {
			panic(err)
		}


		// 解码数据
		respRPC, err := decode(respBytes)
		if err != nil {
			panic(err)
		}
		// 处理服务端返回的数据
		outArgs := make([]reflect.Value, 0, len(respRPC.Args))
		for _, arg := range respRPC.Args {
			// 必须进行nil转换
			if arg != nil {
				// 必须填充一个真正的类型，不能是nil
				outArgs = append(outArgs, reflect.ValueOf(arg))
				continue
			}

		}
		return outArgs
	}

	//fn为函数原型，v是函数实体

	v := reflect.MakeFunc(fn.Type(), f)
	// 为函数fPtr赋值
	fn.Set(v)
}
