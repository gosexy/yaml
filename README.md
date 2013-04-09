# gosexy/yaml

This package is a wrapper of [goyaml][2].

`gosexy/yaml` provides friendly methods for loading, reading and writing to and
from [YAML][3] formatted files.

Useful if you just want to read/write setting files in your Go programs.

## Installation

```
go get menteslibres.net/gosexy/yaml
```

## Usage

After installing, use the following import path.

```go
import "menteslibres.net/gosexy/yaml"
```

Here's an example that creates a YAML file and writes some values on it:

```go
package main

import (
	"menteslibres.net/gosexy/yaml"
)

func main() {
	settings := yaml.New()

	settings.Set("success", true)
	settings.Set("nested", "tree", 1)
	settings.Set("another", "nested", "tree", []int{1, 2, 3})

	settings.Write("test.yaml")
}
```

The above code generates a `test.yaml` file like this:

```yaml
another:
  nested:
    tree:
    - 1
    - 2
    - 3
nested:
  tree: 1
success: true
```

A detail of another example:

```go
settings, err := yaml.Open("examples/input/settings.yaml")

if err != nil {
	log.Printf("Could not open YAML file: %s", err.Error())
}

s := to.String(settings.Get("test_string"))

fmt.Printf("%s", s)
```

If the referred `settings.yaml` contains a line like this:

```yaml
test_string: "Hello World!"
```

The the output would be:

```
Hello World!
```

Note that you can use nested paths on `Set()` and `Get()`:

```go
/*
path:
	to:
		nested:
			value: 1
*/
settings.Set("path", "to", "nested", "value", 1)

// yaml.*Yaml.Get() returns interface{}.
i := settings.Get("path", "to", "nested", "value")

fmt.Printf("%d\n", i)
// Prints: 1
```

You can also use [gosexy/to][4] to convert from `interface{}` into a compatible
type:

```go
// to.Int64() returns int64.
i := to.Int64(settings.Get("path", "to", "nested", "value"))

fmt.Printf("%d\n", i)
// Prints: 1
```

## Documentation

See the [online docs][1].

[1]: http://godoc.org/menteslibres.net/gosexy/yaml
[2]: http://launchpad.net/goyaml
[3]: http://www.yaml.org
[4]: https://menteslibres.net/gosexy/to
