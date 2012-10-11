# gosexy/yaml

This package is a wrapper of [goyaml](http://launchpad.net/goyaml) that provides methods for loading, reading and writing to and from [YAML](http://www.yaml.org/) formatted files.

## Installation

    $ go get github.com/gosexy/yaml

## Usage

```go
package main

import "github.com/gosexy/yaml"

func main() {
	settings := yaml.New()
	settings.Set("success", true)
	settings.Write("test.yaml")
}
```

## Documentation

You can read ``gosexy/yaml`` documentation from a terminal

```
$ go doc github.com/gosexy/yaml
```

Or you can [browse it](http://go.pkgdoc.org/github.com/gosexy/yaml) online.
