package config

import "github.com/portpowered/agent-factory/pkg/interfaces"

func WriteExpandedFactoryLayoutForTest(
	sourceDir, targetDir string,
	cfg *interfaces.FactoryConfig,
	canonical []byte,
	sourcePath string,
) error {
	return writeExpandedFactoryLayout(sourceDir, targetDir, cfg, canonical, sourcePath)
}
