# Data model

This document captures the Agent Factory data model and separates entities into three layers:
- Customer-facing surface
- Factory internals
- Work history/auditing layer

## Segmented data model

```mermaid
---
title: Agent Factory Data Model
config:
  layout: elk
---
erDiagram
    direction TB

    classDef customer fill:#E0F2FE,stroke:#0369A1,stroke-width:2px,color:#0C4A6E
    classDef factory fill:#ECFCCB,stroke:#4D7C0F,stroke-width:2px,color:#365314
    classDef workload fill:#FFF7ED,stroke:#C2410C,stroke-width:2px,color:#7C2D12
    classDef history fill:#EDE9FE,stroke:#6D28D9,stroke-width:2px,color:#432875

    %% Customer-facing surface
    cust["Customer"] {
        string customerId
        string name
    }
    w["Work"] {
        string workId
        string status
        datetime createdAt
    }
    rel["Relationship"] {
        string type
    }

    w }|--o{ rel: "has"
    %% Factory internals
    f["Factory"] {
        string factoryId
    }
    wt["WorkType"] {
        string name
    }
    ws["Workstation"] {
        string name
    }
    wr["Worker"] {
        string name
    }
    r["Resource"] {
        string name
    }
    f ||--o{ wt : "runs"
    f ||--o{ ws : "contains"
    f ||--o{ wr : "operates"
    f ||--o{ r : "owns"

    wt }o--|| w: "defines"
    w }o--o{ ws : "scheduled on"
    wr }o--o| ws : "assigned to"
    wr }o--o{ r: "uses"
    ws }o--o{ r: "provisions"

    %% Work request/response lane
    req["Workstation Request"] {
        string requestId
        string payload
    }
    resp["Workstation Response"] {
        string responseId
        string status
    }

    chg["Work Change"] {
    }

    chg }o--o{ w: "modifies"
    cust }o--o{ chg: "submits"
    chg }o--o{ f: "submit"

    ws ||--o{ req: "submits"
    w ||--o{ req: "produced by"
    r ||--o{ req: "enables"
    wr ||--o{ req: "executes"
    req ||--|| resp: "returns"
    resp }o--o{ w: "updates"
    

    %% Work history/audit layer
    wh["Work History"] {
        datetime startedAt
        datetime finishedAt
        string result
    }
    req }o--|| wh: "recorded in"
    resp }o--|| wh: "recorded in"
    chg }o--|| wh: "recorded in"
    class cust customer
    class f,wt,ws,wr,r factory
    class w,rel,chg,req,resp workload
    class wh history
```
