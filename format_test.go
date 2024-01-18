/*
Copyright 2023 The Kubernetes Authors.

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
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/pohly/plog/v2"

	"github.com/go-logr/logr"
)

func TestFormat(t *testing.T) {
	obj := config{
		TypeMeta: TypeMeta{
			Kind: "config",
		},
		RealField: 42,
	}

	assertEqual(t, "kind is config", obj.String(), "config.String()")
	assertEqual(t, `{
  "Kind": "config",
  "RealField": 42
}
`, plog.Format(obj).(fmt.Stringer).String(), "Format(config).String()")
	// fmt.Sprintf would call String if it was available.
	str := fmt.Sprintf("%s", plog.Format(obj).(logr.Marshaler).MarshalLog())
	if strings.Contains(str, "kind is config") {
		t.Errorf("fmt.Sprintf called TypeMeta.String for plog.Format(obj).MarshalLog():\n%s", str)
	}

	structured, err := json.Marshal(plog.Format(obj).(logr.Marshaler).MarshalLog())
	if err != nil {
		t.Errorf("JSON Marshal: %v", err)
	} else {
		assertEqual(t, `{"Kind":"config","RealField":42}`, string(structured), "json.Marshal(plog.Format(obj).MarshalLog())")
	}
}

func assertEqual(t *testing.T, expected, actual, what string) {
	if expected != actual {
		t.Errorf("%s:\nExpected\n%s\nActual\n%s\n", what, expected, actual)
	}
}

type TypeMeta struct {
	Kind string
}

func (t TypeMeta) String() string {
	return "kind is " + t.Kind
}

func (t TypeMeta) MarshalLog() interface{} {
	return t.Kind
}

type config struct {
	TypeMeta

	RealField int
}
