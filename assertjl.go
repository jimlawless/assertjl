// Copyright 2021 - by Jim Lawless
// License: MIT / X11
// See: https://github.com/jimlawless/assertjl/LICENSE
//

// I borrowed some code from my other project "whereami" for this package.
// See: https://github.com/jimlawless/whereami

package assertjl

import (
	"fmt"
    "os"
	"runtime"
	"strings"
)


// conditionally display on Stderr a string containing the file name, function name
// and the line number of a specified entry on the call stack *if* the boolean 
// condition is false
func Assert(cond bool,depthList ...int) {
	var depth int
    if cond {
        return
    }
  
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}
	function, file, line, _ := runtime.Caller(depth)
	fmt.Fprintf(os.Stderr,"Assertion! File: %s  Function: %s Line: %d", chopPath(file), runtime.FuncForPC(function).Name(), line)
    os.Exit(0)
}

// return the source filename after the last slash
func chopPath(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	} else {
		return original[i+1:]
	}
}
