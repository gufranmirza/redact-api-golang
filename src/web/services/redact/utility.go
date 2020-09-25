package redact

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func (r *redact) redactJSON(m map[string]interface{}, keys []string, regex []string, replaceall bool) error {
	// handle possible panic due to index out of range and invalid keys etc.
	defer func() error {
		if r := recover(); r != nil {
			return fmt.Errorf("Failed to redact JSON: %s, recovered from panic", m)
		}
		return nil
	}()

	for k, v := range m {
		leafIdx := keys[0]
		// special handing if indexes are given in path
		if strings.Contains(keys[0], "[") {
			leafIdx = strings.Replace(leafIdx, "[", ",", -1)
			leafIdx = strings.Replace(leafIdx, "]", ",", -1)
			res := strings.Split(leafIdx, ",")
			idx, err := strconv.Atoi(res[1])
			// invalid index NaN
			if err != nil {
				return err
			}

			// if current item is array then tranform array
			if reflect.TypeOf(v).Kind() == reflect.Slice {
				v = v.([]interface{})[idx].(map[string]interface{})
				// executed only once
				for k1 := range v.(map[string]interface{}) {
					leafIdx = k1
					k = k1
				}
			}
		}

		if k == leafIdx && len(keys) > 1 {
			r.redactJSON(v.(map[string]interface{}), keys[1:], regex, replaceall)
		} else if len(keys) == 1 && leafIdx == k {
			// handle replace all condition
			if replaceall {
				m[leafIdx] = strings.Repeat("*", len(fmt.Sprint(m[leafIdx])))
				return nil
			}

			// handle regular expression
			for _, reg := range regex {
				if reflect.TypeOf(m[leafIdx]).Kind() != reflect.String {
					buf, _ := json.Marshal(m[leafIdx])
					m[leafIdx] = string(buf)
				}
				re := regexp.MustCompile(reg)
				s := re.ReplaceAllStringFunc(m[leafIdx].(string), func(s string) string { return strings.Repeat("*", len(s)) })
				m[leafIdx] = s
			}
			return nil
		}
	}

	return nil
}
