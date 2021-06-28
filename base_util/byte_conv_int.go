package base_util

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//isSymbol表示有无符号
func BytesToInt(b []byte, isSymbol bool) (int, error) {
	if isSymbol {
		return bytesToIntS(b)
	}
	return bytesToIntU(b)
}

const (
	Uint8Max = ^uint8(0)
	Uint16Max = ^uint16(0)
	Uint32Max = ^uint32(0)
	Uint64Max = ^uint64(0)
	Int8Max  = ^uint8(0) >> 1
	Int16Max  = ^uint16(0) >> 1
	Int32Max  = ^uint32(0) >> 1
	Int64Max  = ^uint64(0) >> 1
)

//字节数(大端)组转成int(无符号的)
func bytesToIntU(b []byte) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0}, b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp uint8
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 2:
		var tmp uint16
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 4:
		var tmp uint32
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	default:
		return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}

//字节数(大端)组转成int(有符号)
func bytesToIntS(b []byte) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0}, b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp int8
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 2:
		var tmp int16
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 4:
		var tmp int32
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	default:
		return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}

//IntToBytes
//int转换成byte[]（大端）
//byteNumber 字节数组的数量.
//强制支持 64 位,最大兼容 32 位
func IntToBytes(n int, byteNumber uint8) ([]byte, error) {
	switch byteNumber {
	case 1:
		tmp := int8(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		err := binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), err
	case 2:
		tmp := int16(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		err := binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), err
	case 4:
		tmp := int32(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		err := binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), err
	case 8:
		tmp := int64(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		err := binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), err
	}
	return nil, fmt.Errorf("IntToBytes bit param is invaild")
}

//UintToBytes
//uint转换成byte[]（大端）
//byteNumber 字节数组的数量.
//强制支持 64 位,最大兼容 32 位
func UintToBytes(n uint64, byteNumber uint8) ([]byte, error) {
	switch byteNumber {
	case 1:
		tmp := uint8(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		err := binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), err
	case 2:
		tmp := uint16(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		err := binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), err
	case 4:
		tmp := uint32(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		err := binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), err
	case 8:
		tmp := uint64(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		err := binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), err
	}
	return nil, fmt.Errorf("IntToBytes bit param is invaild")
}
