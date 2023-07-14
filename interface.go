package mapstructure

import "reflect"

type Unmarshaler interface {
	Unmarshal(from reflect.Value) error
}
