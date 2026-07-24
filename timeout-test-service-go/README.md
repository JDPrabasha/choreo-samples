# Timeout Test Service

A minimal Go service for verifying whether a Choreo endpoint's configured
Resiliency timeout is actually enforced at the gateway. It exposes a single
endpoint that sleeps for a configurable duration before responding, so you
can observe whether the request is cut off before or after that duration.

## Repository File Structure

| Filepath               | Description                                                                                     |
| ----------------------- | ------------------------------------------------------------------------------------------------ |
| main.go                | The Go service code. `GET /timeout-test/sleep?seconds=N` sleeps for N seconds (default 90) before responding. |
| Dockerfile             | Choreo uses the Dockerfile to build the container image of the application.                     |
| .choreo/endpoints.yaml | Choreo-specific configuration. `networkVisibility` is set to `Public` so traffic goes through the managed-API gateway rather than bypassing it. |
| openapi.yaml           | OpenAPI contract for the service, referenced by `.choreo/endpoints.yaml`.                        |

## Deploy Application

Follow the Choreo documentation under [Develop a REST API](https://wso2.com/choreo/docs/develop-components/develop-services/develop-a-rest-api/#step-1-create-a-service-component-from-a-dockerfile) to create a Service component from this directory (`timeout-test-service-go`), using the Go buildpack.

## Testing a resiliency timeout

1. Deploy the component and confirm the endpoint's network visibility is `Public` or `Organization` (not `Project`).
2. Set the endpoint's Resiliency timeout (e.g. 120000 ms) via the Console and click Apply.
3. Call the endpoint with a sleep duration comfortably above the timeout you're testing but below any known upstream default (e.g. 90s):
   ```bash
   curl -w "\ntotal time: %{time_total}s, http code: %{http_code}\n" \
     "https://<your-endpoint>/timeout-test/sleep?seconds=90"
   ```
4. Compare the observed cutoff time/status code against the configured timeout to confirm whether it's actually being enforced.

## Execute the Sample Locally

```bash
cd timeout-test-service-go
go run main.go
curl "http://localhost:9090/timeout-test/sleep?seconds=5"
```
