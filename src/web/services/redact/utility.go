package redact

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var (
	// WILDCARD_KEY is used to represent l[*] type of keys in json path
	WILDCARD_KEY = "WILDCARD_KEY"
	// ARRAY_KEY is used to represent l[0] type of keys in json path
	ARRAY_KEY = "ARRAY_KEY"
	// OBJECT_KEY is used to represent a.b, a.c type of keys in json path
	OBJECT_KEY = "OBJECT_KEY"
	// LIST_KEY is used to represent [0] type of keys in json path
	LIST_KEY = "LIST_KEY"
	// WILDCARD_LIST_KEY is used to represent [*] type of keys in json path
	WILDCARD_LIST_KEY = "WILDCARD_LIST_KEY"

	redacted bool
)

func (r *redact) redactJSON(m interface{}, paths []string, regex []string, replaceall bool) bool {
	// base case
	if m == nil || len(paths) == 0 {
		return redacted
	}

	key := paths[0]
	keyType := determineKeyType(key)
	if keyType == OBJECT_KEY {
		if reflect.TypeOf(m).Kind() != reflect.Map {
			fmt.Println("Interface Did No Match", OBJECT_KEY)
			return redacted
		}

		// key not present into interface
		if m.(map[string]interface{})[key] == nil {
			fmt.Println("Key Not Present", OBJECT_KEY, key)
			return redacted
		}

		// reached the leaf node apply the redact logic at leaf node
		if len(paths) == 1 {
			// check if the key being replaced with * is string only, redacing complex types such as arrays, objects in not allowed
			if reflect.TypeOf(m.(map[string]interface{})[key]).Kind() == reflect.String {
				redacted = true
				m.(map[string]interface{})[key] = strings.Repeat("*", len(fmt.Sprint(m.(map[string]interface{})[key])))
			}
		}

		// continue probing interface with key
		if len(paths) > 1 {
			m = m.(map[string]interface{})[key]
			r.redactJSON(m, paths[1:], regex, replaceall)
		}
	}

	if keyType == ARRAY_KEY {
		keyItem, i := parseArrayIndex(key)
		// check if given key is present in the interface
		if m.(map[string]interface{})[keyItem] == nil {
			fmt.Println("Key Not Present", OBJECT_KEY, key)
			return redacted
		}

		// replace p[index] with actual index item
		m = m.(map[string]interface{})[keyItem]
		if reflect.TypeOf(m).Kind() != reflect.Slice {
			fmt.Println("No Match", ARRAY_KEY)
			return redacted
		}

		idx, _ := strconv.Atoi(i)

		// out of index case
		if idx >= len(m.([]interface{})) {
			fmt.Println("Array Out on Index", OBJECT_KEY, key)
			return redacted
		}

		// reached the leaf node apply the redact logic at leaf node
		if len(paths) == 1 {
			// check if the key being replaced with * is string only, redacing complex types such as arrays, objects in not allowed
			if reflect.TypeOf(m.([]interface{})[idx]).Kind() == reflect.String {
				redacted = true
				m.([]interface{})[idx] = strings.Repeat("*", len(fmt.Sprint(m.([]interface{})[idx])))
			}
			return redacted
		}

		// continue probing interface with key
		if len(paths) > 1 {
			m = m.([]interface{})[idx]
			r.redactJSON(m, paths[1:], regex, replaceall)
		}
	}

	if keyType == WILDCARD_KEY {
		keyItem, _ := parseArrayIndex(key)
		// check if given key is present in the interface
		if m.(map[string]interface{})[keyItem] == nil {
			fmt.Println("Key Not Present", OBJECT_KEY, key)
			return redacted
		}

		// replace p[index] with actual index item
		m = m.(map[string]interface{})[keyItem]
		if reflect.TypeOf(m).Kind() != reflect.Slice {
			fmt.Println("No Match", WILDCARD_KEY)
			return redacted
		}

		// reached the leaf node apply the redact logic at leaf node
		if len(paths) == 1 {
			for k, v := range m.([]interface{}) {
				// check if the key being replaced with * is string only, redacing complex types such as arrays, objects in not allowed
				if reflect.TypeOf(m.([]interface{})[k]).Kind() == reflect.String {
					redacted = true
					m.([]interface{})[k] = strings.Repeat("*", len(fmt.Sprint(v)))
				}
			}
		}

		// need to redact interface recursively again
		if len(paths) > 1 {
			for i, v := range m.([]interface{}) {
				// make sure key is present inside the interface
				if reflect.TypeOf(v).Kind() == reflect.Map && v.(map[string]interface{})[paths[1]] != nil {
					r.redactJSON(v, paths[1:], regex, replaceall)
				} else if reflect.TypeOf(v).Kind() == reflect.String && v == paths[1] {
					redacted = true
					m.([]interface{})[i] = strings.Repeat("*", len(fmt.Sprint(v)))
				}
			}
		}
	}

	if keyType == LIST_KEY {
		_, i := parseArrayIndex(key)
		idx, _ := strconv.Atoi(i)

		// out of index case
		if idx >= len(m.([]interface{})) {
			fmt.Println("Array Out on Index", LIST_KEY, key)
			return redacted
		}

		// replace p[index] with actual index item
		m = m.([]interface{})[idx]
		if reflect.TypeOf(m).Kind() != reflect.Slice {
			fmt.Println("No Match", LIST_KEY)
			return redacted
		}

		// reached the leaf node apply the redact logic at leaf node
		if len(paths) == 1 {
			// check if the key being replaced with * is string only, redacing complex types such as arrays, objects in not allowed
			if reflect.TypeOf(m.([]interface{})[idx]).Kind() == reflect.String {
				redacted = true
				m.([]interface{})[idx] = strings.Repeat("*", len(fmt.Sprint(m.([]interface{})[idx])))
			}
		}

		// continue probing interface with key
		if len(paths) > 1 {
			m = m.([]interface{})[idx]
			r.redactJSON(m, paths[1:], regex, replaceall)
		}
	}

	if keyType == WILDCARD_LIST_KEY {
		// replace p[index] with actual index item
		m = m.([]interface{})

		// redact each element in the list
		for _, item := range m.([]interface{}) {
			if reflect.TypeOf(item).Kind() != reflect.Slice {
				fmt.Println("No Match", WILDCARD_LIST_KEY, item, key)
				return redacted
			}

			m = item.([]interface{})

			// reached the leaf node apply the redact logic at leaf node
			if len(paths) == 1 {
				for k, v := range m.([]interface{}) {
					// check if the key being replaced with * is string only, redacing complex types such as arrays, objects in not allowed
					if reflect.TypeOf(m.([]interface{})[k]).Kind() == reflect.String {
						redacted = true
						m.([]interface{})[k] = strings.Repeat("*", len(fmt.Sprint(v)))
					}
				}
			}

			// need to redact interface recursively again
			if len(paths) > 1 {
				for i, v := range m.([]interface{}) {
					// make sure key is present inside the interface or skip it
					if reflect.TypeOf(v).Kind() == reflect.Map && v.(map[string]interface{})[paths[1]] != nil {
						r.redactJSON(v, paths[1:], regex, replaceall)
					} else if reflect.TypeOf(v).Kind() == reflect.String && v == paths[1] {
						redacted = true
						m.([]interface{})[i] = strings.Repeat("*", len(fmt.Sprint(v)))
					}
				}
			}
		}
	}

	return redacted
}

func determineKeyType(key string) string {
	if strings.Contains(key, "[") {
		key, res := parseArrayIndex(key)
		if len(key) > 0 && res == "*" {
			return WILDCARD_KEY
		} else if len(key) > 0 && res != "*" {
			return ARRAY_KEY
		}

		if len(key) == 0 && res == "*" {
			return WILDCARD_LIST_KEY
		} else if len(key) == 0 && res != "*" {
			return LIST_KEY
		}
	}

	return OBJECT_KEY
}

func parseArrayIndex(key string) (string, string) {
	leafIdx := strings.Replace(key, "[", ",", -1)
	leafIdx = strings.Replace(leafIdx, "]", ",", -1)
	res := strings.Split(leafIdx, ",")

	return res[0], res[1]
}
