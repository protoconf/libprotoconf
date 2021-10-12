package libprotoconf

import (
	"os"
	"testing"

	testdata "github.com/protoconf/libprotoconf/testdata/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/apipb"
)

func TestConfig_Environment(t *testing.T) {
	type fields struct {
		p proto.Message
	}
	type args struct {
		prefix string
		envs   map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    proto.Message
		wantErr bool
	}{
		{
			name: "api",
			fields: fields{
				p: &apipb.Api{},
			},
			args: args{
				envs: map[string]string{
					"GOOGLE_PROTOBUF_API_VERSION": "v1",
				},
			},
			want:    &apipb.Api{Version: "v1"},
			wantErr: false,
		},
		{
			name: "api_with_prefix",
			fields: fields{
				p: &apipb.Api{},
			},
			args: args{
				prefix: "test",
				envs: map[string]string{
					"TEST_VERSION": "v1",
				},
			},
			want:    &apipb.Api{Version: "v1"},
			wantErr: false,
		},
		{
			name: "list",
			fields: fields{
				p: &testdata.TestConfig{},
			},
			args: args{
				envs: map[string]string{
					"LIBPROTOCONF_TESTDATA_V1_TESTCONFIG_STR_ARR": "v1,v2",
				},
			},
			want:    &testdata.TestConfig{StrArr: []string{"v1", "v2"}},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.args.envs {
				os.Setenv(k, v)
			}
			c := NewConfig(tt.fields.p)
			c.DebugLogger()
			c.SetLogger(c.Logger.WithName(t.Name()))
			if tt.args.prefix != "" {
				c.SetEnvKeyPrefix(tt.args.prefix)
			}
			if err := c.Environment(); (err != nil) != tt.wantErr {
				t.Errorf("Config.Environment() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !proto.Equal(tt.fields.p, tt.want) {
				t.Errorf("Config.Environment() = %v, want %v", tt.fields.p, tt.want)
			}
		})
	}
}
