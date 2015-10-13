package nv

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"
)

func check_array(t *testing.T, field string, value interface{}, l int, fn func([]string)) {
	expecteds := strings.Split(field, ";")
	if len(expecteds) < 2 {
		t.Fatal("field does not seem to be a list of expected values:", field)
	}
	if len(expecteds) != l {
		t.Fatal("length mismatch between expected and decoded arrays, expected:",
			len(expecteds), "decoded:", l)
	}
	fn(expecteds)
}

func check_boolean(t *testing.T, field string, value interface{}) {
	check_boolean_value(t, field, value)
}
func check_boolean_value(t *testing.T, field string, value interface{}) {
	exp := false
	_, err := fmt.Sscanf(field, "%t", &exp)
	if err != nil {
		t.Fatal(err)
	}
	got := value.(bool)
	if got != exp {
		t.Fatal("expected:", exp, "got:", got)
	}
}
func check_boolean_array(t *testing.T, field string, value interface{}) {
	arr := value.([]bool)
	check_array(t, field, value, len(arr), func(expecteds []string) {
		for i := range arr {
			check_boolean_value(t, expecteds[i], arr[i])
		}
	})
}

func check_byte(t *testing.T, field string, value interface{}) {
	d := int(0)
	_, err := fmt.Sscanf(field, "%d", &d)
	if err != nil {
		t.Fatal(err)
	}
	exp := byte(d)
	got := value.(byte)
	if got != exp {
		t.Fatal("expected:", exp, "got:", got)
	}
}
func check_byte_array(t *testing.T, field string, value interface{}) {
	arr := value.([]byte)
	check_array(t, field, value, len(arr), func(expecteds []string) {
		for i := range arr {
			check_byte(t, expecteds[i], arr[i])
		}
	})
}

func check_double(t *testing.T, field string, value interface{}) {
	exp := float64(0)
	// is a hack, but good enough
	_, err := fmt.Sscanf(field, "%f", &exp)
	if err != nil {
		t.Fatal(err)
	}

	got := value.(float64)
	if got != exp {
		t.Fatal("expected:", exp, "got:", got)
	}
}

func check_hrtime(t *testing.T, field string, value interface{}) {
	d := int64(0)
	_, err := fmt.Sscanf(field, "%d", &d)
	if err != nil {
		t.Fatal(err)
	}
	exp := time.Duration(d)
	got := value.(time.Duration)
	if got != exp {
		t.Fatal("expected:", exp, "got:", got)
	}
}

func check_int8(t *testing.T, field string, value interface{}) {
	exp := int8(0)
	_, err := fmt.Sscanf(field, "%d", &exp)
	if err != nil {
		t.Fatal(err)
	}
	got := value.(int8)
	if got != exp {
		t.Fatal("expected:", exp, "got:", got)
	}
}
func check_int8_array(t *testing.T, field string, value interface{}) {
	arr := value.([]int8)
	check_array(t, field, value, len(arr), func(expecteds []string) {
		for i := range arr {
			check_int8(t, expecteds[i], arr[i])
		}
	})
}

func check_int16(t *testing.T, field string, value interface{}) {
	exp := int16(0)
	_, err := fmt.Sscanf(field, "%d", &exp)
	if err != nil {
		t.Fatal(err)
	}
	got := value.(int16)
	if got != exp {
		t.Fatal("expected:", exp, "got:", got)
	}
}
func check_int16_array(t *testing.T, field string, value interface{}) {
	arr := value.([]int16)
	check_array(t, field, value, len(arr), func(expecteds []string) {
		for i := range arr {
			check_int16(t, expecteds[i], arr[i])
		}
	})
}

func check_int32(t *testing.T, field string, value interface{}) {
	exp := int32(0)
	_, err := fmt.Sscanf(field, "%d", &exp)
	if err != nil {
		t.Fatal(err)
	}
	got := value.(int32)
	if got != exp {
		t.Fatal("expected:", exp, "got:", got)
	}
}
func check_int32_array(t *testing.T, field string, value interface{}) {
	arr := value.([]int32)
	check_array(t, field, value, len(arr), func(expecteds []string) {
		for i := range arr {
			check_int32(t, expecteds[i], arr[i])
		}
	})
}

func check_int64(t *testing.T, field string, value interface{}) {
	exp := int64(0)
	_, err := fmt.Sscanf(field, "%d", &exp)
	if err != nil {
		t.Fatal(err)
	}
	got := value.(int64)
	if got != exp {
		t.Fatal("expected:", exp, "got:", got)
	}
}
func check_int64_array(t *testing.T, field string, value interface{}) {
	arr := value.([]int64)
	check_array(t, field, value, len(arr), func(expecteds []string) {
		for i := range arr {
			check_int64(t, expecteds[i], arr[i])
		}
	})
}

func check_uint8(t *testing.T, field string, value interface{}) {
	exp := uint8(0)
	_, err := fmt.Sscanf(field, "%d", &exp)
	if err != nil {
		t.Fatal(err)
	}
	got := value.(uint8)
	if got != exp {
		t.Fatal("expected:", exp, "got:", got)
	}
}
func check_uint8_array(t *testing.T, field string, value interface{}) {
	arr := value.([]uint8)
	check_array(t, field, value, len(arr), func(expecteds []string) {
		for i := range arr {
			check_uint8(t, expecteds[i], arr[i])
		}
	})
}

