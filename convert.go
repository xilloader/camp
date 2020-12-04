package camp

import (
	"encoding/binary"
	"fmt"
	"reflect"
)

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// indirect returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil).
func indirect(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		// Avoid creating a reflect.Value if it's not a pointer.
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// indirectToStringerOrError returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil) or an implementation of fmt.Stringer
// or error,
func indirectToStringerOrError(a interface{}) interface{} {
	if a == nil {
		return nil
	}

	var errorType = reflect.TypeOf((*error)(nil)).Elem()
	var fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	v := reflect.ValueOf(a)
	for !v.Type().Implements(fmtStringerType) && !v.Type().Implements(errorType) && v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// ParseBool returns the boolean value represented by the string.
//
// It accepts 1, 1.0, t, T, TRUE, true, True, YES, yes, Yes,Y, y, ON, on, On,
// 0, 0.0, f, F, FALSE, false, False, NO, no, No, N,n, OFF, off, Off.
// Any other value returns an error.
func ParseBool(val interface{}) (value bool, err error) {
	if val != nil {
		switch v := val.(type) {
		case bool:
			return v, nil
		case string:
			switch v {
			case "1", "t", "T", "true", "TRUE", "True", "YES", "yes", "Yes", "Y", "y", "ON", "on", "On":
				return true, nil
			case "0", "f", "F", "false", "FALSE", "False", "NO", "no", "No", "N", "n", "OFF", "off", "Off":
				return false, nil
			}
		case int8, int32, int64:
			strV := fmt.Sprintf("%d", v)
			if strV == "1" {
				return true, nil
			} else if strV == "0" {
				return false, nil
			}
		case float64:
			if v == 1.0 {
				return true, nil
			} else if v == 0.0 {
				return false, nil
			}
		}
		return false, fmt.Errorf("parsing %q: invalid syntax", val)
	}
	return false, fmt.Errorf("parsing <nil>: invalid syntax")
}

func IntToBytesE(i interface{}) ([]byte, error) {
	i = indirect(i)

	switch s := i.(type) {
	case int:
		return Int64ToBytes(int64(s)), nil
	case uint:
		return Int64ToBytes(int64(s)), nil
	case int64:
		return Int64ToBytes(s), nil
	case uint64:
		return Int64ToBytes(int64(s)), nil
	case int32:
		return Int32ToBytes(s), nil
	case uint32:
		return Int32ToBytes(int32(s)), nil
	case int16:
		return Int16ToBytes(s), nil
	case int8:
		return Int16ToBytes(int16(s)), nil
	case uint16:
		return Int16ToBytes(int16(s)), nil
	case uint8:
		return Int16ToBytes(int16(s)), nil
	default:
		return nil, fmt.Errorf("%#v of type %T is not an integer", i, i)
	}
}

func IntToBytes(i interface{}) []byte {
	v, _ := IntToBytesE(i)
	return v
}

// int64 转 []byte
func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

// []byte 转 int64
func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

// int32 转 []byte
func Int32ToBytes(i int32) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

// []byte 转 int32
func BytesToInt32(buf []byte) int64 {
	return int64(binary.BigEndian.Uint32(buf))
}

// int16 转 []byte
func Int16ToBytes(i int16) []byte {
	var buf = make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(i))
	return buf
}

// []byte 转 int16
func BytesToInt16(buf []byte) int16 {
	return int16(binary.BigEndian.Uint16(buf))
}
