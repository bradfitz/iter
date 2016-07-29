// Package iter provides a syntantically different way to iterate over integers. That's it.
package iter

import "reflect"

// N returns a slice of n 0-sized elements, suitable for ranging over.
// n is determined "appropriately" depending on the type and contents of v.
//
// For example:
//
//    for i := range iter.N(10) {
//        fmt.Println(i)
//    }
//
// ... will print 0 to 9, inclusive.
//
// It causes one allocation.
func N(v interface{}) []struct{} {
	return make([]struct{}, n(v))
}

type Inter interface {
	Int() int
}

func n(any interface{}) int {
	if v, ok := any.(Inter); ok {
		return v.Int()
	}

	switch v := any.(type) {
	case int:
		return v
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case uintptr:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)

	case complex64:
		return int(real(v))
	case complex128:
		return int(real(v))

	case string:
		return len(v)

	case bool:
		if v {
			return 1
		}
		return 0
	}

	switch v := reflect.ValueOf(any); v.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return v.Len()
	case reflect.Ptr, reflect.Interface:
		return n(v.Elem().Interface())
	case reflect.Struct:
		if f := v.FieldByName("N"); f.IsValid() {
			return n(f.Interface())
		}
	}

	return 0
}
