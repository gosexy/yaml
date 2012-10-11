# gosexy/yaml

`gosexy/yaml` provides sugar methods for loading, reading and writing to and from [YAML](http://www.yaml.org/) formatted files.

This package is a wrapper of [goyaml](http://launchpad.net/goyaml).

## Installation

```
go get github.com/gosexy/yaml
```

## Usage

```go
package main

import (
	"github.com/gosexy/yaml"
	"github.com/gosexy/sugar"
)

func main() {
	settings := yaml.New()
	settings.Set("success", true)
	settings.Set("nested/tree", 1)
	settings.Set("another/nested/tree", sugar.List { 1, 2, 3 })
	settings.Write("test.yaml")
}
```

The above code would generate a `test.yaml` file like this:

```yaml
another:
  nested:
    tree:
    - 1
    - 2
    - 3
success: true
nested:
  tree: 1
```

## Documentation

You can read ``gosexy/yaml`` documentation from a terminal

```
$ go doc github.com/gosexy/yaml
```

Or you can [browse it](http://go.pkgdoc.org/github.com/gosexy/yaml) online.
