package acceptance_test

import (
	k8s "github.com/gruntwork-io/terratest/modules/k8s"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// "github.com/sirupsen/logrus"
	"github.com/gruntwork-io/terratest/modules/logger"
)

// Given the validate-xp-mrs policy.
// Give an XR composition file using an S3 Managed bucket
// When we apply the XR
// Then XR is created in the cluster
var _ = Describe("Given the validate-xp-mrs policy \n", func() {
	var (
		xrFile string
		xrName string
	)
	Context("Give an XR composition file using an S3 Managed bucket \n ", func() {
		BeforeEach(func() {
			// The policy is applied by the suite, see `acceptance_suite_test.go`
			xrFile = "Composition-NoSQLPlus.yaml"
			xrName = "dynamo-with-bucket"
		})
		When("we apply the XR \n", func() {
			BeforeEach(func() {
				k8sOptions := &k8s.KubectlOptions{Logger: logger.Discard}
				_err := k8s.KubectlApplyE(GinkgoT(), k8sOptions, xrFile)
				Expect(_err).NotTo(HaveOccurred())
			})
			It("It should allow deployment \n", func() {
				k8sOptions := &k8s.KubectlOptions{Logger: logger.Discard}
				str_output, _err := k8s.RunKubectlAndGetOutputE(GinkgoT(), k8sOptions, "get", "composition", xrName)
				Expect(_err).NotTo(HaveOccurred())
				Expect(str_output).To(ContainSubstring(xrName))
				_err = k8s.KubectlDeleteE(GinkgoT(), k8sOptions, xrFile)
				Expect(_err).NotTo(HaveOccurred())
			})
		})
	})
})
