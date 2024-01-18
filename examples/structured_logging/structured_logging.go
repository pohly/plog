package main

import (
	"flag"

	"github.com/pohly/plog/v2"
)

// MyStruct will be logged via %+v
type MyStruct struct {
	Name     string
	Data     string
	internal int
}

// MyStringer will be logged as string, with String providing that string.
type MyString MyStruct

func (m MyString) String() string {
	return m.Name + ": " + m.Data
}

func main() {
	plog.InitFlags(nil)
	flag.Parse()

	someData := MyStruct{
		Name:     "hello",
		Data:     "world",
		internal: 42,
	}

	longData := MyStruct{
		Name: "long",
		Data: `Multiple
lines
with quite a bit
of text.`,
	}

	logData := MyStruct{
		Name: "log output from some program",
		Data: `I0000 12:00:00.000000  123456 main.go:42] Starting
E0000 12:00:01.000000  123456 main.go:43] Failed for some reason
`,
	}

	stringData := MyString(longData)

	plog.Infof("someData printed using InfoF: %v", someData)
	plog.Infof("longData printed using InfoF: %v", longData)
	plog.Infof(`stringData printed using InfoF,
with the message across multiple lines:
%v`, stringData)
	plog.Infof("logData printed using InfoF:\n%v", logData)

	plog.Info("=============================================")

	plog.InfoS("using InfoS", "someData", someData)
	plog.InfoS("using InfoS", "longData", longData)
	plog.InfoS(`using InfoS with
the message across multiple lines`,
		"int", 1,
		"stringData", stringData,
		"str", "another value")
	plog.InfoS("using InfoS", "logData", logData)
	plog.InfoS("using InfoS", "boolean", true, "int", 1, "float", 0.1)

	// The Kubernetes recommendation is to start the message with uppercase
	// and not end with punctuation. See
	// https://github.com/kubernetes/community/blob/HEAD/contributors/devel/sig-instrumentation/migration-to-structured-logging.md
	plog.InfoS("Did something", "item", "foobar")
	// Not recommended, but also works.
	plog.InfoS("This is a full sentence.", "item", "foobar")
}
