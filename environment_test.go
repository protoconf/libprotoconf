package libprotoconf

import (
	"os"
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/apipb"
)

func TestConfig_Environment(t *testing.T) {
	type fields struct {
		p proto.Message
	}
	type args struct {
		prefix string
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
			args:    args{},
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
			},
			want:    &apipb.Api{Version: "v1"},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	os.Setenv("TEST_VERSION", "v1")
	os.Setenv("GOOGLE_PROTOBUF_API_VERSION", "v1")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
