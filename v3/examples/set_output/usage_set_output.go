package main

import (
	"bytes"
	"flag"
	"fmt"

	"k8s.io/klog/examples/util/require"
	"github.com/pohly/plog/v3"
)

func main() {
	plog.InitFlags(nil)
	require.NoError(flag.Set("logtostderr", "false"))
	require.NoError(flag.Set("alsologtostderr", "false"))
	flag.Parse()

	buf := new(bytes.Buffer)
	plog.SetOutput(buf)
	plog.Info("nice to meet you")
	plog.Flush()

	fmt.Printf("LOGGED: %s", buf.String())
}
