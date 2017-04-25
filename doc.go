// Package gson provides a toolkit for JSON representation, collation
// and transformation.
//
// Package provides APIs to convert data representation from one format
// to another. Supported formats are:
//   * Json
//   * Golang value
//   * CBOR - Consice Binary Object Representation
//   * binary-collation
//
// CBOR:
//
// Package also provides a RFC-7049 (CBOR) implementation, to encode
// golang data into machine friendly binary format. Following golang
// native types are supported:
//   * nil, true, false.
//   * native integer types, and its alias, of all width.
//   * float32, float64.
//   * slice of bytes.
//   * native string.
//   * slice of interface - []interface{}.
//   * map of string to interface{} - map[string]interface{}.
//
// Types from golang's standard library and custom types provided
// by this package that can be encoded using CBOR:
//   * CborTagBytes: a cbor encoded []bytes treated as value.
//   * CborUndefined: encode a data-item as undefined.
//   * CborIndefinite: encode bytes, string, array and map of unspefied length.
//   * CborBreakStop: to encode end of CborIndefinite length item.
//   * CborTagEpoch: in seconds since epoch.
//   * CborTagEpochMicro: in micro-seconds epoch.
//   * CborTagFraction: m*(10**e)
//   * CborTagFloat: m*(2**e)
//   * CborTagPrefix: to self indentify a binary blog as CBOR.
//
// Package also provides an implementation for encoding json to CBOR
// and vice-versa:
//   * number can be encoded as integer or float.
//   * arrays and maps are encoded using indefinite encoding.
//   * byte-string encoding is not used.
//
// Json-Pointer:
//
// Package also provides a RFC-6901 (JSON-pointers) implementation.
package gson