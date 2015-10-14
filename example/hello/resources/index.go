/**
* Author: CZ cz.theng@gmail.com
 */

package resources

import (
	"github.com/cz-it/ridge"
)

func init() {
	Index = &IndexRes{}
}

var Index *IndexRes

type IndexRes struct {
	ridge.Resource
}
