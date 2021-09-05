package libprotoconf

import (
	"os"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func (c *Config) Environment(prefix string) error {
	return c.iterateFields(func(fd protoreflect.FieldDescriptor) error {
		envName := strings.ToUpper(strings.Join([]string{prefix, fd.TextName()}, "_"))
		c.Logger.V(4).Info("trying to read environtment variable", "variable", envName)
		result := os.Getenv(envName)
		if result != "" {
			c.Logger.V(4).Info("got result", "result", result)
			fl := &flaggable{fd: fd, cfg: c}
			fl.Set(result)
		}
		return nil
	})
}
