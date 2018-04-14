# pathupdate
Manipulates nested maps using "paths"

Given a path of keys like `/foo/bar` and a value like `42`, sets the value `42` under the key `bar` in a map under the key `foo` and returns that map.  So `("/foo/bar", 42) -> {"foo": {"bar": 42}}`.  You can delete values by setting them to `nil`.  See [the tests](https://github.com/davars/pathupdate/blob/master/pathupdate_test.go) for details.
