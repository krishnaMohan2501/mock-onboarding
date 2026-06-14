# mock-onboarding

Mock UPI registration upstream service for local development of the Shield fraud detection system.

Listens on **port 9002** and responds to `POST /upi/register` with a hardcoded success payload. Logs the `X-Risk-Score` header set by the Kong shield-check plugin on every allowed request.

## Run

```bash
go run .
# [MOCK-ONB] Starting on :9002
```

## Response

```json
{ "status": "registration_accepted", "userId": "USER-MOCK-001" }
```

## Part of

See [shield](https://github.com/krishnaMohan2501/shield) for full local setup instructions.
