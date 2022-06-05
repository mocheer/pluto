package fn

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//ToStruct 用map填充结构
func Map2Struct(m map[string]interface{}, s struct{}) error {
	for k, v := range m {
		err := SetStructValue(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

// SetStructValue 设置结构体的值
func SetStructValue(obj any, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem() //结构体属性值
	// structFieldValue := structValue.FieldByName(name) //结构体单个属性值
	structFieldValue := structValue.FieldByNameFunc(func(s string) bool { //临时处理，待完善
		return strings.ToUpper(s) == strings.ToUpper(strings.ReplaceAll(name, "_", ""))
	})
	if !structFieldValue.IsValid() {
		return fmt.Errorf("没有这个字段: %s", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("无法设置值：%s", name)
	}
	//
	structFieldType := structFieldValue.Type() //结构体的类型
	val := reflect.ValueOf(value)              //map值的反射值

	var err error
	if structFieldType != val.Type() { //结构体字段类型和想要赋的值类型不一致，进行类型转换
		val, err = fromString(fmt.Sprintf("%v", value), structFieldValue.Type().Name()) //类型转换
		if err != nil {
			return err
		}
	}
	structFieldValue.Set(val)
	return nil
}

//fromString 类型转换:string类型转其他类型
func fromString(value string, ntype string) (reflect.Value, error) {
	if ntype == "string" {
		return reflect.ValueOf(value), nil
	} else if ntype == "float64" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(i), err
	} else if ntype == "int64" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(i), err
	} else if ntype == "float32" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(float32(i)), err
	} else if ntype == "int32" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int64(i)), err
	} else if ntype == "int8" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int8(i)), err
	} else if ntype == "int" {
		i, err := strconv.Atoi(value)
		return reflect.ValueOf(i), err
	} else if ntype == "time.Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "uint" {
		i, err := strconv.ParseUint(value, 10, 0)
		return reflect.ValueOf(uint(i)), err
	}

	return reflect.ValueOf(value), errors.New("未知的类型：" + ntype)
}
