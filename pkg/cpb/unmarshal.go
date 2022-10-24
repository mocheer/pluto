package cpb

// 暂时弃用，版本没跟上
// func Unmarshal(data []byte) string {
// 	v, n := protowire.ConsumeVarint(data)
// 	return unmarshal(data, v, n)
// }

// func unmarshal(data []byte, v uint64, n int) string {
// 	var builder strings.Builder
// 	switch v {
// 	case Struct:
// 		res := unmarshalStruct(data, v, n)
// 		builder.WriteString(res)
// 	case Int:
// 		res, _ := unmarshalInt(data, v, n)
// 		builder.WriteString(res)
// 	case String:
// 		res, _ := unmarshalString(data, v, n)
// 		builder.WriteString(res)
// 	case MapString:
// 		res, _ := unmarshalMapStringString(data, v, n)
// 		builder.WriteString(res)
// 	}
// 	return builder.String()
// }

// func unmarshalStruct(data []byte, v uint64, n int) string {
// 	var builder strings.Builder
// 	num, n1 := protowire.ConsumeVarint(data[n:])
// 	n = n + n1
// 	val := ""
// 	builder.WriteRune('{')
// 	for i := uint64(0); i < num; i++ {
// 		v1, n2 := protowire.ConsumeVarint(data[n:])
// 		n = n + n2
// 		if v1 == Invalid {
// 			return ""
// 		}
// 		val, n = unmarshalString(data, v, n) // "name"
// 		builder.WriteString(val)
// 		builder.WriteRune(':')
// 		switch v1 {
// 		case Int:
// 			val, n = unmarshalInt(data, v, n)
// 			builder.WriteString(val)
// 		case String:
// 			val, n = unmarshalString(data, v, n)
// 			builder.WriteString(val)
// 		case MapString:
// 			val, n = unmarshalMapStringString(data, v, n)
// 			builder.WriteString(val)
// 		case Bytes: //不方便直接转成js中的Uint8Array

// 		default:
// 			panic("invalid type")
// 		}
// 		if i < num-1 {
// 			builder.WriteByte(',')
// 		}
// 	}
// 	builder.WriteRune('}')
// 	return builder.String()
// }

// var quota = "\""

// func unmarshalString(data []byte, v uint64, n int) (string, int) {
// 	v1, n1 := protowire.ConsumeString(data[n:])
// 	return quota + v1 + quota, n + n1
// }

// func unmarshalInt(data []byte, v uint64, n int) (string, int) {
// 	v1, n1 := protowire.ConsumeVarint(data[n:])
// 	return strconv.FormatUint(v1, 10), n + n1
// }

// func unmarshalMapStringString(data []byte, v uint64, n int) (string, int) {
// 	v1, n1 := protowire.ConsumeVarint(data[n:])
// 	n = n + n1
// 	var builder strings.Builder
// 	builder.WriteRune('{')
// 	for i := uint64(0); i < v1; i++ {
// 		v2, n2 := protowire.ConsumeString(data[n:])
// 		n = n + n2
// 		builder.WriteRune('"')
// 		builder.WriteString(v2)
// 		builder.WriteRune('"')
// 		builder.WriteRune(':')
// 		v3, n3 := protowire.ConsumeString(data[n:])
// 		n = n + n3
// 		builder.WriteRune('"')
// 		builder.WriteString(v3)
// 		builder.WriteRune('"')
// 		if i < v1-1 {
// 			builder.WriteRune(',')
// 		}
// 	}
// 	builder.WriteRune('}')
// 	return builder.String(), n
// }
