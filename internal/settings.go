package internal

import "reflect"

type Settings struct {
	Databases map[string]Database
	Location  string
	TypeMap   map[string]reflect.Type
}

func Init(path string) (Settings, error) {
	var set Settings

	if path == "" {
		set.Location = ".test_folder/"
	} else {
		set.Location = path
	}

	err := set.prepareDb()
	if err != nil {
		return set, err
	}

	set.TypeMap = map[string]reflect.Type{
		"string":  reflect.TypeOf(""),
		"int":     reflect.TypeOf(int(0)),
		"int8":    reflect.TypeOf(int8(0)),
		"int16":   reflect.TypeOf(int16(0)),
		"int32":   reflect.TypeOf(int32(0)),
		"int64":   reflect.TypeOf(int64(0)),
		"float":   reflect.TypeOf(float64(0)),
		"float32": reflect.TypeOf(float32(0)),
		"float64": reflect.TypeOf(float64(0)),
		"bool":    reflect.TypeOf(false),
	}

	return set, nil
}
