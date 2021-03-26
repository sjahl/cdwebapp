# cdwebapp

A simple hello world application for testing a CD pipeline.

## Github Actions

The interesting parts of this repository live in the `.github/workflows` directory, which I'll try to walk through here.

### .github/workflows/test.yaml

#### Go Test matrix

  - Runs a `go test` with the specified versions of golang

#### Go Build step

  - Runs a build of the go binary after tests pass.
  - Uploads the binary as a workflow artifact that will be copied into the Docker image in the docker release step

#### Docker Release

  - Derives the 7 character short SHA of the current git commit. This is used to tag the docker image.

  - Downloads the binary artifact from the `go build` step.

  - Builds and Pushes a docker image to Docker Hub, and tags the image with the 7 character short sha of the current git commit.

#### Bump Image Tag in Helm

  - This task notifies the repository where our helm charts live, and informs a github action workflow in that repository that there is a new docker image tag.

  - Using the `peter-evans/repository-dispatch` action, specify the target repository, and the desired event type, and a payload. In our case, the action sends an event dispatch of type 'update-chart' to the helm-cdwebapp repository, and includes the new docker image tag in the client payload.

  - The receiveing repository needs two things:

    - A github action that uses the `yq` tool to manipulate yaml files: https://github.com/sjahl/helm-cdwebapp/blob/main/.github/actions/yq/action.yaml

    - A github workflow that listens to repository-dispatch events called 'update-chart', and invokes the `yq` action in response.
  
      - An example invocation can be seen here: https://github.com/sjahl/helm-cdwebapp/blob/main/.github/workflows/update_chart.yaml#L17-L22

  - The result of this that a new pull request will be opened in the helm-cdwebapp repository that updates the dockerImageTag in the values-dev.yaml file in that repository.

