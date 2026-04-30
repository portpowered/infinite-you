package workers

import (
	"sort"
	"strings"
)

type commandEnvEntry struct {
	name  string
	value string
}

func commandEnvEntriesFromMap(envVars map[string]string) []commandEnvEntry {
	if len(envVars) == 0 {
		return nil
	}

	keys := make([]string, 0, len(envVars))
	for name := range envVars {
		keys = append(keys, name)
	}
	sort.Strings(keys)

	entries := make([]commandEnvEntry, 0, len(keys))
	for _, name := range keys {
		entries = append(entries, commandEnvEntry{name: name, value: envVars[name]})
	}
	return entries
}

func mergeCommandEnv(base []string, overlays ...[]commandEnvEntry) []string {
	values := make(map[string]string)
	order := make([]string, 0, len(base))
	setEnv := func(name, value string) {
		if name == "" {
			return
		}
		if _, exists := values[name]; !exists {
			order = append(order, name)
		}
		values[name] = value
	}

	for _, entry := range base {
		name, value, ok := strings.Cut(entry, "=")
		if !ok {
			continue
		}
		setEnv(name, value)
	}

	for _, overlay := range overlays {
		for _, entry := range overlay {
			setEnv(entry.name, entry.value)
		}
	}

	env := make([]string, 0, len(order))
	for _, name := range order {
		env = append(env, name+"="+values[name])
	}
	return env
}
