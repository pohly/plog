//go:build go1.21
// +build go1.21

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

package plog_test

import (
	"log/slog"
	"os"

	"github.com/pohly/plog/v3"
)

func ExampleSetSlogLogger() {
	state := plog.CaptureState()
	defer state.Restore()

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				// Avoid non-deterministic output.
				return slog.Attr{}
			}
			return a
		},
	})
	logger := slog.New(handler)
	plog.SetSlogLogger(logger)
	plog.Info("hello world")

	// Output:
	// level=INFO msg="hello world"
}
