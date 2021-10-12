package libprotoconf

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	v1 "github.com/protoconf/libprotoconf/config/v1"
	testdata "github.com/protoconf/libprotoconf/testdata/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/apipb"
)

func TestConfig_LoadFromSystemDir(t *testing.T) {
	type fields struct {
		msg proto.Message
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*v1.LibprotoconfConfig_Loadable
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				msg: &apipb.Api{},
			},
			want: []*v1.LibprotoconfConfig_Loadable{
				&v1.LibprotoconfConfig_Loadable{Priority: 100, Path: "/etc/google/protobuf"},
				&v1.LibprotoconfConfig_Loadable{Priority: 100, Path: "/etc/google"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewConfig(tt.fields.msg)
			if err := c.LoadFromSystemDir(); (err != nil) != tt.wantErr {
				t.Errorf("Config.LoadFromSystemDir() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(c.config.ConfigDirs, tt.want) {
				t.Errorf("Config.LoadFromSystemDir() =  %v, %v", c.config.ConfigDirs, tt.want)
			}
		})
	}
}

func TestConfig_DetectWorkspace(t *testing.T) {
	wd, _ := os.Getwd()
	type fields struct {
		msg proto.Message
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				msg: &testdata.TestConfig{},
			},
			want:    filepath.Join(wd, ".libprotoconf"),
			wantErr: false,
		},
		{
			name: "test_google",
			fields: fields{
				msg: &apipb.Api{},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewConfig(tt.fields.msg)
			got, err := c.DetectWorkspace()
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.DetectWorkspace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Config.DetectWorkspace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_LoadFromWorkspace(t *testing.T) {
	type fields struct {
		msg proto.Message
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				msg: &testdata.TestConfig{},
			},
			wantErr: false,
		},
		{
			name: "test_google",
			fields: fields{
				msg: &apipb.Api{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewConfig(tt.fields.msg)
			if err := c.LoadFromWorkspace(); (err != nil) != tt.wantErr {
				t.Errorf("Config.LoadFromWorkspace() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfig_LoadFromUserDir(t *testing.T) {
	type fields struct {
		msg  proto.Message
		home string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*v1.LibprotoconfConfig_Loadable
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				msg:  &apipb.Api{},
				home: "/tmp",
			},
			want: []*v1.LibprotoconfConfig_Loadable{
				&v1.LibprotoconfConfig_Loadable{Priority: 90, Path: "/tmp/.google/protobuf"},
				&v1.LibprotoconfConfig_Loadable{Priority: 90, Path: "/tmp/.google"},
			},
			wantErr: false,
		},
		{
			name: "empty_home",
			fields: fields{
				msg:  &apipb.Api{},
				home: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("HOME", tt.fields.home)
			c := NewConfig(tt.fields.msg)
			if err := c.LoadFromUserDir(); (err != nil) != tt.wantErr {
				t.Errorf("Config.LoadFromUserDir() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(c.config.ConfigDirs, tt.want) {
				t.Errorf("Config.LoadFromUserDir() =  %v, want %v", c.config.ConfigDirs, tt.want)
			}
		})
	}
}

func TestConfig_LoadFromDefaultDirs(t *testing.T) {
	type fields struct {
		msg  proto.Message
		home string
	}
	tests := []struct {
		name   string
		fields fields
		want   []*v1.LibprotoconfConfig_Loadable
	}{
		{
			name: "test",
			fields: fields{
				msg:  &apipb.Api{},
				home: "/tmp",
			},
			want: []*v1.LibprotoconfConfig_Loadable{
				&v1.LibprotoconfConfig_Loadable{Priority: 100, Path: "/etc/google/protobuf"},
				&v1.LibprotoconfConfig_Loadable{Priority: 100, Path: "/etc/google"},
				&v1.LibprotoconfConfig_Loadable{Priority: 90, Path: "/tmp/.google/protobuf"},
				&v1.LibprotoconfConfig_Loadable{Priority: 90, Path: "/tmp/.google"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewConfig(tt.fields.msg)
			os.Setenv("HOME", tt.fields.home)
			c.LoadFromDefaultDirs()
			if !reflect.DeepEqual(c.config.ConfigDirs, tt.want) {
				t.Errorf("Config.LoadFromDefaultDirs() =  %v, %v", c.config.ConfigDirs, tt.want)
			}
		})
	}
}

func TestConfig_DefaultLogger(t *testing.T) {
	c := NewConfig(&apipb.Api{})
	l := c.DefaultLogger()
	if l.GetSink().Enabled(10) {
		t.Error("verbose logging should be disabled")
	}
	if !l.GetSink().Enabled(1) {
		t.Error("info logging should be enabled")
	}
}

func TestConfig_DebugLogger(t *testing.T) {
	c := NewConfig(&apipb.Api{})
	l := c.DebugLogger()
	if !l.GetSink().Enabled(10) {
		t.Error("verbose logging should be enabled")
	}
}

func TestConfig_iterateFields(t *testing.T) {
	type fields struct {
		msg proto.Message
	}
	type args struct {
		f func(protoreflect.FieldDescriptor) error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "error",
			fields: fields{
				msg: &apipb.Api{},
			},
			args: args{
				f: func(fd protoreflect.FieldDescriptor) error { return fmt.Errorf("intendend error") },
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewConfig(tt.fields.msg)
			if err := c.iterateFields(tt.args.f); (err != nil) != tt.wantErr {
				t.Errorf("Config.iterateFields() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfig_AppendLoadDir(t *testing.T) {
	type fields struct {
		msg proto.Message
	}
	type args struct {
		priority uint32
		paths    []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "bad_path",
			fields: fields{
				msg: &apipb.Api{},
			},
			args: args{
				priority: 10,
				paths:    []string{"/\\"},
			},
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewConfig(tt.fields.msg)
			if err := c.AppendLoadDir(tt.args.priority, tt.args.paths...); (err != nil) != tt.wantErr {
				t.Errorf("Config.AppendLoadDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
