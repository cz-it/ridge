/**
* Author: CZ cz.theng@gmail.com
 */

package ridge

import (
	"fmt"
)

type Resourcer interface {
	Index()
	/*
		Create()
		Query()
		Update()
		Delate()
	*/
}

type Resource struct {
	Name string
}

func (res *Resource) Index() {
	fmt.Println("Index...")
}

func Run() {
	fmt.Println("Ridge.Start...")
}
