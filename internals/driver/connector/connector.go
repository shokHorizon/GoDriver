package connector

import "io"

type Connector interface {
	io.Reader
	io.Writer
	io.Closer
}
