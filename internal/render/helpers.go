package render

import (
	"fmt"
	"strings"

	"github.com/cweill/gotests/internal/models"
)

const nFile = 7 // Number of files to be read from template (package) template (directory)

func fieldName(f *models.Field) string {
	var n string
	if f.IsNamed() {
		n = f.Name
	} else {
		n = f.Type.String()
	}
	return n
}

func receiverName(f *models.Receiver) string {
	var n string
	if f.IsNamed() {
		n = f.Name
	} else {
		n = f.ShortName()
	}
	if n == "name" {
		// Avoid conflict with test struct's "name" field.
		n = "n"
	}
	if n == "t" {
		// Avoid conflict with test argument.
		// "tr" is short for t receiver.
		n = "tr"
	}
	if n == "f" {
		// Avoid conflict with fuzzing argument (f *testing.F)
		n = "fr"
	}
	return n
}

func parameterName(f *models.Field) string {
	var n string
	if f.IsNamed() {
		n = f.Name
	} else {
		n = fmt.Sprintf("in%v", f.Index)
	}
	return n
}

func defaultValueForType(f *models.Field) string {
	switch f.Type.Underlying {
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return "0"
	case "float32", "float64":
		return "0.0"
	case "string":
		return `"hello"`
	case "byte":
		return `byte(0x01)`
	default:
		// go test fuzz does not support non basic types
		// We can import well known libs to handle that case
		return "[]byte{}"
	}
}

func wantName(f *models.Field) string {
	var n string
	if f.IsNamed() {
		n = "want" + strings.Title(f.Name)
	} else if f.Index == 0 {
		n = "want"
	} else {
		n = fmt.Sprintf("want%v", f.Index)
	}
	return n
}

func gotName(f *models.Field) string {
	var n string
	if f.IsNamed() {
		n = "got" + strings.Title(f.Name)
	} else if f.Index == 0 {
		n = "got"
	} else {
		n = fmt.Sprintf("got%v", f.Index)
	}
	return n
}
