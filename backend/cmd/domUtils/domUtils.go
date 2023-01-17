package domUtils

import (
	"fmt"
	"strings"
)

func ToJson(vars ...string) []byte {
	res := strings.Builder{}
	res.WriteString("{")
	vLen := len(vars)
	for idx, v := range vars {
		data := strings.Split(v, ":=")
		if len(data) != 2 {
			continue
		}
		res.WriteString(fmt.Sprintf("\"%s\"=\"%s\"", data[0], data[1]))
		if idx < vLen-1 {
			res.WriteString(",")
		}
	}
	res.WriteString("}")

	return []byte(res.String())
}
