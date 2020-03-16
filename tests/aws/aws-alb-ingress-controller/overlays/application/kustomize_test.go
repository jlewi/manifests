package application

import (
	"github.com/kubeflow/manifests/tests"
	"testing"
)

func TestKustomize(t *testing.T) {
	testCase := &tests.KustomizeTestCase{
		Package: "../../../../../aws/aws-alb-ingress-controller/overlays/application",
		Expected: "test_data/expected",
	}

	tests.RunTestCase(t, testCase)
}