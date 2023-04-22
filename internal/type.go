package internal

import (
	"fmt"
	"reflect"
)

func (s *Settings) CheckTypeMatch(f Field, value interface{}) (bool, error) {
	fieldType, ok := s.TypeMap[f.Type]
	if !ok {
		return false, fmt.Errorf("unknown type: %s", f.Type)
	}

	if reflect.TypeOf(value) != fieldType {
		return false, fmt.Errorf("type mismatch: expected %v, got %v", fieldType, reflect.TypeOf(value))
	}

	return true, nil
}
