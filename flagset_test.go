package libprotoconf

import (
	"flag"
	"reflect"
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
		name    string
		fields  fields
		want    proto.Message
		wantErr bool
	}{
		{
			name: "apipb",
			fields: fields{
				p:    &apipb.Api{},
				args: []string{"-name", "libprotoconf", "-version", "v1", "-syntax", "SYNTAX_PROTO3"},
			},
			want:    &apipb.Api{Name: "libprotoconf", Version: "v1", Syntax: typepb.Syntax_SYNTAX_PROTO3},
			wantErr: false,
		},
		{
			name: "apipb_default_enum",
			fields: fields{
				p:    &apipb.Api{},
				args: []string{"-name", "libprotoconf", "-version", "v1", "-syntax", "SYNTAX_PROTO2"},
			},
			want:    &apipb.Api{Name: "libprotoconf", Version: "v1"},
			wantErr: false,
		},
		{
			name: "apipb_bad_enum",
			fields: fields{
				p:    &apipb.Api{},
				args: []string{"-name", "libprotoconf", "-version", "v1", "-syntax", "SYNTAX_PROTO4"},
			},
			want:    &apipb.Api{Name: "libprotoconf", Version: "v1"},
			wantErr: true,
		},
		{
			name: "timestamppb",
			fields: fields{
				p:    timestamppb.Now(),
				args: []string{"-nanos", "123", "-seconds", "456"},
			},
			want:    &timestamppb.Timestamp{Seconds: 456, Nanos: 123},
			wantErr: false,
		},
		{
			name: "help",
			fields: fields{
				p:    &apipb.Api{},
				args: []string{"-h"},
			},
			want:    &apipb.Api{},
			wantErr: true,
		},
		{
			name: "bool_true",
			fields: fields{
				p:    &wrapperspb.BoolValue{},
				args: []string{"-value", "true"},
			},
			want:    &wrapperspb.BoolValue{Value: true},
			wantErr: false,
		},
		{
			name: "bool_false",
			fields: fields{
				p:    &wrapperspb.BoolValue{},
				args: []string{"-value=false"},
			},
			want:    &wrapperspb.BoolValue{Value: false},
			wantErr: false,
		},
		{
			name: "bad_bool",
			fields: fields{
				p:    &wrapperspb.BoolValue{},
				args: []string{"-value=hello"},
			},
			want:    &wrapperspb.BoolValue{Value: false},
			wantErr: true,
		},
		{
			name: "uint32",
			fields: fields{
				p:    &wrapperspb.UInt32Value{},
				args: []string{"-value", "123"},
			},
			want:    &wrapperspb.UInt32Value{Value: 123},
			wantErr: false,
		},
		{
			name: "bad_uint32",
			fields: fields{
				p:    &wrapperspb.UInt32Value{},
				args: []string{"-value", "abc"},
			},
			want:    &wrapperspb.UInt32Value{},
			wantErr: true,
		},
		{
			name: "bad_int32",
			fields: fields{
				p:    &wrapperspb.Int32Value{},
				args: []string{"-value", "abc"},
			},
			want:    &wrapperspb.Int32Value{},
			wantErr: true,
		},
		{
			name: "uint64",
			fields: fields{
				p:    &wrapperspb.UInt64Value{},
				args: []string{"-value", "123"},
			},
			want:    &wrapperspb.UInt64Value{Value: 123},
			wantErr: false,
		},
		{
			name: "bad_uint64",
			fields: fields{
				p:    &wrapperspb.UInt64Value{},
				args: []string{"-value", "hello"},
			},
			want:    &wrapperspb.UInt64Value{},
			wantErr: true,
		},
		{
			name: "bad_int64",
			fields: fields{
				p:    &wrapperspb.Int64Value{},
				args: []string{"-value", "hello"},
			},
			want:    &wrapperspb.Int64Value{},
			wantErr: true,
		},
		{
			name: "float64",
			fields: fields{
				p:    &wrapperspb.FloatValue{},
				args: []string{"-value", "123.456"},
			},
			want:    &wrapperspb.FloatValue{Value: 123.456},
			wantErr: false,
		},
		{
			name: "int_to_float64",
			fields: fields{
				p:    &wrapperspb.FloatValue{},
				args: []string{"-value", "123"},
			},
			want:    &wrapperspb.FloatValue{Value: 123},
			wantErr: false,
		},
		{
			name: "bad_float64",
			fields: fields{
				p:    &wrapperspb.FloatValue{},
				args: []string{"-value", "abc"},
			},
			want:    &wrapperspb.FloatValue{},
			wantErr: true,
		},
		{
			name: "double",
			fields: fields{
				p:    &wrapperspb.DoubleValue{},
				args: []string{"-value", "123.456"},
			},
			want:    &wrapperspb.DoubleValue{Value: 123.456},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewConfig(tt.fields.p)
			c.DebugLogger()
			c.SetLogger(c.Logger.WithName(t.Name()))
			fs := flag.NewFlagSet(tt.name, flag.ContinueOnError)
			c.PopulateFlagSet(fs)
			if err := fs.Parse(tt.fields.args); (err != nil) != tt.wantErr {
				t.Errorf("Config.FlagSet() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !proto.Equal(tt.fields.p, tt.want) {
				t.Errorf("Config.FlagSet() = %v, want %v", tt.fields.p, tt.want)
			}
		})
	}
}

func TestConfig_DefaultFlagSet(t *testing.T) {
	type fields struct {
		msg proto.Message
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test",
			fields: fields{
				msg: &apipb.Api{},
			},
			want: "google.protobuf.Api",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewConfig(tt.fields.msg)
			fs := c.DefaultFlagSet()
			if got := fs.Name(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.DefaultFlagSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
