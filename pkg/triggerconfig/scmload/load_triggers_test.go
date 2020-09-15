package scmload_test

import (
	"testing"

	"github.com/jenkins-x/go-scm/scm/driver/fake"
	"github.com/jenkins-x/lighthouse/pkg/config"
	"github.com/jenkins-x/lighthouse/pkg/plugins"
	"github.com/jenkins-x/lighthouse/pkg/triggerconfig/scmload"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMergeConfig(t *testing.T) {
	scmClient, _ := fake.NewDefault()
	owner := "myorg"
	repo := "myrepo"
	ref := "master"
	cfg := &config.Config{}
	pluginCfg := &plugins.Configuration{}
	flag, err := scmload.MergeTriggers(cfg, pluginCfg, scmClient, owner, repo, ref)
	require.NoError(t, err, "failed to merge configs")
	assert.True(t, flag, "did not return merge flag")

	LogConfig(t, cfg)

	r := owner + "/" + repo
	assert.Len(t, cfg.Presubmits[r], 2, "presubmits for repo %s", r)
	assert.Len(t, cfg.Postsubmits[r], 1, "postsubmits for repo %s", r)
}

func TestInvalidConfigs(t *testing.T) {
	scmClient, _ := fake.NewDefault()

	invalidRepos := []string{"duplicate-presubmit", "duplicate-postsubmit"}
	for _, repo := range invalidRepos {
		owner := "myorg"
		ref := "master"
		_, err := scmload.LoadTriggerConfig(scmClient, owner, repo, ref)
		require.Errorf(t, err, "should have failed to load triggers from repo %s/%s with ref %s", owner, repo, ref)

		t.Logf("got expected error loading invalid configuration on repo %s of: %s", repo, err.Error())
	}
}
