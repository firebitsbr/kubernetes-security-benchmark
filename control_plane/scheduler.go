package control_plane

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mesosphere/kubernetes-security-benchmark/framework"
	"github.com/mesosphere/kubernetes-security-benchmark/matcher"
)

const schedulerProcessName = "kube-scheduler"

func Scheduler(index, subIndex int) {
	f := framework.New(schedulerProcessName)

	It(fmt.Sprintf("[%d.%d.1] Ensure that the --profiling argument is set to false", index, subIndex), func() {
		Expect(f.Process.CmdlineSlice()).To(matcher.HaveFlagWithValue("--profiling", "false"))
	})
}