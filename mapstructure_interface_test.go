package mapstructure

import (
	"fmt"
	"reflect"
	"testing"
)

type Listable[T any] []T

func (l *Listable[T]) Unmarshal(from reflect.Value) error {
	data := from.Interface()
	if item, ok := data.(T); ok {
		*l = []T{item}
		return nil
	}
	if items, ok := data.([]T); ok {
		*l = items
		return nil
	}
	return fmt.Errorf("invalid type: %s", from.Type().String())
}

func TestInterfaceDecode(t *testing.T) {
	m := map[string]any{
		"list": "a",
	}

	type A struct {
		List Listable[string] `config:"list"`
	}

	var a A

	decoder, err := NewDecoder(&DecoderConfig{
		DecodeHook: UnmarshalInterfaceHookFunc(),
		TagName:    "config",
		Result:     &a,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = decoder.Decode(m)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(a)
}
