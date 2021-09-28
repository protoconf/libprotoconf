package main

import (
	"fmt"
	"os"

	"github.com/protoconf/libprotoconf"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/apipb"
)

func main() {
	cfg := &apipb.Api{
		Name:    "protoconf",
		Version: "v1",
	}
	// cfg := &v1.LibprotoconfConfig{}
	lpc := libprotoconf.NewConfig(cfg)
	lpc.DebugLogger()
	fs := lpc.DefaultFlagSet()
	lpc.Logger.Info("starting", "args", os.Args)
	lpc.LoadFromDefaultDirs()
	lpc.Environment()
	fs.Parse(os.Args[1:])
	if fs.Parsed() {
		lpc.Logger.Info("result", "config", fmt.Sprintf("%v", cfg))
		fmt.Print(protojson.Format(cfg))
	} else {
		fs.Usage()
	}
}
