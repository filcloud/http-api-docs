# http-api-docs

> A generator for go-filecoin API endpoints documentation.

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Captain](#captain)
- [Contribute](#contribute)
- [License](#license)

## Install

In order to build this project, you need to first install Go, clone this repo, and finally run `make install`:

```sh
> git clone https://github.com/ipfs/http-api-docs "$(go env GOPATH)/src/github.com/ipfs/http-api-docs"
> cd "$(go env GOPATH)/src/github.com/ipfs/http-api-docs"
> make install
```

## Usage

After installing you can run:

```
> http-api-docs
```

This should spit out a Markdown document. This is exactly the `http.md` documentation at https://github.com/filcloud/http-api-docs/blob/master/http.md, so you can redirect the output to just overwrite that file.

## Captain

This project is captained by @hsanjuan.

## Contribute

PRs accepted.

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT (C) Protocol Labs, Inc.
