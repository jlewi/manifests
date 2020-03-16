package application

import (
	"github.com/kubeflow/manifests/tests"
	"testing"
)

func TestKustomize(t *testing.T) {
	testCase := &tests.KustomizeTestCase{
		Package: "../../../../../pipeline/pipeline-visualization-service/overlays/application",
		Expected: "test_data/expected",
	}

	tests.RunTestCase(t, testCase)
}