func check_uint16(t *testing.T, field string, value interface{}) {
	exp := uint16(0)
	_, err := fmt.Sscanf(field, "%d", &exp)
	if err != nil {
		t.Fatal(err)
	}
	got := value.(uint16)
	if got != exp {
		t.Fatal("expected:", exp, "got:", got)
	}
}
func check_uint16_array(t *testing.T, field string, value interface{}) {
	arr := value.([]uint16)
	check_array(t, field, value, len(arr), func(expecteds []string) {
		for i := range arr {
			check_uint16(t, expecteds[i], arr[i])
		}
	})
}

func check_uint32(t *testing.T, field string, value interface{}) {
	exp := uint32(0)
	_, err := fmt.Sscanf(field, "%d", &exp)
	if err != nil {
		t.Fatal(err)
	}
	got := value.(uint32)
	if got != exp {
		t.Fatal("expected:", exp, "got:", got)
	}
}
func check_uint32_array(t *testing.T, field string, value interface{}) {
	arr := value.([]uint32)
	check_array(t, field, value, len(arr), func(expecteds []string) {
		for i := range arr {
			check_uint32(t, expecteds[i], arr[i])
		}
	})
}

func check_uint64(t *testing.T, field string, value interface{}) {
	exp := uint64(0)
	_, err := fmt.Sscanf(field, "%d", &exp)
	if err != nil {
		t.Fatal(err)
	}
	got := value.(uint64)
	if got != exp {
		t.Fatal("expected:", exp, "got:", got)
	}
}
func check_uint64_array(t *testing.T, field string, value interface{}) {
	arr := value.([]uint64)
	check_array(t, field, value, len(arr), func(expecteds []string) {
		for i := range arr {
			check_uint64(t, expecteds[i], arr[i])
		}
	})
}

func check_nvlist(t *testing.T, field string, value interface{}) {
	list := value.(map[string]interface{})
	for f, v := range list {
		checkers[reflect.TypeOf(v).String()](t, f, v)
	}
}
func check_nvlist_array(t *testing.T, field string, value interface{}) {
	arr := value.([]map[string]interface{})
	check_array(t, field, value, len(arr), func(expecteds []string) {
		for i := range arr {
			check_nvlist(t, expecteds[i], arr[i])
		}
	})
}

func check_string(t *testing.T, field string, value interface{}) {
	exp := field
	got := value.(string)
	if got != exp {
		t.Fatal("expected:", exp, "got:", got)
	}
}
func check_string_array(t *testing.T, field string, value interface{}) {
	arr := value.([]string)
	check_array(t, field, value, len(arr), func(expecteds []string) {
		for i := range arr {
			check_string(t, expecteds[i], arr[i])
		}
	})
}

var checkers map[string]func(*testing.T, string, interface{})

func init() {
	checkers = map[string]func(*testing.T, string, interface{}){
		"bool":                      check_boolean,
		"byte":                      check_byte,
		"int16":                     check_int16,
		"uint16":                    check_uint16,
		"int32":                     check_int32,
		"uint32":                    check_uint32,
		"int64":                     check_int64,
		"uint64":                    check_uint64,
		"string":                    check_string,
		"[]byte":                    check_byte_array,
		"[]int16":                   check_int16_array,
		"[]uint16":                  check_uint16_array,
		"[]int32":                   check_int32_array,
		"[]uint32":                  check_uint32_array,
		"[]int64":                   check_int64_array,
		"[]uint64":                  check_uint64_array,
		"[]string":                  check_string_array,
		"time.Duration":             check_hrtime,
		"map[string]interface {}":   check_nvlist,
		"[]map[string]interface {}": check_nvlist_array,
		"int8":    check_int8,
		"uint8":   check_uint8,
		"[]bool":  check_boolean_array,
		"[]int8":  check_int8_array,
		"[]uint8": check_uint8_array,
		"float64": check_double,
	}
}

func assertFields(t *testing.T, name string, m map[string]interface{}) {
	for field, value := range m {
		valueType := reflect.TypeOf(value).String()
		fn, ok := checkers[valueType]
		if !ok {
			t.Fatal(name, "unknown type:", valueType)
		}
		fn(t, field, value)
	}
}

//go:generate make -C _test ../known_good_data_test.go
func TestDecodeGood(t *testing.T) {
	for _, test := range good {
		t.Log(test.name)

		m := map[string]interface{}{}
		err := Decode(test.payload, &m)
		if err != nil {
			t.Fatal(test.name, "failed:", err)
		}

		t.Log(l)
		assertFields(t, test.name, l)
	}
}

func TestDecodeBad(t *testing.T) {
	for _, test := range decode_bad {
		t.Log(test.err)

		m := map[string]interface{}{}
		err := Decode(test.payload, &m)
		if err == nil {
			t.Fatalf("expected an error, wanted:|%s| payload:|%v|\n",
				test.err, test.payload)
		}
		if test.err != err.Error() {
			t.Fatalf("error mismatch, want:|%s| got:|%s| payload:|%v|\n",
				test.err, err.Error(), test.payload)
		}
	}
}
