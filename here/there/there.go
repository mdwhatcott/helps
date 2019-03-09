package there

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// Locate gathers caller info (file, line) on
func Locate(skip int) Location {
	_, file, line, _ := runtime.Caller(skip + 1)
	return Location{file: file, line: line}
}

type Location struct {
	file string
	line int
}

// JoinedFileLine returns a string in the form of line+sep+line.
func (this Location) Joined(sep string) string {
	file, line := this.FileLine()
	return fmt.Sprintf("%s%s%d", file, sep, line)
}

// FileLine returns the file and line as separate return values.
func (this Location) FileLine() (string, int) { return this.file, this.line }

// Line returns the line number of the caller.
func (this Location) Line() int { return this.line }

// File returns the file path of the caller.
func (this Location) File() string { return this.file }

// Dir returns the directory containing the file of the caller.
func (this Location) Dir() string { return filepath.Dir(this.File()) }
