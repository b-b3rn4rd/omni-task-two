# Test 2
## Description
The following test will require you to do the following:
* Create a simple application which has a single "/healthcheck" endpoint.
* Containerise your application as a single deployable artifact, encapsulating all dependencies. Create a CI pipeline for your application
* The application can be written in any programming language. We'd recommend using one the following: NodeJS or GoLang. Please indicate your preferred programming language.
* The application should be a simple, small, operable web-style API or service provider. 
* It should implement the following:
* * An endpoint which returns basic information about your application in JSON format which is generated; 
* * The following is expected: Applications Version.
Description. ("static variable") Last Commit SHA.

## Implementation
### Microservice
It's a golang restful microservice using [go-restful](github.com/emicklei/go-restful), the same library that's being used by kubernetes for it's internal apis.

The microservice comes with a single resource that provides two endpoints
`/healtcheck` and `/` running on port `8080`

There is a simple unit test to check that the index page is returning `hello world`, additionally
I'm also performing integration test using additional docker container.

For the dependency management I'm using `go dep`, for linting `gometalinter`.

The build specific variables for the `/healthcheck` page are passed using `ldflags`

### Docker
The application is containerised using multistage build process.
The build stage is installing dependencies, performs linting and builds the binary.

The second stage copies binary from builder stage and runs it as unprivileged user

I'm also using `docker-compose` to orchestrate image build and publish processes.


### CI
The aim was to provide a platform agnostic CI. The inspiration comes from [3musketeers](https://3musketeers.io/) approach, which is based on
gmake, docker-compose and dockerfile.

During a container test process I'm also running an additional container to perform
integration test and in theory it should run automation tests. The idea was taken from [docker.com](https://success.docker.com/article/dev-pipeline#cicdworkflow) and
[this article](https://docs.microsoft.com/en-us/azure/devops/pipelines/languages/docker?view=azure-devops&tabs=yaml#use-docker-compose)


In order to organise pipeline workflow, I'm tagging docker images differently depending on a branch.
* If a branch is `master` then the tag format is `major.manor.patch` where `patch` is a build id
* For all other brancges the tag format is `major.manor.patch-gitsha1` where `patch` is a build id and `gitsha1` is current commit sha1

**Dockerhub link: https://cloud.docker.com/repository/registry-1.docker.io/bernard/omni-task-two**

### Limitations
Obviously there's a room to extend CI process with additional tools and processes such as:

* sonarqube - to store test coverage reports and linting scores
* integrate sonarqube with github or bitbucket for PRs
* use artifactoy for artifact management
* use mockery for tests
* use goreleaser during publishing
  