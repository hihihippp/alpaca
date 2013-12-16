package alpaca

import (
	"github.com/robertkrimen/terst"
	"testing"
)

func TestArrayInterfaceToString(t *testing.T) {
	terst.Terst(t)

	old := make([]interface{}, 2)

	old[0] = "hello"
	old[1] = "world"

	terst.Is(ArrayInterfaceToString(old), []string{"hello", "world"})
}

func TestMapKeysToStringArray(t *testing.T) {
	terst.Terst(t)

	old := make(map[string]interface{})

	old["a"] = "aaa"
	old["b"] = "bbb"
	old["c"] = "ccc"

	terst.Is(MapKeysToStringArray(old, []string{"b"}), []string{"a", "c"})
}

func TestActiveClassInfo(t *testing.T) {
	terst.Terst(t)

	old := make(map[string]interface{})

	old["args"] = []string{"id", "url"}
	old["create"] = make(map[string]interface{})
	old["update"] = make(map[string]interface{})

	v := ActiveClassInfo("test", old)

	terst.Is(v["name"], "test")
	terst.Is(v["args"].([]string), []string{"id", "url"})
	terst.Is(v["methods"].([]string), []string{"create", "update"})
}

func TestArgsFunctionMaker(t *testing.T) {
	terst.Terst(t)

	f := ArgsFunctionMaker("$", ", ").(func(interface{}, ...bool) string)
	args, noargs := make([]interface{}, 2), make([]interface{}, 0)

	args[0] = "id"
	args[1] = "url"

	terst.Is(f(args), "$id, $url, ")
	terst.Is(f(args, false), "$id, $url, ")
	terst.Is(f(args, true), "$id, $url")
	terst.Is(f(args, true, false), "$id, $url")
	terst.Is(f(args, true, true), ", $id, $url")
	terst.Is(f(args, false, true), ", $id, $url, ")

	terst.Is(f(noargs), "")
	terst.Is(f(noargs, true), "")
	terst.Is(f(noargs, false, true), "")
}

func TestPathFunctionMaker(t *testing.T) {
	terst.Terst(t)

	f := PathFunctionMaker("\"+@", "+\"").(func(string, interface{}) string)
	args := make([]interface{}, 2)

	args[0] = "id"
	args[1] = "url"

	terst.Is(f("/user/:id/:not/:url/wow", args), "/user/\"+@id+\"/:not/\"+@url+\"/wow")
}
