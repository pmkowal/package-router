##### Table of Contents

<!-- MarkdownTOC autolink="true" autoanchor="true" -->

- [Ingrid Coding Assignment](#ingrid-coding-assignment)
    - [Solution](#solution)
        - [Google App Engine](#google-app-engine)
        - [Building/deploying the application manually](#buildingdeploying-the-application-manually)
        - [Testing the application](#testing-the-application)
        - [Assumptions/decisions made](#assumptionsdecisions-made)
    - [Task Description](#task-description)

<!-- /MarkdownTOC -->


<a id="ingrid-coding-assignment"></a>
# Ingrid Coding Assignment

<a id="solution"></a>
## Solution

<a id="google-app-engine"></a>
### Google App Engine

The application has been deployed to Google App Engine and it is available at the following address: https://package-manager-311510.appspot.com/routes
Example:
https://package-manager-311510.appspot.com/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219

<a id="buildingdeploying-the-application-manually"></a>
### Building/deploying the application manually

1. The executable binary can be generated using `go build` command inside the main directory. It will build the binary which can be executed on on any machine with the same system architecture, regardless of whether the Go is installed on the system.
    ```shell
    $ go build
    ```
2. The new executable file is named after the directory in which the command was executed:
    ```shell
    $ ls
    ```
    - macOS or Linux: `packageRouter`
    - Windows: `packageRouter.exe`
3. In order to run the application the following command needs to be executed:
    - macOS or Linux:
    ```shell
    $ ./packageRouter
    ```
    - Windows:
    ```shell
    $ packageRouter.exe
    ```
4. Application will be available locally in the web browser at the following address: http://localhost:8080/. Example:
http://localhost:8080/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219

<a id="testing-the-application"></a>
### Testing the application

1. The application can be tested by running the following command:
    ```shell
    $ go test -v ./...
    ```

<a id="assumptionsdecisions-made"></a>
### Assumptions/decisions made

1. The appliaction can handle multiple pickup locations, however, a maximum of 5 concurrent API calls to the OSRM service are assumed to avoid overloading it. If there are more than 5 pickup locations queried, they are "queued" using channels mechanism. The maximum number of concurrent calls to OSRM service can be modified using `maxConcurrentRequests` constant.

<a id="task-description"></a>
## Task Description

Task description can be found in a separate file: [ingrid-backend-coding-task.md](ingrid-backend-coding-task.md)