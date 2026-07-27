# Timeout Test Service (WSO2 MI)

A minimal WSO2 Micro Integrator service for verifying whether a Choreo
endpoint's configured Resiliency timeout is actually enforced at the gateway.
It exposes a single API resource that sleeps for a configurable duration
before responding, so you can observe whether the request is cut off before
or after that duration.

## Repository File Structure

| Filepath                                                          | Description                                                                                     |
| ------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| timeoutTestConfigs/src/main/synapse-config/api/TimeoutTest.xml     | The Synapse API. `GET /timeout-test/sleep?seconds=N` sleeps for N seconds (default 90) before responding. |
| timeoutTestConfigs/src/main/resources/metadata                    | API metadata and OpenAPI definition used by the Micro Integrator/Choreo tooling.                 |
| pom.xml, timeoutTestConfigs, timeoutTestCompositeExporter         | Maven multi-module project structure required by the WSO2 MI buildpack.                          |

## Deploying on Choreo

1. Fork this repository.
2. Login to the [Choreo console](https://console.choreo.dev/) and create a `Service` component.
3. Select the `WSO2 MI` buildpack and `mi-timeout-test-service` as the `WSO2 MI Project Directory`.
4. Build and deploy the component, and confirm the endpoint's network visibility is `Public` or `Organization` (not `Project`).

## Testing a resiliency timeout

1. Set the endpoint's Resiliency timeout (e.g. 120000 ms) via the Console and click Apply.
2. Call the endpoint with a sleep duration comfortably above the timeout you're testing but below any known upstream default (e.g. 90s):
   ```bash
   curl -w "\ntotal time: %{time_total}s, http code: %{http_code}\n" \
     "https://<your-endpoint>/timeout-test/sleep?seconds=90"
   ```
3. Compare the observed cutoff time/status code against the configured timeout to confirm whether it's actually being enforced.
