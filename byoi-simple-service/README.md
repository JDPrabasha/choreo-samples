# BYOI Simple Service

A minimal Go HTTP service, packaged for **Bring Your Own Image (BYOI)**
deployment as a Choreo Service. It exposes a single `GET /hello` endpoint
that returns a static JSON response — useful as a baseline for verifying
BYOI Service deployments in Choreo.

## Repository File Structure

| Filepath                | Description                                                                       |
| ------------------------ | ---------------------------------------------------------------------------------- |
| main.go                 | Go service. `GET /hello` returns `{"message": "hello world"}`.                    |
| Dockerfile              | Builds the container image to push to your registry for BYOI deployment.          |
| .choreo/endpoints.yaml  | Endpoint configuration to use when creating the BYOI component in Choreo.         |
| openapi.yaml            | OpenAPI contract for `/hello`, referenced by `.choreo/endpoints.yaml` so it shows up in the Choreo Test Console. |

## Execute the Sample Locally

```bash
cd byoi-simple-service
go run main.go
curl "http://localhost:8080/hello"
```

## Build and Push the Container Image

The image is published to GitHub Container Registry:

```
ghcr.io/jdprabasha/byoi-simple-service:latest
```

To rebuild and push your own copy:

```bash
cd byoi-simple-service
docker build -t ghcr.io/<your-gh-username>/byoi-simple-service:latest .
docker push ghcr.io/<your-gh-username>/byoi-simple-service:latest
```

## Deploy as a BYOI Service in Choreo

1. In your project, select **Create Component** with the **Service** type.
2. Provide a component name and description.
3. Select **Deploy an image from a Container Registry** as the source, and provide `ghcr.io/jdprabasha/byoi-simple-service:latest` (or your own image from the previous step).
4. Configure the endpoint using [.choreo/endpoints.yaml](.choreo/endpoints.yaml) (port `8080`, context `/`).
5. Deploy the component.
