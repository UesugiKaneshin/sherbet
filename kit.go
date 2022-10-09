package sherbet

import (
	"crypto/rand"
	"math/big"
	"reflect"
	"time"

	"github.com/viant/toolbox"
)

// EncryptArrayToInt encrypt array to int64
func EncryptArrayToInt(array []bool) int64 {
	var result int64 = 0

	for _, val := range array {
		result = result << 1
		if val {
			result++
		} else {
		}
	}

	return result
}

// DecryptArrayToInt decrypt int64 to array
func DecryptArrayToInt(data int64, length int) []bool {
	var result = make([]bool, length)

	for length > 0 {
		result[length-1] = data%2 == 1
		length--
		data = data >> 1
	}

	return result
}

// DatetimeFormat format a datetime
func DatetimeFormat(datetime *string, format string) error {
	var result error = nil

	temp, err := time.Parse(time.RFC3339Nano, *datetime)
	if err == nil {
		*datetime = temp.Format(format)
	} else {
		result = err
	}

	return result
}

// DatetimesFormat format datetimes
func DatetimesFormat(datetimes *[]*string, format string) error {
	var result error = nil

	for _, val := range *datetimes {
		if result == nil {
			result = DatetimeFormat(val, format)
		} else {
			break
		}
	}

	return result
}

// ReflectTags get object's tags
func ReflectTags(obj interface{}, tag string) *[]string {
	var result = new([]string)

	types := reflect.TypeOf(obj)
	reflectTags(types, tag, result)

	return result
}
func reflectTags(types reflect.Type, tag string, data *[]string) {
	for i := 0; i < types.NumField(); i++ {
		if types.Field(i).Tag.Get(tag) == "" {
			reflectTags(types.Field(i).Type, tag, data)
		} else {
			tag := types.Field(i).Tag.Get(tag)

			if !toolbox.HasSliceAnyElements(*data, tag) {
				*data = append(*data, tag)
			} else {
			}
		}
	}
}

// ReflectValues get object's value
func ReflectValues(obj interface{}) *[]interface{} {
	var result = new([]interface{})

	reflectValues(obj, result)

	return result
}
func reflectValues(obj interface{}, data *[]interface{}) {
	values := reflect.ValueOf(obj)
	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).Kind() == reflect.Struct {
			reflectValues(values.Field(i).Interface(), data)
		} else {
			switch values.Field(i).Elem().Kind() {
			case reflect.Bool:
				*data = append(*data, values.Field(i).Elem().Bool())
			case reflect.Int:
				*data = append(*data, values.Field(i).Elem().Int())
			case reflect.Uint8:
				*data = append(*data, values.Field(i).Elem().Uint())
			case reflect.String:
				*data = append(*data, values.Field(i).Elem().String())
			case reflect.Float64:
				*data = append(*data, values.Field(i).Elem().Float())
			default:
				*data = append(*data, nil)
			}
		}
	}
}

// HasValueFromSliceForInt has value from slice
func HasValueFromSliceForInt(array *[]*int, value *int) (result bool) {
	result = false

	for _, val := range *array {
		if *value == *val {
			result = true
		} else {
		}
	}

	return result
}

// HasValueFromSliceForString has value from slice
func HasValueFromSliceForString(array *[]*string, value *string) (result bool) {
	result = false

	for _, val := range *array {
		if *value == *val {
			result = true
		} else {
		}
	}

	return result
}

// HasValueFromSlice has value from slice
func HasValueFromSlice[T int | float64](array *[]*T, value *T) (result bool) {
	result = false

	for _, val := range *array {
		if *value == *val {
			result = true
		} else {
		}
	}

	return result
}

// GetPoint get object's point
func GetPoint[T string | int | int64 | float64](obj T) *T { return &obj }

// GetRandomPassword get password
func GetRandomPassword(length int, kind string) string {
	var password = make([]rune, length)
	var codeModel = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+=-!@#$%*,.[]")

	for key := range password {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(codeModel))))
		password[key] = codeModel[int(index.Int64())]
	}

	return string(password)
}
