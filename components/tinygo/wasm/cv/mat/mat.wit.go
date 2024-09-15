// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package mat represents the imported interface "wasm:cv/mat".
//
// mat resource is a matrix of bytes, wrapper around the cv::Mat type.
package mat

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

// Mattype represents the enum "wasm:cv/mat#mattype".
//
//	enum mattype {
//		cv8u,
//		cv8s,
//		cv16u,
//		cv16s,
//		cv32s,
//		cv32f,
//		cv64f
//	}
type Mattype uint8

const (
	MattypeCv8u Mattype = iota
	MattypeCv8s
	MattypeCv16u
	MattypeCv16s
	MattypeCv32s
	MattypeCv32f
	MattypeCv64f
)

var stringsMattype = [7]string{
	"cv8u",
	"cv8s",
	"cv16u",
	"cv16s",
	"cv32s",
	"cv32f",
	"cv64f",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e Mattype) String() string {
	return stringsMattype[e]
}

// Mat represents the imported resource "wasm:cv/mat#mat".
//
//	resource mat
type Mat cm.Resource

// ResourceDrop represents the imported resource-drop for resource "mat".
//
// Drops a resource handle.
//
//go:nosplit
func (self Mat) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_MatResourceDrop((uint32)(self0))
	return
}

//go:wasmimport wasm:cv/mat [resource-drop]mat
//go:noescape
func wasmimport_MatResourceDrop(self0 uint32)

// NewMat represents the imported constructor for resource "mat".
//
//	constructor(cols: u32, rows: u32, %type: mattype)
//
//go:nosplit
func NewMat(cols uint32, rows uint32, type_ Mattype) (result Mat) {
	cols0 := (uint32)(cols)
	rows0 := (uint32)(rows)
	type0 := (uint32)(type_)
	result0 := wasmimport_NewMat((uint32)(cols0), (uint32)(rows0), (uint32)(type0))
	result = cm.Reinterpret[Mat]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/mat [constructor]mat
//go:noescape
func wasmimport_NewMat(cols0 uint32, rows0 uint32, type0 uint32) (result0 uint32)

// Close represents the imported method "close".
//
// close the Mat
//
//	close: func()
//
//go:nosplit
func (self Mat) Close() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_MatClose((uint32)(self0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.close
//go:noescape
func wasmimport_MatClose(self0 uint32)

// Cols represents the imported method "cols".
//
// Cols returns the number of columns for this Mat.
//
//	cols: func() -> u32
//
//go:nosplit
func (self Mat) Cols() (result uint32) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_MatCols((uint32)(self0))
	result = (uint32)((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.cols
//go:noescape
func wasmimport_MatCols(self0 uint32) (result0 uint32)

// Empty represents the imported method "empty".
//
// empty returns true if the Mat is empty.
//
//	empty: func() -> bool
//
//go:nosplit
func (self Mat) Empty() (result bool) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_MatEmpty((uint32)(self0))
	result = cm.U32ToBool((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.empty
//go:noescape
func wasmimport_MatEmpty(self0 uint32) (result0 uint32)

// Reshape represents the imported method "reshape".
//
// Reshape changes the shape and/or the number of channels of a 2D matrix without
// copying the data.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d3/d63/classcv_1_1Mat.html#a4eb96e3251417fa88b78e2abd6cfd7d8
//
//	reshape: func(channels: u32, rows: u32) -> mat
//
//go:nosplit
func (self Mat) Reshape(channels uint32, rows uint32) (result Mat) {
	self0 := cm.Reinterpret[uint32](self)
	channels0 := (uint32)(channels)
	rows0 := (uint32)(rows)
	result0 := wasmimport_MatReshape((uint32)(self0), (uint32)(channels0), (uint32)(rows0))
	result = cm.Reinterpret[Mat]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.reshape
//go:noescape
func wasmimport_MatReshape(self0 uint32, channels0 uint32, rows0 uint32) (result0 uint32)

// Rows represents the imported method "rows".
//
// Rows returns the number of rows for this Mat.
//
//	rows: func() -> u32
//
//go:nosplit
func (self Mat) Rows() (result uint32) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_MatRows((uint32)(self0))
	result = (uint32)((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.rows
//go:noescape
func wasmimport_MatRows(self0 uint32) (result0 uint32)

// Size represents the imported method "size".
//
// Size returns an array with one element for each dimension containing the size of
// that dimension for the Mat.
//
//	size: func() -> list<u32>
//
//go:nosplit
func (self Mat) Size() (result cm.List[uint32]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_MatSize((uint32)(self0), &result)
	return
}

//go:wasmimport wasm:cv/mat [method]mat.size
//go:noescape
func wasmimport_MatSize(self0 uint32, result *cm.List[uint32])

// Type represents the imported method "type".
//
// type returns the type of the Mat.
//
//	%type: func() -> mattype
//
//go:nosplit
func (self Mat) Type() (result Mattype) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_MatType((uint32)(self0))
	result = (Mattype)((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/mat [method]mat.type
//go:noescape
func wasmimport_MatType(self0 uint32) (result0 uint32)
