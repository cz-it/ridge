/**
* Author: CZ cz.theng@gmail.com
 */

package main

import (
	"github.com/cz-it/ridge"
	"github.com/cz-it/ridge/example/hello/resources"
)

func init() {
	ridge.CreateResource("index", resources.Index)
}
