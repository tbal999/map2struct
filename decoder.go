package mapstringinterfacedecoder

import (
	"fmt"
	"strconv"
)

// Decoder - object used to ingest(decode) map[string]interface{}
type Decoder struct {
	Name  string
	Root  *Decoder
	sub   map[string][]*Decoder
	field map[string]string
	array map[string][]string
}

// Decode - decode a map[string]interface{} ready to be investigated into more
func (e *Decoder) Decode(m map[string]interface{}) {
	e.field = make(map[string]string)
	e.array = make(map[string][]string)
	e.sub = make(map[string][]*Decoder)
	if e.Root == nil {
		e.Name = "root"
	}

	for key, v := range m {
		switch vv := v.(type) {
		case string:
			value := vv
			e.field[key] = value
		case float64:
			value := strconv.FormatFloat(vv, 'f', -1, 64)
			e.field[key] = value
		case bool:
			value := strconv.FormatBool(vv)
			e.field[key] = value
		case []interface{}:
			contents := e.decodeInterfaceSlice(vv, key)
			e.array[key] = contents
		case map[string]interface{}:
			sub := &Decoder{}
			sub.Name = key
			sub.Root = e
			sub.Decode(vv)
			e.sub[key] = append(e.sub[key], sub)
		case nil:
			value := "NULL"
			e.field[key] = value
		default:
			e.field[key] = vv.(string) + "_type_not_supported"
		}
	}
}

// Get - grab an array of struct subentities
func (e *Decoder) Get(sub string) []*Decoder {
	if _, ok := e.sub[sub]; ok {
		return e.sub[sub]
	} else {
		fmt.Println(sub, "- is not a subentity")
	}

	return []*Decoder{e}
}

// Field - return the field [by name] of a struct
func (e *Decoder) Field(field string) string {
	if _, ok := e.field[field]; ok {
		return e.field[field]
	} else {
		fmt.Println(field, "- is not a field of this entity")
	}

	return ""
}

// Fields - return the field names of a struct as an array
func (e *Decoder) Fields() []string {
	fields := []string{}

	if len(e.field) != 0 {
		for key := range e.field {
			fields = append(fields, key)
		}
	}

	return fields
}

// Arrays - return the array names of a struct
func (e *Decoder) Arrays() []string {
	arrays := []string{}

	if len(e.array) != 0 {
		for key := range e.array {
			if len(e.array[key]) != 0 {
				arrays = append(arrays, key)
			}
		}
	}

	return arrays
}

// Array - return the array [by name] of a struct
func (e *Decoder) Array(array string) []string {
	if _, ok := e.array[array]; ok {
		return e.array[array]
	} else {
		fmt.Println(array, "- is not an array of this entity")
	}

	return []string{}
}

// Print - print out a struct
func (e *Decoder) Print() {
	var (
		rootName string
		entities []string
		fields   []string
		arrays   []string
	)
	if e.Root != nil {
		rootName = e.Root.Name
	}
	if len(e.field) != 0 {
		for key := range e.field {
			fields = append(fields, key)
		}
	}
	if len(e.array) != 0 {
		for key := range e.array {
			if len(e.array[key]) != 0 {
				arrays = append(arrays, key)
			}
		}
	}
	if len(e.sub) != 0 {
		for key := range e.sub {
			entities = append(entities, key)
		}
	}
	fmt.Printf("name: %s\nroot: %s\nfield: %s\narray: %v\nsub: %v\n",
		e.Name, rootName, fields, arrays, entities)
}

func (e *Decoder) decodeInterfaceSlice(m []interface{}, key string) []string {
	values := make([]string, 0, len(m))
	for _, v := range m {
		switch vv := v.(type) {
		case map[string]interface{}:
			sub := &Decoder{}
			sub.Name = key
			sub.Root = e
			sub.Decode(vv)
			e.sub[key] = append(e.sub[key], sub)
		case string:
			value := vv
			values = append(values, value)
		case float64:
			value := strconv.FormatFloat(vv, 'f', -1, 64)
			values = append(values, value)
		case []interface{}:
			contents := e.decodeInterfaceSlice(vv, key)
			e.array[key] = contents
		case bool:
			value := strconv.FormatBool(vv)
			values = append(values, value)
		case nil:
			value := "NULL"
			values = append(values, value)
		}
	}
	return values
}
