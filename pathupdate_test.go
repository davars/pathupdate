package pathupdate

import (
	"log"
	"reflect"
	"testing"
)

type pathUpdateTest struct {
	in    interface{}
	path  string
	value interface{}
	out   interface{}
	err   error
}

func TestPathUpdate(t *testing.T) {
	tests := []pathUpdateTest{
		{
			in:   map[string]interface{}{},
			path: "/",
			value: map[string]interface{}{
				"foo": 42,
			},
			out: map[string]interface{}{
				"foo": 42,
			},
		},
		{
			in:    map[string]interface{}{},
			path:  "/",
			value: 42,
			out:   42,
		},
		{
			in:    42,
			path:  "/foo",
			value: 42,
			out: map[string]interface{}{
				"foo": 42,
			},
		},
		{
			in: map[string]interface{}{
				"foo": 56,
				"bar": 78,
			},
			path:  "/foo",
			value: 42,
			out: map[string]interface{}{
				"foo": 42,
				"bar": 78,
			},
		},
		{
			in: map[string]interface{}{
				"foo": 56,
				"bar": 78,
			},
			path: "/foo/baz",
			value: map[string]interface{}{
				"bat": "boz",
			},
			out: map[string]interface{}{
				"foo": map[string]interface{}{
					"baz": map[string]interface{}{
						"bat": "boz",
					},
				},
				"bar": 78,
			},
		},
		{
			in: map[string]interface{}{
				"second": map[string]interface{}{
					"args":    []string{"arg1"},
					"command": "unblock",
				},
				"third": map[string]interface{}{
					"args":    []string{"arg2"},
					"command": "block",
				},
			},
			path:  "/second/command",
			value: "block",
			out: map[string]interface{}{
				"second": map[string]interface{}{
					"args":    []string{"arg1"},
					"command": "block",
				},
				"third": map[string]interface{}{
					"args":    []string{"arg2"},
					"command": "block",
				},
			},
		},
		{
			in: map[string]interface{}{
				"second": map[string]interface{}{
					"args":    []string{"arg1"},
					"command": "unblock",
				},
				"third": map[string]interface{}{
					"args":    []string{"arg2"},
					"command": "block",
				},
			},
			path:  "/second",
			value: nil,
			out: map[string]interface{}{
				"third": map[string]interface{}{
					"args":    []string{"arg2"},
					"command": "block",
				},
			},
		},
	}
	for _, test := range tests {
		m := &test.in
		PathUpdate(test.path, m, test.value)
		log.Printf("got %v, want %v", *m, test.out)
		if !reflect.DeepEqual(*m, test.out) {
			t.Errorf("got %v, want %v", *m, test.out)
		}
	}

}
