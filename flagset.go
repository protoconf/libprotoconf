package libprotoconf

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
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
		if f.fd.IsList() {
			list := f.cfg.msg.ProtoReflect().Get(f.fd).List()
			var arr []string
			for i := 0; i < list.Len(); i++ {
				v := list.Get(i)
				if v.IsValid() {
					arr = append(arr, fmt.Sprintf("%v", v))
				}
			}
			return fmt.Sprintf("%v", arr)
		}
		return string(values.Get(int(f.cfg.msg.ProtoReflect().Get(f.fd).Enum())).Name())
	}
	if f.fd.IsList() {
		list := f.cfg.msg.ProtoReflect().Get(f.fd).List()
		var arr []string
		for i := 0; i < list.Len(); i++ {
			v := list.Get(i)
			if v.IsValid() {
				arr = append(arr, fmt.Sprintf("%v", v))
			}
		}
		return fmt.Sprintf("[%v]", strings.Join(arr, ", "))
	}
	v := f.cfg.msg.ProtoReflect().Get(f.fd)
	return fmt.Sprintf("%v", v)
}

func (f *flaggable) IsBoolFlag() bool {
	return f.fd.Kind() == protoreflect.BoolKind
}

func (f *flaggable) detect(value string) (protoreflect.Value, error) {
	switch f.fd.Kind() {
	case protoreflect.BoolKind:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return protoreflect.Value{}, err
		}
		return protoreflect.ValueOfBool(b), nil
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		i, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return protoreflect.Value{}, err
		}
		return protoreflect.ValueOf(i), nil
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return protoreflect.Value{}, err
		}
		return protoreflect.ValueOf(i), nil
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		i, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return protoreflect.Value{}, err
		}
		return protoreflect.ValueOf(i), nil
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		i, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return protoreflect.Value{}, err
		}
		return protoreflect.ValueOf(i), nil
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		i, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return protoreflect.Value{}, err
		}
		return protoreflect.ValueOf(i), nil
	case protoreflect.EnumKind:
		values := f.fd.Enum().Values()
		pname := protoreflect.Name(value)
		v := values.ByName(pname)
		if v == nil {
			return protoreflect.Value{}, fmt.Errorf(`invalid option %s value for field "%s"`, value, f.fd.Name())
		}
		f.cfg.Logger.V(10).Info("trying to set enum field", "input", value, "name", pname, "result", v.FullName(), "number", v.Number())
		return protoreflect.ValueOfEnum(v.Number()), nil
	case protoreflect.StringKind:
		return protoreflect.ValueOfString(value), nil
	}
	return protoreflect.Value{}, fmt.Errorf("cannot set field %s, kind %s not implemented", f.fd.JSONName(), f.fd.Kind())
}

func (f *flaggable) Set(value string) error {
	f.cfg.Logger.V(4).Info("setting field", "value", value, "kind", f.fd.Kind().GoString())
	if f.fd.IsList() {
		list := f.cfg.msg.ProtoReflect().Mutable(f.fd).List()
		for _, item := range strings.Split(value, ",") {
			v, err := f.detect(strings.TrimSpace(item))
			if err != nil {
				return err
			}
			list.Append(v)
		}
		return nil
	}
	v, err := f.detect(value)
	if err != nil {
		return err
	}
	f.cfg.msg.ProtoReflect().Set(f.fd, v)
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
		case protoreflect.GroupKind, protoreflect.BytesKind:
			return nil
		case protoreflect.MessageKind:
			msgDesc := f.Message()
			l := NewConfig(dynamicpb.NewMessage(msgDesc))
			l.SetEnvKeyPrefix(strings.Join([]string{c.envKeyPrefix, toEnvKey(f.JSONName())}, "_"))
			l.DefaultFlagSet().VisitAll(func(fl *flag.Flag) {
				fs.Var(fl.Value, strings.Join([]string{f.JSONName(), fl.Name}, "-"), fl.Usage)

			})
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
