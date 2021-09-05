package libprotoconf

import (
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
		wantErr bool
	}{
		{
			name: "api",
			fields: fields{
				p: &apipb.Api{},
			},
			args: args{
				prefix: "test",
			},
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				p: tt.fields.p,
			}
			c.DebugLogger()
			if err := c.Environment(tt.args.prefix); (err != nil) != tt.wantErr {
				t.Errorf("Config.Environment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
