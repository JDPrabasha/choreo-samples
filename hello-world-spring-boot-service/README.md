# Hello World Spring Boot

Sample for Hello World service, using Spring Boot (`spring-boot-starter-web`).

### Prerequisites
1. Fork the repositoy

## Getting started

Please refer to the Choreo documentation under the [Develop an Application with Buildpacks](https://wso2.com/choreo/develop-components/deploy-an-application-with-buildpacks) to learn how to deploy the application.

1. Select `Service` Card from Component Creation Wizard
2. Select `Java` as the buildpack. Fill as follow according to selected Buildpack.

    | **Field**             | **Description**                               |
    |-----------------------|-----------------------------------------------|
    |Name           | Hello World Spring Boot Service             |
    |Description    | Hello World Spring Boot Service      |
    | **GitHub Account**    | Your account                                  |
    | **GitHub Repository** | choreo-samples |
    | **Branch**            | **`main`**                               |
    | **Buildpack**      | Java|
    | **Project Directory**       | hello-world-spring-boot-service|
    | **Select Language Version**              | 17.x |

3. Click Create. Once the component creation is complete, you will see the component overview page.
4. Deploy the created component

No `Procfile` is needed — Choreo's Java buildpack detects the Spring Boot executable jar (built via `spring-boot-maven-plugin`) and runs it automatically. The service reads the port to bind from the `PORT` environment variable Choreo injects (see `src/main/resources/application.properties`).

## Execute the Sample Locally

Navigate to the Spring Boot application directory

```bash
cd hello-world-spring-boot-service
```

Build the project

```bash
mvn clean package
```

Run the service

```bash
PORT=8080 java -jar target/hello-world-spring-boot-service-1.0.0.jar
```

Invoke the endpoint

```bash
curl http://localhost:8080/hello
```
