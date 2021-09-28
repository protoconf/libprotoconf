package libprotoconf

import (
	"flag"
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/apipb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/typepb"
	"google.golang.org/protobuf/types/known/wrapperspb"
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
		{
			name: "bool_value_true",
			fields: fields{
				p:    &wrapperspb.BoolValue{},
				args: []string{"-value", "true"},
			},
			want: &wrapperspb.BoolValue{Value: true},
		},
		{
			name: "bool_value_false",
			fields: fields{
				p:    &wrapperspb.BoolValue{},
				args: []string{"-value=false"},
			},
			want: &wrapperspb.BoolValue{Value: false},
		},
		{
			name: "uint32",
			fields: fields{
				p:    &wrapperspb.UInt32Value{},
				args: []string{"-value", "123"},
			},
			want: &wrapperspb.UInt32Value{Value: 123},
		},
		{
			name: "uint64",
			fields: fields{
				p:    &wrapperspb.UInt64Value{},
				args: []string{"-value", "123"},
			},
			want: &wrapperspb.UInt64Value{Value: 123},
		},
		{
			name: "float64",
			fields: fields{
				p:    &wrapperspb.FloatValue{},
				args: []string{"-value", "123.456"},
			},
			want: &wrapperspb.FloatValue{Value: 123.456},
		},
		{
			name: "double",
			fields: fields{
				p:    &wrapperspb.DoubleValue{},
				args: []string{"-value", "123.456"},
			},
			want: &wrapperspb.DoubleValue{Value: 123.456},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewConfig(tt.fields.p)
			c.DebugLogger()
			c.SetLogger(c.Logger.WithName(t.Name()))
			fs := flag.NewFlagSet(tt.name, flag.ContinueOnError)
			c.PopulateFlagSet(fs)
			if fs.Parse(tt.fields.args); !proto.Equal(tt.fields.p, tt.want) {
				t.Errorf("Config.FlagSet() = %v, want %v", tt.fields.p, tt.want)
			}
		})
	}
}
