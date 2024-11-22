### [Freshbooks](https://www.freshbooks.com/) API Client

[Bitsnap](https://bitsnap.io) Operational Intelligence Platform [Freshbooks](https://www.freshbooks.com/) API client package.

Existing accounting platforms often suffer from reliability issues, such as potential ledger data loss and inconsistent API specifications, 
making long-term use and interoperability risky. To address these challenges, Yuriy is publishing up-to-date Accounting Services API Clients with a focus on long-term support. 
These clients aim to uncover and highlight inconsistencies in API specifications and documentation.

Bitsnap focuses on importing and data processing, instead of synchronization.

## Refactoring

NOT READY FOR PRODUCTION USE.

### Bitsnap Accounting Read-only API Clients

 - [Freshbooks](https://github.com/bitsnap/freshbooks-api-client-api-client) this one
 - [Chargebee](https://github.com/bitsnap/chargebee-api-client) 
 - [Myob](https://github.com/bitsnap/myob-api-client)
 - [Paypal](https://github.com/bitsnap/paypal-api-client)
 - [QBO](https://github.com/bitsnap/qbo-api-client)
 - [Stripe](https://github.com/bitsnap/stripe-api-client)
 - [Wave](https://github.com/bitsnap/wave-api-client)
 - [Xero](https://github.com/bitsnap/xero-api-client)
 - [Zoho Books](https://github.com/bitsnap/zohobooks-api-client)

### Usage

```go
package main 

import (
  "time"
  
  "github.com/bitsnap/freshbooks-api-client-api-client"
)

// set FRESHBOOKS_API_TOKEN and FRESHBOOKS_SITE variables
//
func main() {

}

```

### Development

Install [asdf](https://asdf-vm.com/guide/getting-started.html) with all the dependencies

```bash
git clone https://github.com/asdf-vm/asdf.git ~/.asdf --branch v0.14.0
# Add the following to ~/.bashrc:
echo '. "$HOME/.asdf/asdf.sh"' >> ~/.bashrc
# Completions must be configured by adding the following to your .bashrc:
echo '. "$HOME/.asdf/completions/asdf.bash"' >> ~/.bashrc
```

```bash
# install asdf

asdf plugin add golang
asdf plugin add golangci-lint 
asdf plugin add ko
asdf plugin add goreleaser

go install golang.org/x/tools/cmd/goimports@latest
go install mvdan.cc/gofumpt@latest
go install github.com/onsi/ginkgo/v2/ginkgo@latest
```

### Code generation

```bash
# will scrape the existing API documentation pages and emit the respective API client methods
# enums and models validation, where applicable

go generate ./...
```

### Testing

```bash
FRESHBOOKS_API_TOKEN=test_xxx ginkgo run ./...
```

### TODO
 - [ ] refactor scraping
 - [ ] maybe introduce some attr filtering
 - [ ] implement webhook API endpoints for CDC streaming
 - [ ] provide [Apache Flink](https://flink.apache.org/) [JNI source](https://nightlies.apache.org/flink/flink-docs-master/docs/dev/datastream/sources/) binding  
 - [ ] provide [Argo Events](https://argoproj.github.io/argo-events/) source binding
 - [ ] benchmarks

## License

[Bitsnap](https://bitsnap.io) is a proprietary product developed by [Yuriy Yarosh](mailto:yuriy@yarosh.dev).

[Freshbooks](https://www.freshbooks.com/) golang API Client is licensed under the terms of [MIT License](LICENSE).