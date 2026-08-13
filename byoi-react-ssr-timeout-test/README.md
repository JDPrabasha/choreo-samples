# BYOI React SSR Timeout Test

A minimal server-side rendered (SSR) React web app, packaged for **Bring Your
Own Image (BYOI)** deployment as a Choreo Web Application. It reproduces a
customer-reported issue: a BYOI Web Application that generates reports taking
longer than 60 seconds has its requests terminated around the 60-second mark,
and the endpoint timeout needs to be raised to ~300 seconds (5 minutes) for
that component only — not platform-wide.

The page renders a "Generate Report" button. Clicking it calls
`GET /api/generate-report`, which sleeps for a configurable number of seconds
(default 90) before responding, simulating a slow report-generation backend.
This lets you observe whether the request is cut off before or after the
configured Resiliency timeout.

## Repository File Structure

| Filepath                | Description                                                                                       |
| ------------------------ | --------------------------------------------------------------------------------------------------- |
| server.js               | Express server. `GET /` renders the page via `react-dom/server`. `GET /api/generate-report?seconds=N` sleeps for N seconds (default 90) before responding. |
| public/client.js        | Browser script that wires up the "Generate Report" button and calls the API.                      |
| Dockerfile              | Builds the container image to push to your registry for BYOI deployment.                          |
| .choreo/endpoints.yaml  | Endpoint configuration to use when creating the BYOI component in Choreo.                         |

## Execute the Sample Locally

```bash
cd byoi-react-ssr-timeout-test
npm install
npm start
```

Open http://localhost:8080 and click **Generate Report**, or call the API directly:

```bash
curl -w "\ntotal time: %{time_total}s\n" "http://localhost:8080/api/generate-report?seconds=5"
```

## Build and Push the Container Image

The image is published to GitHub Container Registry:

```
ghcr.io/jdprabasha/byoi-react-ssr-timeout-test:latest
```

To rebuild and push your own copy:

```bash
cd byoi-react-ssr-timeout-test
docker build -t ghcr.io/<your-gh-username>/byoi-react-ssr-timeout-test:latest .
docker push ghcr.io/<your-gh-username>/byoi-react-ssr-timeout-test:latest
```

## Deploy as a BYOI Web Application in Choreo

1. In your project, select **Create Component** with the **Web Application** type.
2. Provide a component name and description.
3. Select **Deploy an image from a Container Registry** as the source, and provide `ghcr.io/jdprabasha/byoi-react-ssr-timeout-test:latest` (or your own image from the previous step).
4. Configure the endpoint using [.choreo/endpoints.yaml](.choreo/endpoints.yaml) (port `8080`, context `/`).
5. Deploy the component.

## Reproducing the Timeout Cutoff and Verifying the Fix

1. With the component freshly deployed (default Resiliency timeout, ~60s), click **Generate Report** or call:
   ```bash
   curl -w "\ntotal time: %{time_total}s, http code: %{http_code}\n" \
     "https://<your-endpoint>/api/generate-report?seconds=90"
   ```
   Confirm the request is terminated around the 60-second mark rather than completing.
2. In the Choreo Console, open this component's endpoint settings and increase the **Resiliency timeout** to `300000` ms (5 minutes), then click Apply. This applies to this component's endpoint only, not the platform default.
3. Repeat the request from step 1 (or click the button again) and confirm it now completes successfully, returning the "Report generated after ~90.0s" message instead of being cut off.
