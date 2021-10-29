package decoder

import (
	"fmt"
	"strconv"
)

type Decoder struct {
	Name  string
	Root  *entity
	Sub   map[string]*Decoder
	Field map[string]string
	Array map[string][]string
}

func (e *Decoder) Print() {
	var (
		rootName string
		entities []string
	)
	if e.Root != nil {
		rootName = e.Root.Name
	}
	if len(e.Sub) != 0 {
		for key := range e.Sub {
			entities = append(entities, key)
		}
	}
	fmt.Printf("Name: %s\nRoot: %s\nField: %s\nArray: %v\nSub: %v\n",
		e.Name, rootName, e.Field, e.Array, entities)
}

func (e *Decoder) Ingest(m map[string]interface{}) {
	e.Field = make(map[string]string)
	e.Array = make(map[string][]string)
	e.Sub = make(map[string]*Decoder)
	if e.Root == nil {
		e.Name = "root"
	}

	for key, v := range m {
		switch vv := v.(type) {
		case string:
			value := vv
			e.Field[key] = value
		case float64:
			value := strconv.FormatFloat(vv, 'f', -1, 64)
			e.Field[key] = value
		case bool:
			value := strconv.FormatBool(vv)
			e.Field[key] = value
		case nil:
			value := "NULL"
			e.Field[key] = value
		}
	}

	for key, v := range m {
		switch vv := v.(type) {
		case []interface{}:
			contents := e.decodeInterfaceSlice(vv, key)
			e.Array[key] = contents
		}
	}

	for key, v := range m {
		switch vv := v.(type) {
		case map[string]interface{}:
			sub := &Decoder{}
			sub.Name = key
			sub.Root = e
			sub.Ingest(vv)
			e.Sub[key] = sub
		}
	}
}

func (e *Decoder) decodeInterfaceSlice(m []interface{}, key string) []string {
	values := make([]string, 0, len(m))
	for _, v := range m {
		switch vv := v.(type) {
		case map[string]interface{}:
			sub := &Decoder{}
			sub.Name = key
			sub.Root = e
			sub.Ingest(vv)
			e.Sub[key] = sub
		case string:
			value := vv
			values = append(values, value)
		case float64:
			value := strconv.FormatFloat(vv, 'f', -1, 64)
			values = append(values, value)
		case []interface{}:
			contents := e.decodeInterfaceSlice(vv, key)
			e.Array[key] = contents
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
