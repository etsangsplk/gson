package gson

import "errors"

//**** error codes ****//

// ErrorInvalidDocumentText is returned for misconstructed JSON text.
var ErrorInvalidDocumentText = errors.New("gson.invalidDocumentText")

// ErrorInvalidDocumentType is returned for misconstructed JSON text.
var ErrorInvalidDocumentType = errors.New("gson.invalidDocumentType")

// ErrorInvalidValueType is returned for misconstructed JSON text.
var ErrorInvalidValueType = errors.New("gson.invalidValueType")

// ErrorEmptyDocument is returned when Document does not contain a
// valid value.
var ErrorEmptyDocument = errors.New("gson.emptyDocument")

// ErrorEmptyText to scan
var ErrorEmptyText = errors.New("gson.emptyText")

// ErrorExpectedNil expected a `nil` token while scanning.
var ErrorExpectedNil = errors.New("gson.exptectedNil")

// ErrorExpectedTrue expected a `true` token while scanning.
var ErrorExpectedTrue = errors.New("gson.exptectedTrue")

// ErrorExpectedFalse expected a `false` token while scanning.
var ErrorExpectedFalse = errors.New("gson.exptectedFalse")

// ErrorExpectedClosearray expected a `]` token while scanning.
var ErrorExpectedClosearray = errors.New("gson.exptectedCloseArray")

// ErrorExpectedKey expected a `key-string` token while scanning.
var ErrorExpectedKey = errors.New("gson.exptectedKey")

// ErrorExpectedColon expected a `:` token while scanning.
var ErrorExpectedColon = errors.New("gson.exptectedColon")

// ErrorExpectedCloseobject expected a `}` token while scanning.
var ErrorExpectedCloseobject = errors.New("gson.exptectedCloseobject")

// ErrorExpectedToken expected a valid json token while scanning.
var ErrorExpectedToken = errors.New("gson.exptectedToken")

// ErrorExpectedNum expected a `number` token while scanning.
var ErrorExpectedNum = errors.New("gson.exptectedNum")

// ErrorExpectedString expected a `string` token while scanning.
var ErrorExpectedString = errors.New("gson.exptectedString")

var MetaFields = []string{"_id"}
