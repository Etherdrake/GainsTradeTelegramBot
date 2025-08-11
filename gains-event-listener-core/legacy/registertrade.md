# Flow for processing a registered trade 

```mermaid
flowchart TD
    A[RegisterTrade] --> B{Is it an order?};
    B -- Yes --> C[ProcessNewOrder];
    B -- No --> D[ProcessNewTrade];
```