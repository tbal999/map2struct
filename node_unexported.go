package map2struct

import (
	"strconv"
)

// ingestInterfaceSlice - for when we're working with []interface{}
func (e *Node) ingestInterfaceSlice(m []interface{}, key string) []string {
	values := make([]string, 0, len(m))
	for _, v := range m {
		switch vv := v.(type) {
		case map[string]interface{}:
			sub := &Node{}
			sub.Name = key
			sub.Root = e
			sub.Ingest(vv)
			e.sub[key] = append(e.sub[key], sub)
		case string:
			value := vv
			values = append(values, value)
		case float64:
			value := strconv.FormatFloat(vv, 'f', -1, 64)
			values = append(values, value)
		case []interface{}:
			contents := e.ingestInterfaceSlice(vv, key)
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
