package alpaca

import (
	"os"
	"regexp"
)

func MakeDir(name string) {
	HandleError(os.Mkdir(name, 0755))
	MoveDir(name)
}

func MoveDir(name string) {
	HandleError(os.Chdir(name))
}

func ArrayInterfaceToString(inter interface{}) []string {
	old := inter.([]interface{})
	new := make([]string, len(old))

	for i, v := range old {
		new[i] = v.(string)
	}

	return new
}

func MapKeysToStringArray(inter interface{}, exclude []string) []string {
	old := inter.(map[string]interface{})
	new := make([]string, 0, len(old))

	for v, _ := range old {
		flag := true

		for _, e := range exclude {
			if e == v {
				flag = false
			}
		}

		if flag {
			new = append(new, v)
		}
	}

	return new
}

func ActiveClassInfo(name string, class interface{}) map[string]interface{} {
	data := make(map[string]interface{})

	data["name"] = name
	data["methods"] = MapKeysToStringArray(class, []string{"args"})
	data["args"] = class.(map[string]interface{})["args"]

	return data
}

func ArgsFunctionMaker(before, after string) interface{} {
	return func(args interface{}, options ...bool) string {
		str := ""

		if args != nil && len(args.([]interface{})) > 0 {
			for _, v := range ArrayInterfaceToString(args) {
				str += before + v + after
			}

			if len(options) > 0 && options[0] {
				str = str[0 : len(str)-len(after)]
			}

			if len(options) > 1 && options[1] {
				str = after + str
			}
		}

		return str
	}
}

func PathFunctionMaker(before, after string) interface{} {
	return func(path string, args interface{}) string {
		if args != nil {
			for _, v := range ArrayInterfaceToString(args) {
				reg := regexp.MustCompile(":(" + v + ")")
				path = reg.ReplaceAllString(path, before+"$1"+after)
			}
		}

		return path
	}
}
