package models

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

type Error struct {
	File     string `json:"file"`
	Function string `json:"function"`
	Detail   any    `json:"detail"`
	Line     int    `json:"line"`
	Code     string `json:"code"`
}

func (e *Error) ToJson() []byte {
	b, err := json.MarshalIndent(e, "", " ")
	if err != nil {
		return nil
	}
	return b
}

func (e *Error) PrintConsole() {
	objectJSON := e.ToJson()
	fmt.Printf("%s\n", objectJSON)
}

func GetError(err error) *Error {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	function := runtime.FuncForPC(pc[0])
	file, line := function.FileLine(pc[0])
	return &Error{
		Code:     "code",
		File:     strings.Replace(filepath.Base(file), ".go", "", 1),
		Function: filepath.Base(function.Name()),
		Line:     line,
		Detail:   err.Error(),
	}
}

func UserError(msg string) *Error {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	function := runtime.FuncForPC(pc[0])
	file, line := function.FileLine(pc[0])
	return &Error{
		Code:     "code",
		File:     strings.Replace(filepath.Base(file), ".go", "", 1),
		Function: function.Name(),
		Line:     line,
		Detail:   msg,
	}
}
