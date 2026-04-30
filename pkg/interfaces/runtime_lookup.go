package interfaces

// RuntimeWorkstationLookup resolves runtime workstation definitions by authored name.
type RuntimeWorkstationLookup interface {
	Workstation(name string) (*FactoryWorkstationConfig, bool)
}

// RuntimeDefinitionLookup resolves runtime worker and workstation definitions by authored name.
type RuntimeDefinitionLookup interface {
	RuntimeWorkstationLookup
	Worker(name string) (*WorkerConfig, bool)
}

// RuntimeConfigLookup exposes the canonical public runtime-facing lookup
// contract for consumers that need runtime definitions plus path-aware
// execution lookups.
type RuntimeConfigLookup interface {
	RuntimeDefinitionLookup
	FactoryDir() string
	RuntimeBaseDir() string
}

// FirstRuntimeDefinitionLookup returns the first non-nil runtime definition
// lookup from the provided candidates.
func FirstRuntimeDefinitionLookup(lookups ...RuntimeDefinitionLookup) RuntimeDefinitionLookup {
	for _, lookup := range lookups {
		if lookup != nil {
			return lookup
		}
	}
	return nil
}

// FirstRuntimeWorkstationLookup returns the first non-nil runtime workstation
// lookup from the provided candidates.
func FirstRuntimeWorkstationLookup(lookups ...RuntimeWorkstationLookup) RuntimeWorkstationLookup {
	for _, lookup := range lookups {
		if lookup != nil {
			return lookup
		}
	}
	return nil
}
