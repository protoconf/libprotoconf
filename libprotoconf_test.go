package libprotoconf

import (
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
