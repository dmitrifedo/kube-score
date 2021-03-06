package score

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProbesPodAllMissing(t *testing.T) {
	comments := testExpectedScore(t, "pod-probes-all-missing.yaml", "Pod Probes", 1)
	assert.Len(t, comments, 2)
	assert.Equal(t, "Container is missing a readinessProbe", comments[0].Summary)
	assert.Equal(t, "Container is missing a livenessProbe", comments[1].Summary)
}

func TestProbesPodMissingReady(t *testing.T) {
	// not targeted by service, expected full score
	comments := testExpectedScore(t, "pod-probes-missing-ready.yaml", "Pod Probes", 10)
	assert.Len(t, comments, 0)
}

func TestProbesPodIdenticalHTTP(t *testing.T) {
	comments := testExpectedScore(t, "pod-probes-identical-http.yaml", "Pod Probes", 7)
	assert.Len(t, comments, 1)
	assert.Equal(t, "Pod has the same readiness and liveness probe", comments[0].Summary)
}

func TestProbesPodIdenticalTCP(t *testing.T) {
	comments := testExpectedScore(t, "pod-probes-identical-tcp.yaml", "Pod Probes", 7)
	assert.Len(t, comments, 1)
	assert.Equal(t, "Pod has the same readiness and liveness probe", comments[0].Summary)
}

func TestProbesPodIdenticalExec(t *testing.T) {
	comments := testExpectedScore(t, "pod-probes-identical-exec.yaml", "Pod Probes", 7)
	assert.Len(t, comments, 1)
	assert.Equal(t, "Pod has the same readiness and liveness probe", comments[0].Summary)
}

func TestProbesTargetedByService(t *testing.T) {
	comments := testExpectedScore(t, "pod-probes-targeted-by-service.yaml", "Pod Probes", 1)
	assert.Len(t, comments, 1)
	assert.Equal(t, "Container is missing a readinessProbe", comments[0].Summary)
}

func TestProbesTargetedByServiceSameNamespace(t *testing.T) {
	comments := testExpectedScore(t, "pod-probes-targeted-by-service-same-namespace.yaml", "Pod Probes", 1)
	assert.Len(t, comments, 1)
	assert.Equal(t, "Container is missing a readinessProbe", comments[0].Summary)
}

func TestProbesTargetedByServiceDifferentNamespace(t *testing.T) {
	comments := testExpectedScore(t, "pod-probes-targeted-by-service-different-namespace.yaml", "Pod Probes", 10)
	assert.Len(t, comments, 0)
}

func TestProbesTargetedByServiceNotTargeted(t *testing.T) {
	comments := testExpectedScore(t, "pod-probes-not-targeted-by-service.yaml", "Pod Probes", 10)
	assert.Len(t, comments, 0)
}

func TestProbesMultipleContainers(t *testing.T) {
	comments := testExpectedScore(t, "pod-probes-on-different-containers.yaml", "Pod Probes", 10)
	assert.Len(t, comments, 0)
}

func TestProbesMultipleContainersInit(t *testing.T) {
	comments := testExpectedScore(t, "pod-probes-on-different-containers-init.yaml", "Pod Probes", 10)
	assert.Len(t, comments, 0)
}
