// Code generated by "stringer -type=IndexType -output=index_type_string_gen.go"; DO NOT EDIT.

package model

import "fmt"

const _IndexType_name = "IndexTypeNoneIndexTypeBtreeIndexTypeHash"

var _IndexType_index = [...]uint8{0, 13, 27, 40}

func (i IndexType) String() string {
	if i < 0 || i >= IndexType(len(_IndexType_index)-1) {
		return fmt.Sprintf("IndexType(%d)", i)
	}
	return _IndexType_name[_IndexType_index[i]:_IndexType_index[i+1]]
}
