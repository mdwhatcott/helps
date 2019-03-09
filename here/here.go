package here

import "github.com/mdwhatcott/helps/here/there"

func JoinedFileLine(sep string) string { return there.Locate(1).Joined(sep) }
func FileLine() (string, int)          { return there.Locate(1).FileLine() }
func Line() int                        { return there.Locate(1).Line() }
func File() string                     { return there.Locate(1).File() }
func Dir() string                      { return there.Locate(1).Dir() }
