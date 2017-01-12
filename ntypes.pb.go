// Code generated by protoc-gen-go.
// source: ntypes.proto
// DO NOT EDIT!

/*
Package ntypes is a generated protocol buffer package.

It is generated from these files:
	ntypes.proto

It has these top-level messages:
	String
	StringArray
	Int32
	Int32Array
	Int64
	Int64Array
	Uint32
	Uint32Array
	Uint64
	Uint64Array
	Float32
	Float32Array
	Float64
	Float64Array
	Bool
	BoolArray
*/
package ntypes

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// String represents a string that may be nil.
type String struct {
	Chars string `protobuf:"bytes,1,opt,name=chars" json:"chars,omitempty"`
	Valid bool   `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *String) Reset()                    { *m = String{} }
func (m *String) String() string            { return proto.CompactTextString(m) }
func (*String) ProtoMessage()               {}
func (*String) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// StringArray represents an array of strings that may be nil.
type StringArray struct {
	StringArray []string `protobuf:"bytes,1,rep,name=string_array,json=stringArray" json:"string_array,omitempty"`
	Valid       bool     `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *StringArray) Reset()                    { *m = StringArray{} }
func (m *StringArray) String() string            { return proto.CompactTextString(m) }
func (*StringArray) ProtoMessage()               {}
func (*StringArray) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Int32 represents a int32 that may be nil.
type Int32 struct {
	Int32 int32 `protobuf:"varint,1,opt,name=int32" json:"int32,omitempty"`
	Valid bool  `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *Int32) Reset()                    { *m = Int32{} }
func (m *Int32) String() string            { return proto.CompactTextString(m) }
func (*Int32) ProtoMessage()               {}
func (*Int32) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Int32Array represents an array of int32s that may be nil.
type Int32Array struct {
	Int32Array []int32 `protobuf:"varint,1,rep,packed,name=int32_array,json=int32Array" json:"int32_array,omitempty"`
	Valid      bool    `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (ia *Int32Array) Reset()                   { *ia = Int32Array{} }
func (ia *Int32Array) String() string           { return proto.CompactTextString(ia) }
func (*Int32Array) ProtoMessage()               {}
func (*Int32Array) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// Int64 represents a int64 that may be nil.
type Int64 struct {
	Int64 int64 `protobuf:"varint,1,opt,name=int64" json:"int64,omitempty"`
	Valid bool  `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *Int64) Reset()                    { *m = Int64{} }
func (m *Int64) String() string            { return proto.CompactTextString(m) }
func (*Int64) ProtoMessage()               {}
func (*Int64) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// Int64Array represents an array of int64s that may be nil.
type Int64Array struct {
	Int64Array []int64 `protobuf:"varint,1,rep,packed,name=int64_array,json=int64Array" json:"int64_array,omitempty"`
	Valid      bool    `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (ia *Int64Array) Reset()                   { *ia = Int64Array{} }
func (ia *Int64Array) String() string           { return proto.CompactTextString(ia) }
func (*Int64Array) ProtoMessage()               {}
func (*Int64Array) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// Uint32 represents a Uint32 that may be nil.
type Uint32 struct {
	Uint32 uint32 `protobuf:"varint,1,opt,name=uint32" json:"uint32,omitempty"`
	Valid  bool   `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *Uint32) Reset()                    { *m = Uint32{} }
func (m *Uint32) String() string            { return proto.CompactTextString(m) }
func (*Uint32) ProtoMessage()               {}
func (*Uint32) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// Uint32Array represents an array of uint32s that may be nil.
type Uint32Array struct {
	Uint32Array []uint32 `protobuf:"varint,1,rep,packed,name=uint32_array,json=uint32Array" json:"uint32_array,omitempty"`
	Valid       bool     `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (ua *Uint32Array) Reset()                   { *ua = Uint32Array{} }
func (ua *Uint32Array) String() string           { return proto.CompactTextString(ua) }
func (*Uint32Array) ProtoMessage()               {}
func (*Uint32Array) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// Uint64 represents a Uint64 that may be nil.
type Uint64 struct {
	Uint64 uint64 `protobuf:"varint,1,opt,name=uint64" json:"uint64,omitempty"`
	Valid  bool   `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *Uint64) Reset()                    { *m = Uint64{} }
func (m *Uint64) String() string            { return proto.CompactTextString(m) }
func (*Uint64) ProtoMessage()               {}
func (*Uint64) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// Uint64Array represents an array of uint64s that may be nil.
type Uint64Array struct {
	Uint64Array []uint64 `protobuf:"varint,1,rep,packed,name=uint64_array,json=uint64Array" json:"uint64_array,omitempty"`
	Valid       bool     `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (ua *Uint64Array) Reset()                   { *ua = Uint64Array{} }
func (ua *Uint64Array) String() string           { return proto.CompactTextString(ua) }
func (*Uint64Array) ProtoMessage()               {}
func (*Uint64Array) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// Float32 represents a Float32 that may be nil.
type Float32 struct {
	Float32 float32 `protobuf:"fixed32,1,opt,name=float32" json:"float32,omitempty"`
	Valid   bool    `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *Float32) Reset()                    { *m = Float32{} }
func (m *Float32) String() string            { return proto.CompactTextString(m) }
func (*Float32) ProtoMessage()               {}
func (*Float32) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

// Float32Array represents an array of float32s that may be nil.
type Float32Array struct {
	Float32Array []float32 `protobuf:"fixed32,1,rep,packed,name=float32_array,json=float32Array" json:"float32_array,omitempty"`
	Valid        bool      `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (fa *Float32Array) Reset()                   { *fa = Float32Array{} }
func (fa *Float32Array) String() string           { return proto.CompactTextString(fa) }
func (*Float32Array) ProtoMessage()               {}
func (*Float32Array) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

// Float64 represents a Float64 that may be nil.
type Float64 struct {
	Float64 float64 `protobuf:"fixed64,1,opt,name=float64" json:"float64,omitempty"`
	Valid   bool    `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *Float64) Reset()                    { *m = Float64{} }
func (m *Float64) String() string            { return proto.CompactTextString(m) }
func (*Float64) ProtoMessage()               {}
func (*Float64) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

// Float64Array represents an array of float64s that may be nil.
type Float64Array struct {
	Float64Array []float64 `protobuf:"fixed64,1,rep,packed,name=float64_array,json=float64Array" json:"float64_array,omitempty"`
	Valid        bool      `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (fa *Float64Array) Reset()                   { *fa = Float64Array{} }
func (fa *Float64Array) String() string           { return proto.CompactTextString(fa) }
func (*Float64Array) ProtoMessage()               {}
func (*Float64Array) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

// Bool represents a bool that may be nil.
type Bool struct {
	Bool  bool `protobuf:"varint,1,opt,name=bool" json:"bool,omitempty"`
	Valid bool `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *Bool) Reset()                    { *m = Bool{} }
func (m *Bool) String() string            { return proto.CompactTextString(m) }
func (*Bool) ProtoMessage()               {}
func (*Bool) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

// BoolArray represents an array of booleans that may be nil.
type BoolArray struct {
	BoolArray []bool `protobuf:"varint,1,rep,packed,name=bool_array,json=boolArray" json:"bool_array,omitempty"`
	Valid     bool   `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (ba *BoolArray) Reset()                   { *ba = BoolArray{} }
func (ba *BoolArray) String() string           { return proto.CompactTextString(ba) }
func (*BoolArray) ProtoMessage()               {}
func (*BoolArray) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func init() {
	proto.RegisterType((*String)(nil), "ntypes.String")
	proto.RegisterType((*StringArray)(nil), "ntypes.StringArray")
	proto.RegisterType((*Int32)(nil), "ntypes.Int32")
	proto.RegisterType((*Int32Array)(nil), "ntypes.Int32Array")
	proto.RegisterType((*Int64)(nil), "ntypes.Int64")
	proto.RegisterType((*Int64Array)(nil), "ntypes.Int64Array")
	proto.RegisterType((*Uint32)(nil), "ntypes.Uint32")
	proto.RegisterType((*Uint32Array)(nil), "ntypes.Uint32Array")
	proto.RegisterType((*Uint64)(nil), "ntypes.Uint64")
	proto.RegisterType((*Uint64Array)(nil), "ntypes.Uint64Array")
	proto.RegisterType((*Float32)(nil), "ntypes.Float32")
	proto.RegisterType((*Float32Array)(nil), "ntypes.Float32Array")
	proto.RegisterType((*Float64)(nil), "ntypes.Float64")
	proto.RegisterType((*Float64Array)(nil), "ntypes.Float64Array")
	proto.RegisterType((*Bool)(nil), "ntypes.Bool")
	proto.RegisterType((*BoolArray)(nil), "ntypes.BoolArray")
}

func init() { proto.RegisterFile("ntypes.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x94, 0xc1, 0x6b, 0xf2, 0x40,
	0x10, 0xc5, 0x89, 0x26, 0x31, 0x99, 0xc4, 0x4b, 0xf8, 0xf8, 0xc8, 0xa5, 0x54, 0xed, 0xc5, 0x53,
	0x29, 0x75, 0x59, 0xe8, 0xad, 0xb5, 0x20, 0x78, 0x9d, 0xe2, 0xa5, 0x97, 0x12, 0xdb, 0xda, 0x06,
	0x82, 0x2b, 0x49, 0x2c, 0xf8, 0xdf, 0x97, 0xdd, 0xd9, 0xd5, 0xb5, 0x24, 0xf1, 0xb6, 0xcf, 0x9d,
	0x37, 0xbf, 0xf7, 0x16, 0x0c, 0xc4, 0xdb, 0xfa, 0xb0, 0xfb, 0xac, 0x6e, 0x77, 0xa5, 0xa8, 0x45,
	0xe2, 0x93, 0x9a, 0x30, 0xf0, 0x5f, 0xea, 0x32, 0xdf, 0x7e, 0x25, 0xff, 0xc0, 0x7b, 0xff, 0xce,
	0xca, 0x2a, 0x75, 0x46, 0xce, 0x34, 0x44, 0x12, 0xf2, 0xd7, 0x9f, 0xac, 0xc8, 0x3f, 0xd2, 0xde,
	0xc8, 0x99, 0x06, 0x48, 0x62, 0xb2, 0x80, 0x88, 0x5c, 0x4f, 0x65, 0x99, 0x1d, 0x92, 0x31, 0xc4,
	0x95, 0x92, 0x6f, 0x99, 0xd4, 0xa9, 0x33, 0xea, 0x4f, 0x43, 0x8c, 0x2a, 0x6b, 0xa4, 0x79, 0xcf,
	0x0c, 0xbc, 0xe5, 0xb6, 0x9e, 0xdd, 0xcb, 0xeb, 0x5c, 0x1e, 0x14, 0xdc, 0x43, 0x12, 0x2d, 0xa6,
	0x67, 0x00, 0x65, 0xa2, 0xc5, 0xd7, 0x10, 0xa9, 0x61, 0x0b, 0xed, 0x21, 0xe4, 0xa7, 0x81, 0x2e,
	0x32, 0x67, 0x9a, 0xcc, 0x99, 0x22, 0xf7, 0x91, 0x44, 0x27, 0x99, 0x33, 0x9b, 0xcc, 0x99, 0x45,
	0xee, 0x2b, 0xb2, 0x19, 0x68, 0x5e, 0xc2, 0xc1, 0x5f, 0x51, 0xbd, 0xff, 0xe0, 0xef, 0x4f, 0xad,
	0x87, 0xa8, 0x55, 0xfb, 0x9b, 0xaf, 0xac, 0x5a, 0x63, 0x88, 0xf7, 0x7f, 0x8b, 0x0f, 0x31, 0xda,
	0x5f, 0x6c, 0xae, 0xf9, 0x9c, 0x19, 0xbe, 0xee, 0xee, 0xa2, 0x56, 0xdd, 0x7c, 0x53, 0x4e, 0xf3,
	0xcf, 0xea, 0xbb, 0xc4, 0xef, 0xee, 0xff, 0x00, 0x83, 0x45, 0x21, 0x32, 0x59, 0x34, 0x85, 0xc1,
	0x86, 0x8e, 0x2a, 0x41, 0x0f, 0x8d, 0x6c, 0xb1, 0x2e, 0x21, 0xd6, 0x56, 0x02, 0xdc, 0xc0, 0x50,
	0x1b, 0xac, 0x10, 0x3d, 0x8c, 0x37, 0xf6, 0x50, 0x77, 0x0a, 0xce, 0x8e, 0x29, 0xf4, 0x3b, 0x38,
	0x68, 0xe4, 0x85, 0x14, 0xa6, 0xa6, 0x49, 0x71, 0xf6, 0x14, 0x8e, 0x4e, 0xd1, 0xfd, 0x16, 0x77,
	0xe0, 0xce, 0x85, 0x28, 0x92, 0x04, 0xdc, 0xb5, 0x10, 0x85, 0xe2, 0x07, 0xa8, 0xce, 0x2d, 0x8e,
	0x47, 0x08, 0xa5, 0x83, 0x96, 0x5e, 0x01, 0xc8, 0x51, 0x0b, 0x1b, 0x60, 0xb8, 0x3e, 0x5e, 0x37,
	0x6e, 0x98, 0x07, 0xaf, 0xfa, 0xbf, 0xbf, 0xf6, 0xd5, 0xa7, 0x60, 0xf6, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0xe1, 0xb2, 0x8a, 0x2c, 0x1a, 0x04, 0x00, 0x00,
}
