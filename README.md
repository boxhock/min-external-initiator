# Minimum Chainlink External Initiator

This simple program will act as an external initiator for a Chainlink node. It will trigger a job if the URL provided
returns a new body.

This implementation has not actually been tested against a Chainlink node, nor does it do necessary input validation. It
only serves as a quick introduction into how External Initiators work.

### Example params in job spec

```json
{
  "name": "min-ei",
  "params": {
    "url": "https://ethgasstation.info/json/ethgasAPI.json",
    "method": "GET"
  }
}
```
