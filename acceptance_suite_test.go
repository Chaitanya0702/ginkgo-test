package acceptance_test

import (
	// "bytes"
	// "context"
	// "os"
	k8s "github.com/gruntwork-io/terratest/modules/k8s"
	"testing"

	"github.com/gruntwork-io/terratest/modules/logger"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Acceptance test Suite")
}

var _ = BeforeSuite(func() {
	logrus.Debug("Deploy the policy before the suite")
	deployJsPolicies()
})

var _ = AfterSuite(func() {
	deleteJsPolicies()
	logrus.Debug("delete the policy after the suite")
})

func deployJsPolicies() {
	logrus.Debug("Deploying jspolicies")
	k8sOptions := &k8s.KubectlOptions{Logger: logger.Discard}
	_err := k8s.KubectlApplyE(GinkgoT(), k8sOptions, "policies")
	Expect(_err).NotTo(HaveOccurred())
	logrus.Debug("Deployed jspolicies")
}

func deleteJsPolicies() {
	logrus.Debug("Deleting jspolicies")
	k8sOptions := &k8s.KubectlOptions{Logger: logger.Discard}
	_err := k8s.KubectlDeleteE(GinkgoT(), k8sOptions, "policies")
	Expect(_err).NotTo(HaveOccurred())
	logrus.Debug("Deleted jspolicies")
}
