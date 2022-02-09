/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package test

import (
	"io"
	"testing"

	"github.com/go-logr/logr"

	"k8s.io/klog/v2/klogr"
)

func init() {
	InitKlog()
}

// TestKlogOutput tests klog output without a logger.
func TestKlogOutput(t *testing.T) {
	Output(t, OutputConfig{})
}

// TestKlogrOutput tests klogr output via klog.
func TestKlogrOutput(t *testing.T) {
	// klogr currently doesn't produce exactly the same output as klog.
	// TODO: fix that.
	mapping := map[string]string{
		`I output.go:<LINE>] "test" keyWithoutValue="(MISSING)"
I output.go:<LINE>] "test" keyWithoutValue="(MISSING)" anotherKeyWithoutValue="(MISSING)"
I output.go:<LINE>] "test" keyWithoutValue="(MISSING)"
`: `I output.go:<LINE>] "test" keyWithoutValue="(MISSING)"
I output.go:<LINE>] "test" keyWithoutValue="anotherKeyWithoutValue"
I output.go:<LINE>] "test" keyWithoutValue="(MISSING)"
`,
	}
	Output(t, OutputConfig{
		NewLogger: func(out io.Writer, v int, vmodule string) logr.Logger {
			return klogr.NewWithOptions(klogr.WithFormat(klogr.FormatKlog))
		},
		ExpectedOutputMapping: mapping,
	})
}
