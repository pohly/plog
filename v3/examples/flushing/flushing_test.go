package main

import (
	"flag"
	"testing"

	"go.uber.org/goleak"

	"k8s.io/klog/examples/util/require"
	"github.com/pohly/plog/v3"
)

func main() {
	plog.InitFlags(nil)

	// By default klog writes to stderr. Setting logtostderr to false makes klog
	// write to a log file.
	require.NoError(flag.Set("logtostderr", "false"))
	require.NoError(flag.Set("log_file", "myfile.log"))
	flag.Parse()

	// Info writes the first log message. When the first log file is created,
	// a flushDaemon is started to frequently flush bytes to the file.
	plog.Info("nice to meet you")

	// klog won't ever stop this flushDaemon. To exit without leaking a goroutine,
	// the daemon can be stopped manually.
	plog.StopFlushDaemon()

	// After you stopped the flushDaemon, you can still manually flush.
	plog.Info("bye")
	plog.Flush()
}

func TestLeakingFlushDaemon(t *testing.T) {
	// goleak detects leaking goroutines.
	defer goleak.VerifyNone(t)

	// Without calling StopFlushDaemon in main, this test will fail.
	main()
}
