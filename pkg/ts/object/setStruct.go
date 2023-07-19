package object

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// SetStructValue 设置结构体的值
func SetStructValue(obj any, name string, value any) error {
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
		ntype := structFieldValue.Type().Kind()
		if ntype != reflect.Interface {
			val, err = fromString(fmt.Sprintf("%v", value), ntype) //类型转换
			if err != nil {
				return err
			}
		}
	}
	structFieldValue.Set(val)
	return nil
}

// fromString 类型转换:string类型转其他类型
func fromString(value string, ntype reflect.Kind) (reflect.Value, error) {
	switch ntype {
	case reflect.String:
		return reflect.ValueOf(value), nil
	case reflect.Float64:
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(i), err
	case reflect.Int64:
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(i), err
	case reflect.Float32:
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(float32(i)), err
	case reflect.Int32:
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int32(i)), err
	case reflect.Int8:
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int8(i)), err
	case reflect.Int:
		i, err := strconv.Atoi(value)
		return reflect.ValueOf(i), err
	case reflect.Uint:
		i, err := strconv.ParseUint(value, 10, 0)
		return reflect.ValueOf(uint(i)), err
	}

	return reflect.ValueOf(value), errors.New("未知的类型：" + ntype.String())
}
