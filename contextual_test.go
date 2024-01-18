/*
Copyright 2022 The Kubernetes Authors.

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
	"context"
	"fmt"
	"runtime"
	"testing"

	"github.com/go-logr/logr"
	"github.com/pohly/plog/v2"
)

func ExampleSetLogger() {
	defer plog.ClearLogger()

	// Logger is only used as backend, Background() returns klogr.
	plog.SetLogger(logr.Discard())
	fmt.Printf("logger after SetLogger: %T\n", plog.Background().GetSink())

	// Logger is only used as backend, Background() returns klogr.
	plog.SetLoggerWithOptions(logr.Discard(), plog.ContextualLogger(false))
	fmt.Printf("logger after SetLoggerWithOptions with ContextualLogger(false): %T\n", plog.Background().GetSink())

	// Logger is used as backend and directly.
	plog.SetLoggerWithOptions(logr.Discard(), plog.ContextualLogger(true))
	fmt.Printf("logger after SetLoggerWithOptions with ContextualLogger(true): %T\n", plog.Background().GetSink())

	// Output:
	// logger after SetLogger: *plog.klogger
	// logger after SetLoggerWithOptions with ContextualLogger(false): *plog.klogger
	// logger after SetLoggerWithOptions with ContextualLogger(true): <nil>
}

func ExampleFlushLogger() {
	defer plog.ClearLogger()

	// This simple logger doesn't need flushing, but others might.
	plog.SetLoggerWithOptions(logr.Discard(), plog.FlushLogger(func() {
		fmt.Print("flushing...")
	}))
	plog.Flush()

	// Output:
	// flushing...
}

func BenchmarkPassingLogger(b *testing.B) {
	b.Run("with context", func(b *testing.B) {
		ctx := plog.NewContext(context.Background(), plog.Background())
		var finalCtx context.Context
		for n := b.N; n > 0; n-- {
			finalCtx = passCtx(ctx)
		}
		runtime.KeepAlive(finalCtx)
	})

	b.Run("without context", func(b *testing.B) {
		logger := plog.Background()
		var finalLogger plog.Logger
		for n := b.N; n > 0; n-- {
			finalLogger = passLogger(logger)
		}
		runtime.KeepAlive(finalLogger)
	})
}

func BenchmarkExtractLogger(b *testing.B) {
	b.Run("from context", func(b *testing.B) {
		ctx := plog.NewContext(context.Background(), plog.Background())
		var finalLogger plog.Logger
		for n := b.N; n > 0; n-- {
			finalLogger = extractCtx(ctx)
		}
		runtime.KeepAlive(finalLogger)
	})
}

//go:noinline
func passCtx(ctx context.Context) context.Context { return ctx }

//go:noinline
func extractCtx(ctx context.Context) plog.Logger { return plog.FromContext(ctx) }

//go:noinline
func passLogger(logger plog.Logger) plog.Logger { return logger }
