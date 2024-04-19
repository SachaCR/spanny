package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var expectedDefaultConfig = SpannyConfig{
	MigrationFilesPath:   "./migrations",
	ServicePath:          "localhost",
	Port:                 9010,
	Env:                  "default",
	ProjectId:            "default-project",
	InstanceId:           "default-instance",
	DatabaseId:           "default-database",
	UsingSpannerEmulator: true,
}

func TestDefaultConfig(t *testing.T) {
	loadedConfig, err := LoadConfiguration("test", "./DirectoryDoesNotExist")

	if err != nil {
		t.Errorf("Error loading configuration: %v", err)
	}

	assert.Equal(t, expectedDefaultConfig, loadedConfig, "Should return the default config")
}

func TestLoadConfigWithTestEnv(t *testing.T) {
	loadedConfig, err := LoadConfiguration("test", "../../../")

	var expectedTestConfig = SpannyConfig{
		MigrationFilesPath:   "./migrations",
		ServicePath:          "localhost",
		Port:                 9010,
		Env:                  "test",
		ProjectId:            "test-project",
		InstanceId:           "test-instance",
		DatabaseId:           "test-database",
		UsingSpannerEmulator: true,
	}

	if err != nil {
		t.Errorf("Error loading configuration: %v", err)
	}

	assert.Equal(t, expectedTestConfig, loadedConfig, "Should return the test config")
}

func TestLoadUnknownEnv(t *testing.T) {
	t.Run("Try to load an unknown environment", func(t *testing.T) {
		loadedConfig, err := LoadConfiguration("unknown", "../../../")

		var expectedDefaultConfig = SpannyConfig{
			MigrationFilesPath:   "./migrations",
			ServicePath:          "localhost",
			Port:                 9010,
			Env:                  "default",
			ProjectId:            "local-project",
			InstanceId:           "local-instance",
			DatabaseId:           "local-database",
			UsingSpannerEmulator: true,
		}

		if err != nil {
			t.Errorf("Error loading configuration: %v", err)
		}

		assert.Equal(t, expectedDefaultConfig, loadedConfig, "Should return the default config from config file")

	})
}
