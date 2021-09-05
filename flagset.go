package libprotoconf

import (
	"flag"
	"fmt"
	"strconv"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type flaggable struct {
	fd  protoreflect.FieldDescriptor
	cfg *Config
}

func (f *flaggable) String() string {
	return ""
}

func (f *flaggable) Set(value string) error {
	f.cfg.Logger.V(4).Info("setting field", "value", value, "kind", f.fd.Kind().GoString())
	switch f.fd.Kind() {
	case protoreflect.BoolKind:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		f.cfg.p.ProtoReflect().Set(f.fd, protoreflect.ValueOfBool(b))
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		i, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return err
		}
		f.cfg.p.ProtoReflect().Set(f.fd, protoreflect.ValueOf(i))
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		f.cfg.p.ProtoReflect().Set(f.fd, protoreflect.ValueOf(i))
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		i, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return err
		}
		f.cfg.p.ProtoReflect().Set(f.fd, protoreflect.ValueOf(i))
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		i, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		f.cfg.p.ProtoReflect().Set(f.fd, protoreflect.ValueOf(i))
	case protoreflect.EnumKind:
		values := f.fd.Enum().Values()
		pname := protoreflect.Name(value)
		v := values.ByName(pname)
		if v == nil {
			return fmt.Errorf(`invalid option %s value for field "%s"`, value, f.fd.Name())
		}

		f.cfg.p.ProtoReflect().Set(f.fd, protoreflect.ValueOfEnum(v.Number()))
	case protoreflect.StringKind:
		f.cfg.p.ProtoReflect().Set(f.fd, protoreflect.ValueOfString(value))
	default:
		return fmt.Errorf("cannot set field %s, kind %s not implemented", f.fd.JSONName(), f.fd.Kind())
	}
	return nil
}

// FlagSet returns a flag.FlagSet instance for the config
func (c *Config) FlagSet() *flag.FlagSet {
	c.Logger.V(4).Info("creating a flagset")
	fs := flag.NewFlagSet(string(c.p.ProtoReflect().Descriptor().FullName()), flag.ExitOnError)
	fs.Usage = func() {
		fs.PrintDefaults()
	}
	// for i := 0; i < fields.Len(); i++ {
	c.iterateFields(func(f protoreflect.FieldDescriptor) error {
		c.Logger.V(4).Info("got field", "name", f.Name(), "fullName", f.FullName(), "type", f.Kind().GoString())
		switch f.Kind() {
		case protoreflect.MessageKind, protoreflect.GroupKind, protoreflect.BytesKind:
			return nil
		case protoreflect.EnumKind:
			fl := &flaggable{fd: f, cfg: c}
			values := f.Enum().Values()
			valueStrings := make([]string, values.Len())
			for j := 0; j < values.Len(); j++ {
				valueStrings = append(valueStrings, string(values.Get(j).Name()))
			}
			fs.Var(fl, f.JSONName(), fmt.Sprintf("type: %s, default: %s, options: %s", f.Kind(), values.Get(int(c.p.ProtoReflect().Get(f).Enum())).Name(), valueStrings))
		default:
			fl := &flaggable{fd: f, cfg: c}
			fs.Var(fl, f.JSONName(), fmt.Sprintf("type: %s, default: %s", f.Kind(), c.p.ProtoReflect().Get(f)))
		}
		return nil
	})
	return fs
}
