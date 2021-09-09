package libprotoconf

import (
	stdlog "log"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/logr/funcr"
	"github.com/go-logr/stdr"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Config struct {
	msg          proto.Message
	Logger       logr.Logger
	envKeyPrefix string
}

func NewConfig(p proto.Message) *Config {
	c := &Config{msg: p, Logger: funcr.New(func(prefix, args string) {}, funcr.Options{})}
	c.SetEnvKeyPrefix(string(p.ProtoReflect().Descriptor().FullName()))
	return c
}

func (c *Config) SetLogger(logger logr.Logger) logr.Logger {
	c.Logger = logger
	return c.Logger
}

func (c *Config) LoggerWithLevel(v int) logr.Logger {
	stdr.SetVerbosity(v)
	return c.SetLogger(
		stdr.NewWithOptions(
			stdlog.New(os.Stderr, "", stdlog.LstdFlags),
			stdr.Options{LogCaller: stdr.None},
		).WithName("libprotoconf"))
}

func (c *Config) DefaultLogger() logr.Logger {
	return c.LoggerWithLevel(1)
}

func (c *Config) DebugLogger() logr.Logger {
	return c.LoggerWithLevel(10)
}

func (c *Config) iterateFields(f func(protoreflect.FieldDescriptor) error) error {
	r := c.msg.ProtoReflect()
	fields := r.Descriptor().Fields()
	c.Logger.V(10).Info("got reflector", "messageType", r.Descriptor().FullName(), "fields", fields.Len())
	for i := 0; i < fields.Len(); i++ {
		err := f(fields.Get(i))
		if err != nil {
			return err
		}
	}
	return nil
}
