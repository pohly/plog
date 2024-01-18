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

package testdata

import (
	"github.com/pohly/plog/v3"
)

func calls() {
	plog.Infof("%s") // want `github.com/pohly/plog/v3.Infof format %s reads arg #1, but call has 0 args`
	plog.Infof("%s", "world")
	plog.Info("%s", "world") // want `github.com/pohly/plog/v3.Info call has possible formatting directive %s`
	plog.Info("world")
	plog.Infoln("%s", "world") // want `github.com/pohly/plog/v3.Infoln call has possible formatting directive %s`
	plog.Infoln("world")

	plog.InfofDepth(1, "%s") // want `github.com/pohly/plog/v3.InfofDepth format %s reads arg #1, but call has 0 args`
	plog.InfofDepth(1, "%s", "world")
	plog.InfoDepth(1, "%s", "world") // want `github.com/pohly/plog/v3.InfoDepth call has possible formatting directive %s`
	plog.InfoDepth(1, "world")
	plog.InfolnDepth(1, "%s", "world") // want `github.com/pohly/plog/v3.InfolnDepth call has possible formatting directive %s`
	plog.InfolnDepth(1, "world")

	plog.Warningf("%s") // want `github.com/pohly/plog/v3.Warningf format %s reads arg #1, but call has 0 args`
	plog.Warningf("%s", "world")
	plog.Warning("%s", "world") // want `github.com/pohly/plog/v3.Warning call has possible formatting directive %s`
	plog.Warning("world")
	plog.Warningln("%s", "world") // want `github.com/pohly/plog/v3.Warningln call has possible formatting directive %s`
	plog.Warningln("world")

	plog.WarningfDepth(1, "%s") // want `github.com/pohly/plog/v3.WarningfDepth format %s reads arg #1, but call has 0 args`
	plog.WarningfDepth(1, "%s", "world")
	plog.WarningDepth(1, "%s", "world") // want `github.com/pohly/plog/v3.WarningDepth call has possible formatting directive %s`
	plog.WarningDepth(1, "world")
	plog.WarninglnDepth(1, "%s", "world") // want `github.com/pohly/plog/v3.WarninglnDepth call has possible formatting directive %s`
	plog.WarninglnDepth(1, "world")

	plog.Errorf("%s") // want `github.com/pohly/plog/v3.Errorf format %s reads arg #1, but call has 0 args`
	plog.Errorf("%s", "world")
	plog.Error("%s", "world") // want `github.com/pohly/plog/v3.Error call has possible formatting directive %s`
	plog.Error("world")
	plog.Errorln("%s", "world") // want `github.com/pohly/plog/v3.Errorln call has possible formatting directive %s`
	plog.Errorln("world")

	plog.ErrorfDepth(1, "%s") // want `github.com/pohly/plog/v3.ErrorfDepth format %s reads arg #1, but call has 0 args`
	plog.ErrorfDepth(1, "%s", "world")
	plog.ErrorDepth(1, "%s", "world") // want `github.com/pohly/plog/v3.ErrorDepth call has possible formatting directive %s`
	plog.ErrorDepth(1, "world")
	plog.ErrorlnDepth(1, "%s", "world") // want `github.com/pohly/plog/v3.ErrorlnDepth call has possible formatting directive %s`
	plog.ErrorlnDepth(1, "world")

	plog.Fatalf("%s") // want `github.com/pohly/plog/v3.Fatalf format %s reads arg #1, but call has 0 args`
	plog.Fatalf("%s", "world")
	plog.Fatal("%s", "world") // want `github.com/pohly/plog/v3.Fatal call has possible formatting directive %s`
	plog.Fatal("world")
	plog.Fatalln("%s", "world") // want `github.com/pohly/plog/v3.Fatalln call has possible formatting directive %s`
	plog.Fatalln("world")

	plog.FatalfDepth(1, "%s") // want `github.com/pohly/plog/v3.FatalfDepth format %s reads arg #1, but call has 0 args`
	plog.FatalfDepth(1, "%s", "world")
	plog.FatalDepth(1, "%s", "world") // want `github.com/pohly/plog/v3.FatalDepth call has possible formatting directive %s`
	plog.FatalDepth(1, "world")
	plog.FatallnDepth(1, "%s", "world") // want `github.com/pohly/plog/v3.FatallnDepth call has possible formatting directive %s`
	plog.FatallnDepth(1, "world")

	plog.V(1).Infof("%s") // want `\(github.com/pohly/plog/v3.Verbose\).Infof format %s reads arg #1, but call has 0 args`
	plog.V(1).Infof("%s", "world")
	plog.V(1).Info("%s", "world") // want `\(github.com/pohly/plog/v3.Verbose\).Info call has possible formatting directive %s`
	plog.V(1).Info("world")
	plog.V(1).Infoln("%s", "world") // want `\(github.com/pohly/plog/v3.Verbose\).Infoln call has possible formatting directive %s`
	plog.V(1).Infoln("world")

	plog.V(1).InfofDepth(1, "%s") // want `\(github.com/pohly/plog/v3.Verbose\).InfofDepth format %s reads arg #1, but call has 0 args`
	plog.V(1).InfofDepth(1, "%s", "world")
	plog.V(1).InfoDepth(1, "%s", "world") // want `\(github.com/pohly/plog/v3.Verbose\).InfoDepth call has possible formatting directive %s`
	plog.V(1).InfoDepth(1, "world")
	plog.V(1).InfolnDepth(1, "%s", "world") // want `\(github.com/pohly/plog/v3.Verbose\).InfolnDepth call has possible formatting directive %s`
	plog.V(1).InfolnDepth(1, "world")

	// Detecting format specifiers for plog.InfoS and other structured logging calls would be nice,
	// but doesn't work the same way because of the extra "msg" string parameter. logcheck
	// can be used instead of "go vet".
	plog.InfoS("%s", "world")
}
