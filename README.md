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

## Set up

Set the CL node credentials as environment variables:

```.dotenv
CHAINLINK_URL=http://localhost:6688
CHAINLINK_ACCESS=abc # Access key generated for this external initiator
CHAINLINK_SECRET=abc # Access secret generated for this external initiator
```
