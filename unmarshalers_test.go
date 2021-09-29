package libprotoconf

import (
	"encoding/base64"
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/apipb"
	"google.golang.org/protobuf/types/known/sourcecontextpb"
	"google.golang.org/protobuf/types/known/typepb"
)

func TestConfig_Unmarshal(t *testing.T) {
	type fields struct {
		msg proto.Message
	}
	type args struct {
		filename string
		data     []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wants   proto.Message
		wantErr bool
	}{
		{
			name: "json",
			fields: fields{
				msg: &apipb.Api{},
			},
			args:    args{filename: "api.json", data: []byte(`{"name":"protoconf", "version":"v1", "syntax":"SYNTAX_PROTO3"}`)},
			wants:   &apipb.Api{Name: "protoconf", Version: "v1", Syntax: typepb.Syntax_SYNTAX_PROTO3},
			wantErr: false,
		},
		{
			name: "yaml",
			fields: fields{
				msg: &apipb.Api{},
			},
			args:    args{filename: "api.yaml", data: []byte("name: protoconf\nversion: v1")},
			wants:   &apipb.Api{Name: "protoconf", Version: "v1"},
			wantErr: false,
		},
		{
			name: "bad_yaml",
			fields: fields{
				msg: &apipb.Api{},
			},
			args:    args{filename: "api.yaml", data: []byte("name: protoconfversion: v1")},
			wants:   &apipb.Api{},
			wantErr: true,
		},
		{
			name: "pbtext",
			fields: fields{
				msg: &apipb.Api{},
			},
			args:    args{filename: "api.pb", data: []byte(`name: "protoconf" version: "v1"`)},
			wants:   &apipb.Api{Name: "protoconf", Version: "v1"},
			wantErr: false,
		},
		{
			name: "binary",
			fields: fields{
				msg: &apipb.Api{},
			},
			args:    args{filename: "api.data", data: getBinary(&apipb.Api{Name: "protoconf", Version: "v1"})},
			wants:   &apipb.Api{Name: "protoconf", Version: "v1"},
			wantErr: false,
		},
		{
			name: "base64",
			fields: fields{
				msg: &apipb.Api{},
			},
			args: args{filename: "api.base64", data: []byte(base64.URLEncoding.EncodeToString(
				getBinary(&apipb.Api{Name: "protoconf", Version: "v1"})))},
			wants:   &apipb.Api{Name: "protoconf", Version: "v1"},
			wantErr: false,
		},
		{
			name: "bad_base64",
			fields: fields{
				msg: &apipb.Api{},
			},
			args: args{filename: "api.base64", data: []byte(base64.URLEncoding.EncodeToString(
				getBinary(&apipb.Api{Name: "protoconf", Version: "v1"})) + "bad")},
			wants:   &apipb.Api{},
			wantErr: true,
		},
		{
			name: "jsonnet",
			fields: fields{
				msg: &apipb.Api{},
			},
			args:    args{filename: "api.jsonnet", data: []byte(`{name: "protoconf", version: self.name + "/v1"}`)},
			wants:   &apipb.Api{Name: "protoconf", Version: "protoconf/v1"},
			wantErr: false,
		},
		{
			name: "bad_jsonnet",
			fields: fields{
				msg: &apipb.Api{},
			},
			args:    args{filename: "api.jsonnet", data: []byte(`{name: "protoconf", version: self.name + "/v1"`)},
			wants:   &apipb.Api{},
			wantErr: true,
		},
		{
			name: "dummy_extension",
			fields: fields{
				msg: &apipb.Api{},
			},
			args:    args{filename: "api.dummy", data: []byte(`blah`)},
			wants:   &apipb.Api{},
			wantErr: true,
		},
		{
			name: "toml",
			fields: fields{
				msg: &apipb.Api{},
			},
			args: args{filename: "api.toml", data: []byte(`
			name = "protoconf"
			version = "v1"
			methods = [
				{name = "GET"}
			]
			source_context = {
				file_name = "hello"
			}
			`)},
			wants: &apipb.Api{
				Name:    "protoconf",
				Version: "v1",
				Methods: []*apipb.Method{
					{Name: "GET"},
				},
				SourceContext: &sourcecontextpb.SourceContext{FileName: "hello"},
			},
			wantErr: false,
		},
		{
			name: "bad_toml",
			fields: fields{
				msg: &apipb.Api{},
			},
			args: args{filename: "api.toml", data: []byte(`
			name = "protoconf
			`)},
			wants:   &apipb.Api{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewConfig(tt.fields.msg)
			c.DebugLogger()
			c.SetLogger(c.Logger.WithName(t.Name()))
			if err := c.Unmarshal(tt.args.filename, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Config.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !proto.Equal(c.msg, tt.wants) {
				t.Errorf("Config.Unmarshal() wanted = %v, got = %v", tt.wants, c.msg)
			}
		})
	}
}

func getBinary(msg proto.Message) []byte {
	// This should never fail
	t, _ := proto.Marshal(msg)
	return t
}
