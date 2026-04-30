package interfaces

// FactoryState represents the current lifecycle state of a Factory.
type FactoryState string

const (
	FactoryStateIdle      FactoryState = "IDLE"
	FactoryStateRunning   FactoryState = "RUNNING"
	FactoryStatePaused    FactoryState = "PAUSED"
	FactoryStateCompleted FactoryState = "COMPLETED"
	FactoryStateFailed    FactoryState = "FAILED"
)

// RuntimeMode determines whether the runtime exits on idle completion or stays
// available for future submissions until its context is canceled.
type RuntimeMode string

const (
	RuntimeModeBatch   RuntimeMode = "BATCH"
	RuntimeModeService RuntimeMode = "SERVICE"
)

// SubmitRequest is the internal normalized item used to create work tokens.
type SubmitRequest struct {
	RequestID                string            `json:"request_id,omitempty"`
	WorkID                   string            `json:"work_id,omitempty"`
	Name                     string            `json:"name,omitempty"`
	WorkTypeID               string            `json:"work_type_name"`
	TargetState              string            `json:"target_state,omitempty"`
	CurrentChainingTraceID   string            `json:"current_chaining_trace_id,omitempty"`
	PreviousChainingTraceIDs []string          `json:"previous_chaining_trace_ids,omitempty"`
	TraceID                  string            `json:"trace_id"`
	Payload                  []byte            `json:"payload"`
	Tags                     map[string]string `json:"tags"`
	Relations                []Relation        `json:"relations"`
	ExecutionID              string            `json:"execution_id,omitempty"`
}

// WorkRequestType identifies the canonical request contract accepted by factory submit surfaces.
type WorkRequestType string

const (
	WorkRequestTypeFactoryRequestBatch WorkRequestType = "FACTORY_REQUEST_BATCH"
)

// WorkRequest is the factory-domain representation of the generated WorkRequest schema.
type WorkRequest struct {
	RequestID              string          `json:"request_id"`
	CurrentChainingTraceID string          `json:"current_chaining_trace_id,omitempty"`
	Type                   WorkRequestType `json:"type"`
	Works                  []Work          `json:"works,omitempty"`
	Relations              []WorkRelation  `json:"relations,omitempty"`
}

// WorkRequestSubmitResult describes accepted request metadata.
type WorkRequestSubmitResult struct {
	RequestID string
	TraceID   string
	Accepted  bool
}

// Work is one public item inside a WorkRequest batch.
type Work struct {
	Name                     string            `json:"name"`
	WorkID                   string            `json:"work_id,omitempty"`
	RequestID                string            `json:"request_id,omitempty"`
	WorkTypeID               string            `json:"work_type_name,omitempty"`
	State                    string            `json:"state,omitempty"`
	CurrentChainingTraceID   string            `json:"current_chaining_trace_id,omitempty"`
	PreviousChainingTraceIDs []string          `json:"previous_chaining_trace_ids,omitempty"`
	TraceID                  string            `json:"trace_id,omitempty"`
	Payload                  any               `json:"payload,omitempty"`
	Tags                     map[string]string `json:"tags,omitempty"`
	ExecutionID              string            `json:"-"`
	RuntimeRelations         []Relation        `json:"-"`
}

// WorkRelationType identifies a relationship between work items in a WorkRequest.
type WorkRelationType string

const (
	WorkRelationDependsOn   WorkRelationType = "DEPENDS_ON"
	WorkRelationParentChild WorkRelationType = "PARENT_CHILD"
)

// WorkRelation describes a relation between named work items in a WorkRequest.
type WorkRelation struct {
	Type           WorkRelationType `json:"type"`
	SourceWorkName string           `json:"source_work_name"`
	TargetWorkName string           `json:"target_work_name"`
	RequiredState  string           `json:"required_state,omitempty"`
}

// WorkRequestNormalizeOptions provides context inferred from a submit surface.
type WorkRequestNormalizeOptions struct {
	DefaultWorkTypeID string
	ValidWorkTypes    map[string]bool
	ValidStatesByType map[string]map[string]bool
}

// FactoryWorkItem describes a unit of work at a point in history.
type FactoryWorkItem struct {
	ID                       string            `json:"id"`
	WorkTypeID               string            `json:"work_type_id"`
	State                    string            `json:"state,omitempty"`
	DisplayName              string            `json:"display_name,omitempty"`
	CurrentChainingTraceID   string            `json:"current_chaining_trace_id,omitempty"`
	PreviousChainingTraceIDs []string          `json:"previous_chaining_trace_ids,omitempty"`
	TraceID                  string            `json:"trace_id,omitempty"`
	ParentID                 string            `json:"parent_id,omitempty"`
	PlaceID                  string            `json:"place_id,omitempty"`
	Tags                     map[string]string `json:"tags,omitempty"`
}

// FactoryRelation describes a typed relationship between work items.
type FactoryRelation struct {
	Type           string `json:"type"`
	SourceWorkID   string `json:"source_work_id,omitempty"`
	SourceWorkName string `json:"source_work_name,omitempty"`
	TargetWorkID   string `json:"target_work_id"`
	TargetWorkName string `json:"target_work_name,omitempty"`
	RequiredState  string `json:"required_state,omitempty"`
	RequestID      string `json:"request_id,omitempty"`
	TraceID        string `json:"trace_id,omitempty"`
}

// WorkRequestRecord stores the batch-level request observed before work token injection.
type WorkRequestRecord struct {
	RequestID       string
	Type            WorkRequestType
	TraceID         string
	Source          string
	RelationContext []WorkRelation
	ParentLineage   []string
	WorkItems       []FactoryWorkItem
	Relations       []FactoryRelation
}

// GeneratedSubmissionBatchMetadata captures request-level metadata for generated work.
type GeneratedSubmissionBatchMetadata struct {
	Source          string         `json:"source"`
	RelationContext []WorkRelation `json:"relation_context"`
	ParentLineage   []string       `json:"parent_lineage"`
}

// GeneratedSubmissionBatch carries a canonical generated request with runtime submissions.
type GeneratedSubmissionBatch struct {
	Request     WorkRequest                      `json:"request"`
	Metadata    GeneratedSubmissionBatchMetadata `json:"metadata"`
	Submissions []SubmitRequest                  `json:"submissions"`
}
