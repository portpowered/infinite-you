{
  "id": "US-001",
  "title": "Add health check endpoint",
  "description": "As a developer, I want a /healthz endpoint that returns 200 OK so that load balancers can verify the service is running.",
  "acceptanceCriteria": [
    "GET /healthz returns HTTP 200 with body {\"status\":\"ok\"}",
    "Endpoint is registered in the router",
    "Unit test verifies the response",
    "Typecheck passes"
  ]
}
