package acceptance_test

import (
	"bytes"
	"os/exec"
	// "time"
	// "context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)
// Given the validate-xp-mrs policy. 
// Give an XR with an S3 Managed bucket
// When we apply the XR 
// Then XR is created in the cluster
var _ = Describe("Given the validate-xp-mrs policy", func() {
	var (xr string)	
	Context("Give an XR with an S3 Managed bucket", func(){ 
		BeforeEach(func() {
			// policies is applied by the suite
			xr = "Composition-NoSQLPlus.yaml"
		})
	})
	When("When we apply the XR ", func() {
		BeforeEach(func() {

			cmd := exec.Command("kubectl ", "apply ", "-f ", xr)
			var stdout, stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr
			err := cmd.Run()
			if err == nil {logrus.Info("No Error")}
			Expect(stderr.String()).To(BeEmpty())
			// err = addItemInDynamodbTable(SentinelImageMetadataTable, imageMetadataAttributeValues)
			// Expect(err).NotTo(HaveOccurred())
		})
		It("should allow deployment", func() {
			// filePath := "deployments/common-policies/sps-kabini-acceptance-7.yaml"
			// logrus.Info("Apply delpoyment file ", filePath)
			cmd := exec.Command("kubectl", "get", "composition", "dynamo-with-bucket")
			// time.Sleep(time.Second)
			var stdout, stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr
			err := cmd.Run()
			if err == nil {logrus.Info("No Error")}
			Expect(stderr.String()).To(BeEmpty())
			// cmd = exec.Command("kubectl", "delete", "-f", filePath)

		})
	})
})