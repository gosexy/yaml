/*
  Copyright (c) 2012 José Carlos Nieto, http://xiam.menteslibres.org/

  Permission is hereby granted, free of charge, to any person obtaining
  a copy of this software and associated documentation files (the
  "Software"), to deal in the Software without restriction, including
  without limitation the rights to use, copy, modify, merge, publish,
  distribute, sublicense, and/or sell copies of the Software, and to
  permit persons to whom the Software is furnished to do so, subject to
  the following conditions:

  The above copyright notice and this permission notice shall be
  included in all copies or substantial portions of the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
  LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
  OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
  WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package yaml

import (
	"fmt"
	"github.com/gosexy/sugar"
	"github.com/gosexy/to"
	"launchpad.net/goyaml"
	"os"
	"strings"
)

type Yaml struct {
	file   string
	values *sugar.Tuple
}

/* Creates and returns a YAML struct. */
func New() *Yaml {
	self := &Yaml{}
	self.values = &sugar.Tuple{}
	return self
}

/* Creates and returns a YAML struct, from a file. */
func Open(file string) (*Yaml, error) {
	var err error

	self := New()

	_, err = os.Stat(file)

	if err != nil {
		return nil, err
	}

	self.file = file

	err = self.Read(self.file)

	if err != nil {
		return nil, err
	}

	return self, nil
}

/* Sets a YAML setting */
func (self *Yaml) Set(path string, value interface{}) error {
	return self.values.Set(path, value)
}

/* Returns a YAML setting */
func (self *Yaml) Get(path string) interface{} {
	return self.values.Get(path)
}

func mapValues(data interface{}, parent *sugar.Tuple) {

	var name string

	for key, value := range data.(map[interface{}]interface{}) {

		name = strings.ToLower(to.String(key))

		switch value.(type) {
		case map[interface{}]interface{}:
			values := &sugar.Tuple{}
			mapValues(value, values)
			(*parent)[name] = *values
		default:
			(*parent)[name] = value
		}

	}

}

/* Writes changes to the currently opened YAML file. */
func (self *Yaml) Save() error {
	if self.file != "" {
		return self.Write(self.file)
	} else {
		return fmt.Errorf("No file specified.")
	}
	return nil
}

/* Writes the YAML struct into a file */
func (self *Yaml) Write(filename string) error {

	out, err := goyaml.Marshal(self.values)

	if err != nil {
		return err
	}

	fp, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer fp.Close()

	fp.Write(out)

	return nil
}

/* Loads a YAML file. */
func (self *Yaml) Read(filename string) error {
	var err error
	var data interface{}

	fileinfo, err := os.Stat(filename)

	if err != nil {
		return err
	}

	filesize := fileinfo.Size()

	fp, err := os.Open(filename)

	if err != nil {
		return err
	}

	defer fp.Close()

	buf := make([]byte, filesize)
	fp.Read(buf)

	err = goyaml.Unmarshal(buf, &data)

	if err == nil {
		mapValues(data, self.values)
	} else {
		return err
	}

	return nil
}
