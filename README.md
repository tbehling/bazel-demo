# Bazel build demo with Go

Install bazel:

```
brew tap bazelbuild/tap
brew install bazelbuild/tap/bazel
```

### Build with Bazel

Download all Bazel external dependencies (optional)

```
bazel sync
```

Update 3rd-party Go deps for Bazel rules

```
bazel run //:gazelle -- update-repos -from_file=go.mod
```

Test Go, and run its CLI tool

```
bazel test //hello:go_default_test
bazel run //cmd/hello
```

Cross-compile Go binary for Linux:

```
bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //cmd/hello
```

Compile Go and build a Debian package (requires Python 2)

```
bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 --host_force_python=PY2 //:hello-debian
```

Compile Go and bundle it into a Docker image, and push to the local daemon;
alternately, build it based on the Debian package above.

```
bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //:hello-docker

bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //:hello-docker-deb
```

### Run Bazel under Docker

```
docker pull gcr.io/cloud-builders/bazel
docker run -w /workspace -v $(pwd):/workspace gcr.io/cloud-builders/bazel build //cmd/hello
```
 
