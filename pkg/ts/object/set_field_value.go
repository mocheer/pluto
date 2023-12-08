package object

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// SetFieldValue 设置结构体的值
// 当类型不一致的时候，强转失败不会修改结构体的值，并返回nil
func SetFieldValue(obj any, name string, value any) (fieldValue reflect.Value, err error) {
	fieldValue = FindFieldValue(obj, name, "json")
	// 没有字段，或者无法设置值
	if fieldValue.IsValid() && fieldValue.CanSet() {
		//
		fieldType := fieldValue.Type() //结构体的类型
		rVal := reflect.ValueOf(value) //反射值
		if fieldType != rVal.Type() {  //结构体字段类型和想要赋的值类型不一致，进行类型转换
			ntype := fieldType.Kind()
			if ntype != reflect.Interface { //如果原始字段类型是any，不需要做任何转换
				// 进行类型转换
				var cvtVal any
				// TODO：优化性能，这里不应该转换成字符串
				cvtVal, err = fromString(fmt.Sprintf("%v", value), ntype) //类型转换
				if err == nil {
					rVal = reflect.ValueOf(cvtVal)
				}
			}
		}
		fieldValue.Set(rVal)
	} else {
		err = fmt.Errorf("没有字段或者无法设置值:%s", name)
	}
	return
}

// fromString 类型转换:string类型转其他类型
// 支持类型：string、float64、float32、int、int64、int32、int16、int8、uint、uint64、uint32、uint16、uint8、bool
func fromString(value string, ntype reflect.Kind) (any, error) {
	switch ntype {
	case reflect.String:
		return value, nil
	case reflect.Float64:
		return strconv.ParseFloat(value, 64)
	case reflect.Float32:
		i, err := strconv.ParseFloat(value, 32)
		return float32(i), err
	case reflect.Int64:
		return strconv.ParseInt(value, 10, 64)
	case reflect.Int32:
		i, err := strconv.ParseInt(value, 10, 32)
		return int32(i), err
	case reflect.Int16:
		i, err := strconv.ParseInt(value, 10, 16)
		return int16(i), err
	case reflect.Int8:
		i, err := strconv.ParseInt(value, 10, 8)
		return int8(i), err
	case reflect.Int:
		return strconv.Atoi(value)
	case reflect.Uint64:
		return strconv.ParseUint(value, 10, 64)
	case reflect.Uint32:
		i, err := strconv.ParseUint(value, 10, 32)
		return uint32(i), err
	case reflect.Uint8:
		i, err := strconv.ParseUint(value, 10, 8)
		return uint8(i), err
	case reflect.Uint:
		i, err := strconv.ParseUint(value, 10, 0)
		return uint(i), err
	case reflect.Bool:
		return value == "true", nil
	case reflect.Array:
		return nil, nil
	case reflect.Map:
		return nil, nil
	case reflect.Slice:
		return nil, nil
	default:
		// 未知类型
		return value, errors.New("未知的类型：" + ntype.String())
	}

}
