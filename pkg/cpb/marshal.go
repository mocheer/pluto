package cpb

import (
	"math"
	"reflect"

	"github.com/mocheer/pluto/pkg/ts/object"
	"google.golang.org/protobuf/encoding/protowire"
)

// Marshal
// 目前只计划支持 struct{} 和 []struct{}
// TODO 用gorm/gen+代码生成器直接生成proto编码程序，取消反射
// TODO []any 数组如果是接口会出问题
func Marshal(s any) []byte {
	var data []byte
	data = marshal(data, reflect.ValueOf(s))
	return data
}

// marshal
func marshal(data []byte, v reflect.Value) []byte {
	// 指针、结构体和其属性可能是空对象、零值，当指针为nil时，后续取值，获取属性时会报错
	// 关于 IsValid
	// 1. 当 v 的原始对象是一个any的类型且为空，IsZero 会抛出异常，这个时候 IsValid=false ，可以先验证 IsValid
	// 2. 当 v 是一个结构体或者其他不能为空的值，IsNil 会发生错误，所以在用 IsNil 进行判断前需要先验证 IsValid
	// || v.IsNil()
	if !v.IsValid() {

		return protowire.AppendVarint(data, TypeInvalid)
	}
	//
	typ := v.Type()
	// 结构体中未赋值的属性值，为真
	// 即使赋值，但如果值为空字符串、nil、空指针、数字0等零值时，为真
	// 这里在零值的时候，需要和json保持一致，这里其实不需要再处理，因为已经排除了nil情况
	// if v.IsZero() {
	// 	switch typ.Kind() {
	// 	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
	// 		data = protowire.AppendVarint(data, TypeInt)
	// 		data = protowire.AppendVarint(data, 0)
	// 	case reflect.String:
	// 		data = protowire.AppendVarint(data, TypeString)
	// 		data = protowire.AppendString(data, "")
	// 	default:
	// 		data = protowire.AppendVarint(data, TypeInvalid)
	// 	}
	// 	return data
	// }

	//
	//
	// time.Time、datatypes.JSON
	// 比如说 delete_at 时间类型
	// json.Marshaler 这里不用，因为会多出引号
	// encoding.TextMarshaler 也不用，因为 datatypes.JSON 没有实现该方法，但问题是即使是json前端还需要序列化
	// 这里考虑实现自定义接口 MarshalCPB
	if typ.Implements(marshalerType) {
		vi := v.Interface()
		// typ.Implements(marshalerType)，所以这里能够确保转换正确
		s, _ := vi.(Marshaler)
		//
		cv, ct := s.MarshalCPB()
		data = protowire.AppendVarint(data, ct)
		switch ct {
		case TypeInt:
			val, _ := cv.(uint64)
			data = protowire.AppendVarint(data, val)
		case TypeJSON:
			val, _ := cv.(string)
			data = protowire.AppendString(data, val)
		case TypeString:
			val, _ := cv.(string)
			data = protowire.AppendString(data, val)
		}
		return data
	}
	//
	switch typ.Kind() {
	case reflect.Struct:
		// gorm.DeletedAt
		var isZero = true
		for i := 0; i < v.NumField(); i++ {
			if !v.Field(i).IsZero() {
				isZero = false
			}
		}
		if isZero {
			return protowire.AppendVarint(data, TypeInvalid)
		}
		data = protowire.AppendVarint(data, TypeStruct)
		data = marshalStruct(v, typ, data)
		//
	case reflect.Map:
		num := v.Len()
		if num > 0 {
			data = protowire.AppendVarint(data, TypeMapString)
			data = protowire.AppendVarint(data, uint64(num))
			//
			keys := v.MapKeys()
			switch v.MapIndex(keys[0]).Kind() {
			case reflect.String:
				data = protowire.AppendVarint(data, TypeString)
				for _, key := range keys {
					data = protowire.AppendString(data, key.String())
					data = protowire.AppendString(data, v.MapIndex(key).String())
				}
			case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
				data = protowire.AppendVarint(data, TypeInt)
				for _, key := range keys {
					data = protowire.AppendString(data, key.String())
					data = protowire.AppendVarint(data, v.MapIndex(key).Uint())
				}
			}
		} else {
			data = protowire.AppendVarint(data, TypeInvalid)
		}
		// 数组和切片
	case reflect.Array:

		fallthrough
	case reflect.Slice:
		num := v.Len()
		if num > 0 {
			vi := v.Index(0)
			vitype := vi.Kind()
			if vitype == reflect.Pointer {
				vi = vi.Elem()
				vitype = vi.Kind()
			}
			switch vitype {
			case reflect.Uint8:
				data = protowire.AppendVarint(data, TypeBytes)
				data = protowire.AppendBytes(data, v.Bytes())
			case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64:
				data = protowire.AppendVarint(data, TypeSliceInt)
				data = protowire.AppendVarint(data, uint64(num))
				for i := 0; i < num; i++ {
					data = protowire.AppendVarint(data, uint64(v.Index(i).Int()))
				}
			case reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
				data = protowire.AppendVarint(data, TypeSliceInt)
				data = protowire.AppendVarint(data, uint64(num))
				for i := 0; i < num; i++ {
					data = protowire.AppendVarint(data, uint64(v.Index(i).Uint()))
				}
			case reflect.String:
				data = protowire.AppendVarint(data, TypeSliceString)
				data = protowire.AppendVarint(data, uint64(num))
				for i := 0; i < num; i++ {
					data = protowire.AppendString(data, v.Index(i).String())
				}
			case reflect.Struct:
				data = protowire.AppendVarint(data, TypeSliceStruct)
				data = marshalSliceStruct(v, vi.Type(), data)

			case reflect.Slice:

				vi2 := vi.Index(0)
				vitype2 := vi2.Kind()
				if vitype2 == reflect.Pointer {
					vi2 = vi.Elem()
					vitype2 = vi.Kind()
				}
				if vitype2 == reflect.Uint8 {
					data = protowire.AppendVarint(data, TypeSliceBytes)
					data = protowire.AppendVarint(data, uint64(num))
					for i := 0; i < num; i++ {
						data = protowire.AppendBytes(data, v.Index(i).Bytes())
					}
				}

			}

		} else { //需要给出类型，否则不好解析
			data = protowire.AppendVarint(data, TypeInvalid)
		}
	case reflect.String:
		data = protowire.AppendVarint(data, TypeString)
		data = protowire.AppendString(data, v.String())
	case reflect.Bool:
		data = protowire.AppendVarint(data, TypeBool)
		data = protowire.AppendVarint(data, protowire.EncodeBool(v.Bool()))
	case reflect.Float32:
		data = protowire.AppendVarint(data, TypeFixed32)
		data = protowire.AppendFixed32(data, math.Float32bits(float32(v.Float())))
	case reflect.Float64:
		data = protowire.AppendVarint(data, TypeFixed64)
		data = protowire.AppendFixed64(data, math.Float64bits(v.Float()))
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64:
		data = protowire.AppendVarint(data, TypeInt)
		data = protowire.AppendVarint(data, uint64(v.Int()))
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
		data = protowire.AppendVarint(data, TypeInt)
		data = protowire.AppendVarint(data, uint64(v.Uint()))
	case reflect.Ptr:
		if v.IsNil() {
			return protowire.AppendVarint(data, TypeInvalid)
		}
		data = marshal(data, v.Elem())
	case reflect.Interface: //
		data = marshal(data, reflect.ValueOf(v.Interface()))
	default:
		data = protowire.AppendVarint(data, TypeInvalid)
	}
	return data
}

