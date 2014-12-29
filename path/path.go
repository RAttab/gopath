// Copyright (c) 2014 Datacratic. All rights reserved.

package path

import (
	"bytes"
	"strings"
)

// P represents a path through an object seperated by '.' characters. A path can
// also contain wildcard components indicated by a '*' character. Arrays and
// slice indexes should be specified using non-negative numbers. Only map keyed
// with string are currently supported. Channels can be read by providing either
// a number of values to read or a wildcard character to read all values until
// the channel is closed.
type P []string

// New returns a new P object from a given path string.
func New(path string) P {
	return strings.Split(path, ".")
}

// String returns a simple string representation of the path.
func (path P) String() string {
	buffer := new(bytes.Buffer)

	for i, item := range path {
		buffer.WriteString(item)

		if i < len(path)-1 {
			buffer.WriteString(".")
		}
	}

	return buffer.String()
}

// Last returns the last component of the path.
func (path P) Last() string { return path[len(path)-1] }
