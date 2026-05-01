package harness

import (
	"context"
	"fmt"
	"sync"

	"github.com/portpowered/agent-factory/pkg/interfaces"
	"github.com/portpowered/agent-factory/pkg/petri"
	"github.com/portpowered/agent-factory/pkg/workers"
)

func FirstInputToken(rawTokens any) interfaces.Token {
	switch tokens := rawTokens.(type) {
	case []any:
		if len(tokens) == 0 {
			return interfaces.Token{}
		}
		tok, ok := tokens[0].(interfaces.Token)
		if !ok {
			return interfaces.Token{}
		}
		return tok
	case []interfaces.Token:
		if len(tokens) == 0 {
			return interfaces.Token{}
		}
		return tokens[0]
	default:
		return interfaces.Token{}
	}
}

type FanoutParserExecutor struct {
	mu         sync.Mutex
	calls      int
	ChildCount int
}

func (e *FanoutParserExecutor) Execute(_ context.Context, dispatch interfaces.WorkDispatch) (interfaces.WorkResult, error) {
	e.mu.Lock()
	e.calls++
	e.mu.Unlock()

	parentWorkID := ""
	if len(dispatch.InputTokens) > 0 {
		parentWorkID = FirstInputToken(dispatch.InputTokens).Color.WorkID
	}

	spawned := make([]interfaces.TokenColor, e.ChildCount)
	for i := range spawned {
		spawned[i] = interfaces.TokenColor{
			WorkTypeID: "page",
			WorkID:     fmt.Sprintf("page-%d", i+1),
			ParentID:   parentWorkID,
		}
	}

	return interfaces.WorkResult{
		DispatchID:   dispatch.DispatchID,
		TransitionID: dispatch.TransitionID,
		Outcome:      interfaces.OutcomeAccepted,
		SpawnedWork:  spawned,
	}, nil
}

func (e *FanoutParserExecutor) CallCount() int {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.calls
}

type FailOnNthPageExecutor struct {
	mu     sync.Mutex
	calls  int
	FailOn int
}

func (e *FailOnNthPageExecutor) Execute(_ context.Context, dispatch interfaces.WorkDispatch) (interfaces.WorkResult, error) {
	e.mu.Lock()
	e.calls++
	call := e.calls
	e.mu.Unlock()

	outcome := interfaces.OutcomeAccepted
	if call == e.FailOn {
		outcome = interfaces.OutcomeFailed
	}

	return interfaces.WorkResult{
		DispatchID:   dispatch.DispatchID,
		TransitionID: dispatch.TransitionID,
		Outcome:      outcome,
	}, nil
}

type PanickingExecutor struct{}

func (e *PanickingExecutor) Execute(_ context.Context, _ interfaces.WorkDispatch) (interfaces.WorkResult, error) {
	panic("intentional executor panic for testing")
}

var _ workers.WorkerExecutor = (*PanickingExecutor)(nil)

func TokenPlaces(snap petri.MarkingSnapshot) map[string]int {
	places := make(map[string]int)
	for _, tok := range snap.Tokens {
		places[tok.PlaceID]++
	}
	return places
}
