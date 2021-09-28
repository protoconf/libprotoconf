package libprotoconf

import (
	"flag"
	"testing"

	pb "github.com/protoconf/protoconf/datatypes/proto/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
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
		{
			name: "help",
			fields: fields{
				p:    &apipb.Api{},
				args: []string{"-h"},
			},
			want: &apipb.Api{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewConfig(tt.fields.p)
			c.DebugLogger()
			c.Environment()
			c.SetLogger(c.Logger.WithName(t.Name()))
			fs := flag.NewFlagSet(tt.name, flag.ContinueOnError)
			c.PopulateFlagSet(fs)
			if fs.Parse(tt.fields.args); !proto.Equal(tt.fields.p, tt.want) {
				t.Errorf("Config.FlagSet() = %v, want %v", tt.fields.p, tt.want)
			}
		})
	}
}

func TestMarshalAny(t *testing.T) {

	ts := &apipb.Api{Name: "protoconf", Version: "v1", Syntax: typepb.Syntax_SYNTAX_PROTO3}
	a, _ := anypb.New(ts)
	v := &pb.ProtoconfValue{ProtoFile: "hello.proto", Value: a}
	t.Log(protojson.Format(v))
	t.Fail()
}
