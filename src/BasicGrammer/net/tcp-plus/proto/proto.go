package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// Encode 编码消息
func Encode(message string) ([]byte, error) {
	// 消息长度：占4个字节
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)

	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}

	// 写入消息体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}

	return pkg.Bytes(), nil
}

// Decode 解码消息
func Decode(reader *bufio.Reader) (string, error) {
	// 读取消息长度，前4个字节
	lengthByte, _ := reader.Peek(4) // 获取前4个字节
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}

	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	// 读取包体
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
