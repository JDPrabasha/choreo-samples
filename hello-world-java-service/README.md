# Hello World Java

Sample for Hello World service, using plain Java (`com.sun.net.httpserver`, no framework dependencies).

### Prerequisites
1. Fork the repositoy

## Getting started

Please refer to the Choreo documentation under the [Develop an Application with Buildpacks](https://wso2.com/choreo/develop-components/deploy-an-application-with-buildpacks) to learn how to deploy the application.

1. Select `Service` Card from Component Creation Wizard
2. Select `Java` as the buildpack. Fill as follow according to selected Buildpack.

    | **Field**             | **Description**                               |
    |-----------------------|-----------------------------------------------|
    |Name           | Hello World Java Service             |
    |Description    | Hello World Java Service      |
    | **GitHub Account**    | Your account                                  |
    | **GitHub Repository** | choreo-samples |
    | **Branch**            | **`main`**                               |
    | **Buildpack**      | Java|
    | **Project Directory**       | hello-world-java-service|
    | **Select Language Version**              | 17.x |

3. Click Create. Once the component creation is complete, you will see the component overview page.
4. Deploy the created component

The entrypoint is specified in the `Procfile` (`web: java -jar target/hello-world-java-service.jar`).

## Execute the Sample Locally

Navigate to the Java application directory

```bash
cd hello-world-java-service
```

Build the project

```bash
mvn clean package
```

Run the service

```bash
PORT=8080 java -jar target/hello-world-java-service.jar
```

Invoke the endpoint

```bash
curl http://localhost:8080/hello
```
