---
type: MODEL_WORKSTATION
limits:
  maxExecutionTime: 30m
---


You are processing work item {{ (index .Inputs 0).WorkID }} of type {{ (index .Inputs 0).WorkTypeID }}.

The customer is asking you to convert the following ask into a prd using the /prd and /ralph skills. 

Please convert the file into the corresponding tasks/todo/{{ (index .Inputs 0).Name }}.json.

Note that you are working in autonomous mode, do not ask any questions to the customer.

When you are done, respond with exactly: "<COMPLETE>".

The customer ask is as follows: 

{{ (index .Inputs 0).Payload }}
