// jsResult is a helper package which defines a JS object with result of some operation
// in a unified format: { result: <something>, error: "string" }
package jsresult

import "syscall/js"

func Result(result any, err error) js.Value {
	if err != nil {
		return js.ValueOf(map[string]any{
			"error": err.Error(),
		})
	}
	if result != nil {
		return js.ValueOf(map[string]any{
			"result": result,
		})
	}
	return js.ValueOf(map[string]any{})
}
