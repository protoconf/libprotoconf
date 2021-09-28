package libprotoconf

import (
	"encoding/base64"
	"fmt"
	"path/filepath"

	"github.com/google/go-jsonnet"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
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
	}
	return fmt.Errorf("cannot detect unmarshaller")
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
