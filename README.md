[![Go ReportCard](https://goreportcard.com/badge/github.com/drpaneas/romie)](https://goreportcard.com/report/github.com/drpaneas/romie)

[Romie](https://github.com/romie-gr/romie) is a golang project created by some
friends in a way that is a professional environment (some would say overkill)
with CI and CD systems. Its purpose is to teach us golang and to learn how we
can set up CI/CD and other workflows. The purpose of the code will be a
script/program that would allow you to download ROMs from various sites for
[Retropie](https://retropie.org.uk/). We want to have fun learning and doing a
fun project. See [here](https://github.com/drpaneas/romie/blob/master/README.md#contributors-)
for the authors and contributors.


## Developers

To replicate GH-Actions locally, use [act](https://github.com/nektos/act).

See all the targets:

```shell
$ act -l
ID       Stage  Name
build    0      Building the Project
analyze  0      Run CodeQL analysis
lint     0      Run golangci-lint
deploy   0      Update the website
test     0      Run unit tests
```

Run a target:

```shell
$ act -j <ID> # i.e act -j build
```

#### Testing

Your PR should not break the tests and also pass the linter.

Go testing:

```shell
$ go test -v ./...
```

Go testing with coverage:

```shell
$ go test -coverprofile cover.out ./...
$ go tool cover -html=cover.out
```

Linters:

```shell
act -j lint
```
