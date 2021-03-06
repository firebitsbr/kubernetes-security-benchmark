// Copyright © 2018 Jimmi Dyson <jimmidyson@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package federated

import (
	. "github.com/onsi/ginkgo"

	"github.com/mesosphere/kubernetes-security-benchmark/pkg/framework"
	. "github.com/mesosphere/kubernetes-security-benchmark/pkg/ginkgo/matchers"
)

const controllerManagerProcessName = "federation-controller-manager"

func ControllerManager(index, subIndex int, missingProcessFunc framework.MissingProcessHandlerFunc) {
	f := framework.New(controllerManagerProcessName, missingProcessFunc)
	BeforeEach(f.BeforeEach)

	It("[3.2.1] Ensure that the --profiling argument is set to false [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithValue("--profiling", "false"))
	})
}
