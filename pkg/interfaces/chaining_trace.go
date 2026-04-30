package interfaces

import "sort"

// CanonicalChainingTraceIDs applies the shared chaining-trace fan-in rule:
// keep only non-empty predecessor chain IDs, dedupe them, and sort them.
func CanonicalChainingTraceIDs(traceIDs []string) []string {
	if len(traceIDs) == 0 {
		return nil
	}

	ordered := append([]string(nil), traceIDs...)
	sort.Strings(ordered)

	deduped := ordered[:0]
	var previous string
	for i, traceID := range ordered {
		if traceID == "" || (i > 0 && traceID == previous) {
			previous = traceID
			continue
		}
		deduped = append(deduped, traceID)
		previous = traceID
	}
	if len(deduped) == 0 {
		return nil
	}
	return deduped
}

// PreviousChainingTraceIDsFromTokens collects predecessor chain IDs from
// non-resource input tokens using the canonical fan-in ordering rule.
func PreviousChainingTraceIDsFromTokens(tokens []Token) []string {
	traceIDs := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if token.Color.DataType == DataTypeResource {
			continue
		}
		traceIDs = append(traceIDs, token.Color.TraceID)
	}
	return CanonicalChainingTraceIDs(traceIDs)
}

// PreviousChainingTraceIDsFromTokenColors collects predecessor chain IDs from
// non-resource token colors using the canonical fan-in ordering rule.
func PreviousChainingTraceIDsFromTokenColors(colors []TokenColor) []string {
	traceIDs := make([]string, 0, len(colors))
	for _, color := range colors {
		if color.DataType == DataTypeResource {
			continue
		}
		traceIDs = append(traceIDs, color.TraceID)
	}
	return CanonicalChainingTraceIDs(traceIDs)
}

// PreviousChainingTraceIDsFromWorkItems collects predecessor chain IDs from
// canonical work items using the shared deterministic fan-in rule.
func PreviousChainingTraceIDsFromWorkItems(items []FactoryWorkItem) []string {
	traceIDs := make([]string, 0, len(items))
	for _, item := range items {
		traceIDs = append(traceIDs, item.TraceID)
	}
	return CanonicalChainingTraceIDs(traceIDs)
}

// CurrentChainingTraceIDFromWorkItems resolves the current dispatch chain from
// the first non-system customer work item, falling back to any work item when
// only system work is present.
func CurrentChainingTraceIDFromWorkItems(items []FactoryWorkItem) string {
	for _, item := range items {
		if IsSystemTimeWorkType(item.WorkTypeID) {
			continue
		}
		if item.CurrentChainingTraceID != "" {
			return item.CurrentChainingTraceID
		}
		if item.TraceID != "" {
			return item.TraceID
		}
	}
	for _, item := range items {
		if item.CurrentChainingTraceID != "" {
			return item.CurrentChainingTraceID
		}
		if item.TraceID != "" {
			return item.TraceID
		}
	}
	return ""
}

// CurrentChainingTraceIDFromTokens resolves the current dispatch chain from
// the first non-resource customer work token, falling back to any non-resource
// token when only system work is present.
func CurrentChainingTraceIDFromTokens(tokens []Token) string {
	for _, token := range tokens {
		if token.Color.DataType == DataTypeResource || token.Color.WorkTypeID == SystemTimeWorkTypeID {
			continue
		}
		return token.Color.TraceID
	}
	for _, token := range tokens {
		if token.Color.DataType != DataTypeResource {
			return token.Color.TraceID
		}
	}
	return ""
}
