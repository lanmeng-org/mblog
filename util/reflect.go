package util

import "reflect"

type StructDesc struct {
	FieldDesc map[string]*StructFieldDesc
	FieldArr  []string
}

type StructFieldDesc struct {
	Field string
	Type  reflect.Type
	Value interface{}
}

var structDescCache map[string]*StructDesc

func ReflectStructDesc(obj interface{}, cacheName string) *StructDesc {
	if cacheName != "" {
		if cache, ok := structDescCache[cacheName]; ok {
			return cache
		}
	}

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	structDesc := new(StructDesc)
	structDesc.FieldDesc = make(map[string]*StructFieldDesc, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		fieldDesc := new (StructFieldDesc)
		tf := t.Field(i)
		fieldDesc.Type = tf.Type
		fieldDesc.Field = tf.Name
		fieldDesc.Value = v.Field(i).Interface()

		structDesc.FieldDesc[tf.Name] = fieldDesc;
		structDesc.FieldArr = append(structDesc.FieldArr, tf.Name)
	}

	return structDesc
}
