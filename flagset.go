package libprotoconf

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type flaggable struct {
	fd  protoreflect.FieldDescriptor
	cfg *Config
}

func (f *flaggable) String() string {
	if f.fd == nil || f.cfg == nil {
		return ""
	}
	f.cfg.Logger.V(10).Info("returning string", "fd", f.fd.FullName())
	if f.fd.Kind() == protoreflect.EnumKind {
		values := f.fd.Enum().Values()
		return string(values.Get(int(f.cfg.msg.ProtoReflect().Get(f.fd).Enum())).Name())
	}
	v := f.cfg.msg.ProtoReflect().Get(f.fd)
	if v.IsValid() {
		return fmt.Sprintf("%v", v)
	}
	return ""
}

func (f *flaggable) IsBoolFlag() bool {
	return f.fd.Kind() == protoreflect.BoolKind
}

func (f *flaggable) Set(value string) error {
	f.cfg.Logger.V(4).Info("setting field", "value", value, "kind", f.fd.Kind().GoString())
	switch f.fd.Kind() {
	case protoreflect.BoolKind:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		f.cfg.msg.ProtoReflect().Set(f.fd, protoreflect.ValueOfBool(b))
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		i, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return err
		}
		f.cfg.msg.ProtoReflect().Set(f.fd, protoreflect.ValueOf(i))
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		f.cfg.msg.ProtoReflect().Set(f.fd, protoreflect.ValueOf(i))
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		i, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return err
		}
		f.cfg.msg.ProtoReflect().Set(f.fd, protoreflect.ValueOf(i))
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		i, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		f.cfg.msg.ProtoReflect().Set(f.fd, protoreflect.ValueOf(i))
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		i, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		f.cfg.msg.ProtoReflect().Set(f.fd, protoreflect.ValueOf(i))
	case protoreflect.EnumKind:
		values := f.fd.Enum().Values()
		pname := protoreflect.Name(value)
		v := values.ByName(pname)
		if v == nil {
			return fmt.Errorf(`invalid option %s value for field "%s"`, value, f.fd.Name())
		}
		f.cfg.Logger.V(10).Info("trying to set enum field", "input", value, "name", pname, "result", v.FullName(), "number", v.Number())
		f.cfg.msg.ProtoReflect().Set(f.fd, protoreflect.ValueOfEnum(v.Number()))
	case protoreflect.StringKind:
		f.cfg.msg.ProtoReflect().Set(f.fd, protoreflect.ValueOfString(value))
	default:
		return fmt.Errorf("cannot set field %s, kind %s not implemented", f.fd.JSONName(), f.fd.Kind())
	}
	return nil
}

// FlagSet returns a flag.FlagSet instance for the config
func (c *Config) DefaultFlagSet() *flag.FlagSet {
	fs := flag.NewFlagSet(string(c.msg.ProtoReflect().Descriptor().FullName()), flag.ExitOnError)
	c.PopulateFlagSet(fs)
	return fs
}

func (c *Config) PopulateFlagSet(fs *flag.FlagSet) {
	c.Logger.V(4).Info("creating a flagset")
	fs.Usage = func() {
		fs.PrintDefaults()
	}
	c.iterateFields(func(f protoreflect.FieldDescriptor) error {
		c.Logger.V(4).Info("got field", "name", f.Name(), "fullName", f.FullName(), "type", f.Kind().GoString())
		switch f.Kind() {
		case protoreflect.MessageKind, protoreflect.GroupKind, protoreflect.BytesKind:
			return nil
		case protoreflect.EnumKind:
			fl := &flaggable{fd: f, cfg: c}
			values := f.Enum().Values()
			var valueStrings []string
			for j := 0; j < values.Len(); j++ {
				if str := string(values.Get(j).Name()); str != "" {
					c.Logger.V(10).Info("got enum option", "len", values.Len(), "j", j, "value", str)
					valueStrings = append(valueStrings, str)
				}
			}
			fs.Var(fl, f.JSONName(), fmt.Sprintf(
				"env key: %s\ntype: %s, options: [%s]",
				toEnvKey(c.envKeyPrefix, string(f.Name())),
				f.Kind(),
				strings.Join(valueStrings, ", ")),
			)
		default:
			fl := &flaggable{fd: f, cfg: c}
			fs.Var(fl, f.JSONName(), fmt.Sprintf(
				"env key: %s\ntype: %s",
				toEnvKey(c.envKeyPrefix, string(f.Name())),
				f.Kind(),
			))
		}
		return nil
	})
}
