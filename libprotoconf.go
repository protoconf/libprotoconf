package libprotoconf

import (
	"fmt"
	stdlog "log"
	"os"
	"path/filepath"
	"sort"

	"github.com/go-logr/logr"
	"github.com/go-logr/logr/funcr"
	"github.com/go-logr/stdr"
	v1 "github.com/protoconf/libprotoconf/config/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Config struct {
	msg          proto.Message
	Logger       logr.Logger
	envKeyPrefix string
	config       *v1.LibprotoconfConfig
}

func NewConfig(p proto.Message) *Config {
	c := &Config{
		msg:    p,
		Logger: funcr.New(func(prefix, args string) {}, funcr.Options{}),
		config: &v1.LibprotoconfConfig{},
	}
	c.SetEnvKeyPrefix(string(p.ProtoReflect().Descriptor().FullName()))
	return c
}

func (c *Config) LoadFromDefaultDirs() {
	c.LoadFromSystemDir()
	c.LoadFromUserDir()
	c.LoadFromWorkspace()
	c.Logger.V(4).Info("looking for configs in", "dirs", c.config.ConfigDirs)
}

func (c *Config) LoadFromSystemDir() error {
	for _, p := range c.getPackageTree() {
		c.AppendLoadDir(100, filepath.Join("/etc", pkgToPath(p)))
	}
	return nil
}

func (c *Config) LoadFromUserDir() error {
	for _, p := range c.getPackageTree() {
		if home, err := os.UserHomeDir(); err == nil {
			c.AppendLoadDir(90, filepath.Join(home, "."+pkgToPath(p)))
		} else {
			return err
		}
	}
	return nil
}

func (c *Config) LoadFromWorkspace() error {
	ws, err := c.DetectWorkspace()
	if err == nil {
		c.AppendLoadDir(80, ws)
	}
	return err
}

func (c *Config) DetectWorkspace() (string, error) {
	tree := c.getPackageTree()
	root := "." + tree[len(tree)-1]
	cwd, err := filepath.Abs(".")
	if err != nil {
		return "", err
	}
	for cwd != "/" {
		c.Logger.V(4).Info("trying to detect workspace", "root", root, "cwd", cwd, "tree", tree)
		if ret, err := os.Stat(filepath.Join(cwd, root)); err == nil {
			return filepath.Abs(ret.Name())
		}
		cwd = filepath.Dir(cwd)
	}
	return "", fmt.Errorf("not in workspace")
}

func (c *Config) getPackageTree() []string {
	var result []string
	pkg := c.msg.ProtoReflect().Descriptor().ParentFile().Package()
	result = append(result, string(pkg))
	parent := pkg.Parent()
	for parent != "" {
		result = append(result, string(parent))
		parent = parent.Parent()
	}
	return result
}

func (c *Config) AppendLoadDir(priority uint32, paths ...string) error {
	for _, path := range paths {
		if abs, err := filepath.Abs(path); err == nil {
			c.config.ConfigDirs = append(c.config.ConfigDirs, &v1.LibprotoconfConfig_Loadable{Priority: priority, Path: abs})
		} else {
			return err
		}
	}
	sort.Stable(v1.ByPrio(c.config.ConfigDirs))
	return nil
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
