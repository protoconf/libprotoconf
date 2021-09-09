package libprotoconf

import (
	"os"
	"regexp"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func (c *Config) SetEnvKeyPrefix(p string) {
	c.envKeyPrefix = p
}

func (c *Config) Environment() error {
	return c.iterateFields(func(fd protoreflect.FieldDescriptor) error {
		envName := toEnvKey(c.envKeyPrefix, string(fd.Name()))
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

func toEnvKey(strs ...string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	envName := strings.ToUpper(strings.Join(strs, "_"))
	return reg.ReplaceAllString(envName, "_")

}
