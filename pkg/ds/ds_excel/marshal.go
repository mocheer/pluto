package ds_excel

import (
	"fmt"
	"reflect"

	"github.com/mocheer/pluto/pkg/ts/object"
	"github.com/xuri/excelize/v2"
)

type Marshaler interface {
	MarshalCPB() (any, uint64)
}

var marshalerType = reflect.TypeOf((*Marshaler)(nil)).Elem()

// Marshal
func Marshal(data any) ([]byte, error) {
	f := excelize.NewFile()
	//
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	count := v.Len()
	if count > 0 {
		// 这里反射只执行一次（除非切片元素是不同的对象）
		val := v.Index(0)
		fieldIndexs, fieldNames := object.GetStructFields(val.Type())
		num := len(fieldIndexs)

		// A = 65
		// 设置宽度
		f.SetColWidth("Sheet1", string(rune(65)), string(rune(65+num)), 20)
		for i := 0; i < num; i++ {
			f.SetCellValue("Sheet1", string(rune(65+i))+"1", fieldNames[i])
		}
		for rowIndex := 0; rowIndex < count; rowIndex++ {
			item := v.Index(rowIndex)
			for i := 0; i < num; i++ {
				val := item.Field(i)
				//
				if val.Type().Implements(marshalerType) {
					vi := val.Interface()
					c := vi.(Marshaler)
					data, cpbtype := c.MarshalCPB()
					switch cpbtype {
					// case cpb.TypeDate:
					// 	f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), rowIndex+2), clock.FromUnixMilli(int64(data.(uint64))).Fmt(clock.FmtFullDate))
					default:
						f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), rowIndex+2), data)
					}
				} else {
					f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), rowIndex+2), val)
				}

			}
		}
	}

	bf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return bf.Bytes(), nil
}
