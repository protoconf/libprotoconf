package main

import (
	"fmt"
	"os"

	"github.com/protoconf/libprotoconf"
	testdata "github.com/protoconf/libprotoconf/testdata/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	cfg := &testdata.TestConfig{}
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
