# here
--
    import "github.com/mdwhatcott/here"


## Usage

#### func  Dir

```go
func Dir() string
```
Dir returns the directory containing the file of the caller.

#### func  File

```go
func File() string
```
File returns the file path of the caller.

#### func  FileLine

```go
func FileLine() (string, int)
```
FileLine returns the file and line as separate return values.

#### func  JoinedFileLine

```go
func JoinedFileLine(sep string) string
```
JoinedFileLine returns a string in the form of line+sep+line.

#### func  Line

```go
func Line() int
```
Line returns the line number of the caller.
