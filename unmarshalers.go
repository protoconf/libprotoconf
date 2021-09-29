package libprotoconf

import (
	"encoding/base64"
	"fmt"
	"path/filepath"

	"github.com/google/go-jsonnet"
	"github.com/pelletier/go-toml"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"sigs.k8s.io/yaml"
)

func (c *Config) Unmarshal(filename string, data []byte) error {
	switch filepath.Ext(filename) {
	case ".json":
		return c.UnmarshalJSON(data)
	case ".yaml", ".yml":
		return c.UnmarshalYAML(data)
	case ".pb":
		return c.UnmarshalText(data)
	case ".data":
		return c.UnmarshalBinary(data)
	case ".b64", ".base64":
		return c.UnmarshalBase64(data)
	case ".jsonnet":
		return c.UnmarshalJsonnet(filename, data)
	case ".toml":
		return c.UnmarshalToml(filename, data)
	}
	return fmt.Errorf("cannot detect unmarshaller for: %s", filename)
}

func (c *Config) UnmarshalJSON(data []byte) error {
	c.Logger.V(4).Info("json data", "data", string(data))
	return protojson.Unmarshal(data, c.msg)
}

func (c *Config) UnmarshalText(data []byte) error {
	c.Logger.V(4).Info("text data", "data", string(data))
	return prototext.Unmarshal(data, c.msg)
}

func (c *Config) UnmarshalBinary(data []byte) error {
	c.Logger.V(4).Info("binary data", "data", string(data))
	return proto.Unmarshal(data, c.msg)
}

func (c *Config) UnmarshalYAML(data []byte) error {
	c.Logger.V(4).Info("yaml data", "data", string(data))
	if b, err := yaml.YAMLToJSONStrict(data); err != nil {
		return err
	} else {
		return c.UnmarshalJSON(b)
	}
}

func (c *Config) UnmarshalBase64(data []byte) error {
	c.Logger.V(4).Info("base64 data", "data", string(data))
	target, err := base64.URLEncoding.DecodeString(string(data))
	if err != nil {
		return err
	}
	return c.UnmarshalBinary(target)
}

func (c *Config) UnmarshalJsonnet(filename string, data []byte) error {
	vm := jsonnet.MakeVM()
	jsonStr, err := vm.EvaluateAnonymousSnippet(filename, string(data))
	if err != nil {
		return err
	}
	return c.UnmarshalJSON([]byte(jsonStr))
}

func (c *Config) UnmarshalToml(filename string, data []byte) error {
	t, err := toml.LoadBytes(data)
	if err != nil {
		return err
	}
	return unmarshalInterface(c.msg.ProtoReflect(), t.ToMap())
}

func unmarshalInterface(msg protoreflect.Message, data map[string]interface{}) error {
	// TODO(smintz): support Map
	// TODO(smintz): test scalar list types
	return iterateFields(msg, func(fd protoreflect.FieldDescriptor) error {
		fieldName := fd.Name()
		if v, ok := data[string(fieldName)]; ok {
			if fd.Kind() == protoreflect.MessageKind {
				if fd.IsList() {
					list := msg.Mutable(fd).List()
					for _, vv := range v.([]interface{}) {
						newmsg := list.NewElement().Message().New()
						unmarshalInterface(newmsg, vv.(map[string]interface{}))
						list.Append(protoreflect.ValueOfMessage(newmsg))

					}
					return nil
				}
				newmsg := msg.Get(fd).Message().New()
				unmarshalInterface(newmsg, v.(map[string]interface{}))
				msg.Set(fd, protoreflect.ValueOfMessage(newmsg))
				return nil
			}
			msg.Set(fd, protoreflect.ValueOf(v))
		}
		return nil
	})

}
