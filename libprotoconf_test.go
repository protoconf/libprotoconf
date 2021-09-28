package libprotoconf

import (
	"os"
	"reflect"
	"sort"
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/apipb"
)

func TestConfig_LoadFromSystemDir(t *testing.T) {
	type fields struct {
		msg proto.Message
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				msg: &apipb.Api{},
			},
			want:    []string{"/etc/google", "/etc/google/protobuf"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewConfig(tt.fields.msg)
			if err := c.LoadFromSystemDir(); (err != nil) != tt.wantErr {
				t.Errorf("Config.LoadFromSystemDir() error = %v, wantErr %v", err, tt.wantErr)
			}
			sort.Strings(c.config.ConfigDirs)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(c.config.ConfigDirs, tt.want) {
				t.Errorf("Config.LoadFromSystemDir() =  %v, %v", c.config.ConfigDirs, tt.want)
			}
		})
	}
}

func TestConfig_DetectWorkspace(t *testing.T) {
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
			name:    "test",
			fields:  fields{msg: &apipb.Api{}},
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
		want    []string
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				msg:  &apipb.Api{},
				home: "/tmp",
			},
			want:    []string{"/tmp/.google", "/tmp/.google/protobuf"},
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
			sort.Strings(c.config.ConfigDirs)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(c.config.ConfigDirs, tt.want) {
				t.Errorf("Config.LoadFromUserDir() =  %v, %v", c.config.ConfigDirs, tt.want)
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
		want   []string
	}{
		{
			name: "test",
			fields: fields{
				msg:  &apipb.Api{},
				home: "/tmp",
			},
			want: []string{"/etc/google", "/etc/google/protobuf", "/tmp/.google", "/tmp/.google/protobuf"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewConfig(tt.fields.msg)
			os.Setenv("HOME", tt.fields.home)
			c.LoadFromDefaultDirs()
			sort.Strings(c.config.ConfigDirs)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(c.config.ConfigDirs, tt.want) {
				t.Errorf("Config.LoadFromDefaultDirs() =  %v, %v", c.config.ConfigDirs, tt.want)
			}
		})
	}
}