// marshalType 解析类型
// func marshalType(data []byte, v reflect.Value) []byte {
// 	return nil
// }

// marshalStruct
func marshalStruct(v reflect.Value, typ reflect.Type, data []byte) []byte {
	fieldIndexs, fieldNames := object.GetFieldsByReflectStruct(typ)
	num := len(fieldIndexs)
	data = protowire.AppendVarint(data, uint64(num))
	for i := 0; i < num; i++ {
		// TODO 这里需要判断是不是组合后的struct字段，然后平级输出
		data = protowire.AppendString(data, fieldNames[i])
		data = marshal(data, v.Field(fieldIndexs[i]))
	}
	return data
}

// marshalSliceStruct
func marshalSliceStruct(v reflect.Value, typ reflect.Type, data []byte) []byte {
	//
	num := v.Len()
	data = protowire.AppendVarint(data, uint64(num))
	//
	fieldIndexs, fieldNames := object.GetFieldsByReflectStruct(typ)
	numField := len(fieldIndexs)
	data = protowire.AppendVarint(data, uint64(numField))
	//
	for k := 0; k < numField; k++ {
		data = protowire.AppendString(data, fieldNames[k])
	}
	//
	for i := 0; i < num; i++ {
		val := v.Index(i)
		// 这里一般只需要判断一次，不需要判断这么多次
		if val.Kind() == reflect.Pointer {
			val = val.Elem()
		}
		for j := 0; j < numField; j++ {
			data = marshal(data, val.Field(fieldIndexs[j]))
		}
	}
	return data
}
