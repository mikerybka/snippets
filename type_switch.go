package snippets

import "reflect"

func TypeSwitch(t reflect.Type) {
	switch t.Kind() {
	case reflect.Bool:
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
	case reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
	case reflect.Array, reflect.Slice:
	case reflect.Map:
	case reflect.Pointer:
	case reflect.String:
	case reflect.Struct:
	case reflect.Chan:
	case reflect.Func:
	case reflect.Interface:
	case reflect.UnsafePointer:
	default:
	}
}
