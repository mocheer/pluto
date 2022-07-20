package fn

// javascript : (value = +value) >= value)
// NaN = false
// undefined = false
// number = true
// "number" = true //待完成
// Date =  true

func IsValidNumber(value any) bool {
	return value != nil
}
