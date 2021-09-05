package libprotoconf

import (
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/apipb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/typepb"
)

func TestConfig_FlagSet(t *testing.T) {
	type fields struct {
		p    proto.Message
		args []string
	}

	tests := []struct {
		name   string
		fields fields
		want   proto.Message
	}{
		{
			name: "apipb",
			fields: fields{
				p:    &apipb.Api{},
				args: []string{"-name", "libprotoconf", "-version", "v1", "-syntax", "SYNTAX_PROTO3"},
			},
			want: &apipb.Api{Name: "libprotoconf", Version: "v1", Syntax: typepb.Syntax_SYNTAX_PROTO3},
		},
		{
			name: "apipb_default_enum",
			fields: fields{
				p:    &apipb.Api{},
				args: []string{"-name", "libprotoconf", "-version", "v1", "-syntax", "SYNTAX_PROTO2"},
			},
			want: &apipb.Api{Name: "libprotoconf", Version: "v1"},
		},
		{
			name: "timestamppb",
			fields: fields{
				p:    timestamppb.Now(),
				args: []string{"-nanos", "123", "-seconds", "456"},
			},
			want: &timestamppb.Timestamp{Seconds: 456, Nanos: 123},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewConfig(tt.fields.p)
			c.DebugLogger()
			c.SetLogger(c.Logger.WithName(t.Name()))
			if c.FlagSet().Parse(tt.fields.args); !proto.Equal(tt.fields.p, tt.want) {
				t.Errorf("Config.FlagSet() = %v, want %v", tt.fields.p, tt.want)
			}
		})
	}
}
