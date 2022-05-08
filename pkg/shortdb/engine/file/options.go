package file

import (
	"github.com/batazor/shortlink/pkg/shortdb/engine/options"
)

type Option func(file *file) error

func SetPath(path string) options.Option {
	return func(o interface{}) error {
		f := o.(*file)
		f.path = path

		return nil
	}
}

func SetName(name string) options.Option {
	return func(o interface{}) error {
		f := o.(*file)
		f.database.Name = name

		return nil
	}
}
