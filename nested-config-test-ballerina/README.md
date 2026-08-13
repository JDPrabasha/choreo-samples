# Nested Config Test Service

A minimal Ballerina service for verifying that `choreo describe component-config`
correctly enumerates deeply nested Ballerina configurables — specifically
arrays of primitives, arrays of records, a record nested inside a record, and
an array of records nested inside a record that is itself nested inside
another record.

## Configurables

| Configurable                    | Shape                                         |
| -------------------------------- | ---------------------------------------------- |
| `destinations`                   | array of strings (top-level)                   |
| `profile.roles`                  | array of strings, nested one level             |
| `profile.address`                | record nested inside a record                  |
| `profile.address.tags`           | array of strings, nested two levels            |
| `profile.address.contacts`       | array of records, nested two levels            |

## Repository File Structure

| Filepath               | Description                                                                                     |
| ----------------------- | ------------------------------------------------------------------------------------------------ |
| service.bal             | The Ballerina service code, declaring the nested configurables above and exposing them via two GET endpoints. |
| Ballerina.toml          | Ballerina package descriptor, used by Choreo's Ballerina buildpack to build the component.       |
| .choreo/component.yaml  | Choreo-specific configuration. Declares `Project` and `Public` as allowed network visibilities for the endpoint. |
| openapi.yaml            | OpenAPI contract for the service, referenced by `.choreo/component.yaml`.                        |

## Deploy Application

Follow the Choreo documentation under [Develop a REST API](https://wso2.com/choreo/docs/develop-components/develop-services/develop-a-rest-api/#step-1-create-a-service-component-from-a-dockerfile) to create a Service component from this directory (`nested-config-test-ballerina`), using the Ballerina buildpack.

## Execute the Sample Locally

```bash
cd nested-config-test-ballerina
bal run
curl "http://localhost:9090/nested-config-test/destinations"
curl "http://localhost:9090/nested-config-test/profile"
```
