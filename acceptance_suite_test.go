package acceptance_test

import (
	// "bytes"
	// "context"
	// "os"
	"os/exec"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)


func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Acceptance test Suite")
}

var _ = BeforeSuite(func() {
	// err := 0
	// if err != nil {
	// 	Fail("Error in getting access token: " + err.Error())
	// }
	deployJsPolicies()
})

var _ = AfterSuite(func() {
	// deleteJsPolicies()
	logrus.Info("Nothing to be done")
})

func deployJsPolicies() {
	logrus.Info("Deploying jspolicies")
	cmd := exec.Command("kubectl", "apply", "-f", "policies")
	time.Sleep(time.Second)
	err := cmd.Run()
	if err != nil {
		Fail("Error in applying in policies: " + err.Error())
	}
	logrus.Info("Deployed jspolicies")
}

func deleteJsPolicies() {
	logrus.Info("Deleting jspolicies")
	cmd := exec.Command("kubectl", "delete", "-f", "policies")
	time.Sleep(time.Second)
	err := cmd.Run()
	if err != nil {
		Fail("Error in deleting in policies: " + err.Error())
	}
	logrus.Info("Deleted jspolicies")
}
