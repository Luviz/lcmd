package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"go.yaml.in/yaml/v3"
)

type OutputFormat = string

var (
	TEXT OutputFormat = "text"
	JSON OutputFormat = "json"
	YAML OutputFormat = "yaml"
)

type StdError struct {
	Message string `json:"message"`
}

func (e StdError) String() string {
	return e.Message
}

type StdWriter[T fmt.Stringer] struct {
	data   []T
	err    error
	format OutputFormat
}

func ToArray[T any](data T) []T {
	return []T{data}
}

func New[T fmt.Stringer](data []T, format OutputFormat) StdWriter[T] {
	return StdWriter[T]{
		data:   data,
		err:    nil,
		format: format,
	}
}

func Error(err error, msg string, format OutputFormat) StdWriter[StdError] {
	return StdWriter[StdError]{
		data:   ToArray(StdError{Message: msg}),
		err:    err,
		format: format,
	}
}

func (w StdWriter[T]) fatal(err error) {
	w.err = err
}

func (w StdWriter[T]) yaml() string {
	o, err := yaml.Marshal(w.data)
	w.fatal(err)
	return string(o)

}

func (w StdWriter[T]) json() string {
	o, err := json.MarshalIndent(w.data, "", "  ")
	w.fatal(err)
	return string(o)
}

func (w StdWriter[T]) String() string {
	out := []string{}
	switch w.format {
	case TEXT:
		for _, v := range w.data {
			out = append(out, v.String())
		}
		out = append(out, "\n")
	case JSON:
		out = append(out, w.json())
	case YAML:
		out = append(out, w.yaml())
	}
	return strings.Join(out, "\n")
}

func (w StdWriter[T]) Write() {
	s := w.String()
	os.Stdout.Write([]byte(s))
	if w.err != nil {
		os.Stderr.Write([]byte(w.err.Error()))
	}
}
