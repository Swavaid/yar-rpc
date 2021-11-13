package dis_sys

import (
	"encoding/binary"
	"io"
	"net"
)

// 编写数据会话中读写

// 会话连接的结构体

type Session struct {
	conn net.Conn
}
// 创建新连接

func NewSession(conn net.Conn) *Session {
	return &Session{conn: conn}
}

// 向连接中写数据

func (s Session) Write(data []byte) error {
	// 4字节头+数据长度切片
	buf := make([]byte, 4+len(data))
	// 写入头部数据，记录数据长度
	// binary 只认固定长度的类型，所以使用了uint32，而不是直接写入，采用大端字节排序法
	// 将记录了数据长度的数据的编码形式记录在切片的0~4中
	binary.BigEndian.PutUint32(buf[:4], uint32(len(data)))
	//将剩余数据copy到4以后的位置
	copy(buf[4:], data)
	//将数据写入链接
	_, err := s.conn.Write(buf)
	if err != nil {
		return err
	}
	return nil
}

// 从连接中读数据
func (s Session) Read() ([]byte, error) {
	// 读取头部长度（前四位）
	header := make([]byte, 4)
	// 按头部长度， 读取头部数据
	_, err := io.ReadFull(s.conn, header)
	if err != nil {
		return nil, err
	}
	// 读取数据长度（header信息）
	dataLen := binary.BigEndian.Uint32(header)
	// 按照数据长度去读取数据
	data := make([]byte, dataLen)
	_, err = io.ReadFull(s.conn, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
