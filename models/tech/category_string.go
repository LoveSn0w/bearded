// generated by stringer -type=Category; DO NOT EDIT

package tech

import "fmt"

const _Category_name = "CMSJavascriptFrameworks"

var _Category_index = [...]uint8{0, 3, 23}

func (i Category) String() string {
	if i < 0 || i+1 >= Category(len(_Category_index)) {
		return fmt.Sprintf("Category(%d)", i)
	}
	return _Category_name[_Category_index[i]:_Category_index[i+1]]
}