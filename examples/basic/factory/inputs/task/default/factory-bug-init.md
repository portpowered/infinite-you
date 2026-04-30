# problem statement

We are having a problem with the library/agent-factory. 

When the agent-factory inits a directory, the inited directory is incorrect. What we want the default factory to look like, looks more like the current factory at the {rootdir}/factory directory. 

# Solution

We should update the init factory to roughly do what is in the current root factory dir. 

no sub factory directory. 
there is proper definitions for AGENTS.md for the workers/worstation

## Acceptance criteria

We have a functional test that roughly does: 
- runs the init command on a temporary directory
- runs the generated init command factory as the root directory for a functional test. 
- validates that this works e2e. 
- note that we are transitioning our tests to be more modelled in an async fashion, please use the dispatcher_lifecycle_test as the template for writing new tests. 

The init factory command for the agent-factory CLI generates the appropriate structures necessary for a basic executor that consumes a piece of work, and does it. 
