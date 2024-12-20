# go-sonarcloud

An HTTP client to interact with the [SonarCloud](https://sonarcloud.io) API from Go. It uses a service structure
similar to that of [go-github](https://github.com/google/go-github).

Most code in the [`sonarcloud`](sonarcloud) folder is generated by the code in [`gen`](gen), using [`gen/webservices.json`](gen/webservices.json)
as source of truth for the API endpoints. This is one of the first iterations of generating the API. It only supports
JSON return types and skips some endpoints that need special handling of the return values.

See https://pkg.go.dev/github.com/m-yosefpor/go-sonarcloud/sonarcloud for supported endpoints.

## Installation

```shell
go get github.com/m-yosefpor/go-sonarcloud
```

## Usage

Use `sonarcloud.NewClient` to create a new client. It needs a SonarCloud organization and token, and optionally accepts
an existing `*http.Client`.

After creating the client, create a new request from one of the endpoint-specific packages in `sonarcloud/`, i.e.
`projects.SearchRequest` from `github.com/m-yosefpor/go-sonarcloud/sonarcloud/projects`. If an endpoint supports paging,
you need to supply `paging.PagingParams` as well, though in that case it might be easier to use the `All` variant of the
function. The example below, for instance, uses `client.Projects.SearchAll(req)` instead of `client.Projects.Search(req, paging)`.

The page size for automatic paging is currently set to `100`, but might be configurable later.

```go
func main() {
    org, ok := os.LookupEnv("SONARCLOUD_ORG")
    if !ok {
        log.Fatalf("missing SONARCLOUD_ORG environment variable")
    }

    token, ok := os.LookupEnv("SONARCLOUD_TOKEN")
    if !ok {
        log.Fatalf("mising SONARCLOUD_TOKEN environment variable")
    }

    client := sonarcloud.NewClient(org, token, nil)
    req := projects.SearchRequest{}

    res, err := client.Projects.SearchAll(req)
    if err != nil {
        log.Fatalf("could not search projects: %+v", err)
    }

    fmt.Printf("%+v\n", res)
}
```

## Re-generating the API

Use `make gen` to re-generate the API from the [`gen/webservices.json`](gen/webservices.json) file. This file is sourced
from the [webservices/list](https://sonarcloud.io/api/webservices/list) endpoint.
