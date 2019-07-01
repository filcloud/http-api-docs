package docs

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// MarkdownFormatter implements a markdown doc generator. It is
// used to generate the IPFS website API reference at
// https://github.com/ipfs/website/blob/master/content/pages/docs/api.md
type MarkdownFormatter struct{}

func (md *MarkdownFormatter) GenerateIntro() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, `---
title: "HTTP API"
weight: 20
menu:
    reference:
        parent: api
---

<!-- TODO: Describe how to change ports and configure the API server -->
<!-- TODO: Structure this around command groups (dag, object, files, etc.) -->

<sup>Generated on %s, from go-ipfs v%s.</sup>

When an IPFS node is running as a daemon, it exposes an HTTP API that allows
you to control the node and run the same commands you can from the command
line.

In many cases, using this API this is preferable to embedding IPFS directly in
your program — it allows you to maintain peer connections that are longer
lived than your app and you can keep a single IPFS node running instead of
several if your app can be launched multiple times. In fact, the `+"`ipfs`"+`
CLI commands use this API when operating in [online mode]({{< relref
"usage.md#going-online" >}}).

This document is autogenerated from go-ipfs. For issues and support, check out
the [ipfs-http-api-docs](https://github.com/ipfs/ipfs-http-api-docs)
repository on GitHub.

## Getting started

### Alignment with CLI Commands

[Every command](../commands/) usable from the CLI is also available through
the HTTP API. For example:
`+"```sh"+
		`
> ipfs swarm peers
/ip4/104.131.131.82/tcp/4001/ipfs/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ
/ip4/104.236.151.122/tcp/4001/ipfs/QmSoLju6m7xTh3DuokvT3886QRYqxAzb1kShaanJgW36yx
/ip4/104.236.176.52/tcp/4001/ipfs/QmSoLnSGccFuZQJzRadHn95W2CrSFmZuTdDWP8HXaHca9z

> curl http://127.0.0.1:5001/api/v0/swarm/peers
{
  "Strings": [
    "/ip4/104.131.131.82/tcp/4001/ipfs/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
    "/ip4/104.236.151.122/tcp/4001/ipfs/QmSoLju6m7xTh3DuokvT3886QRYqxAzb1kShaanJgW36yx",
    "/ip4/104.236.176.52/tcp/4001/ipfs/QmSoLnSGccFuZQJzRadHn95W2CrSFmZuTdDWP8HXaHca9z",
  ]
}
`+
		"```"+`

### Arguments

Arguments are added through the special query string key "arg":

`+"```"+`
> curl "http://127.0.0.1:5001/api/v0/swarm/disconnect?arg=/ip4/54.93.113.247/tcp/48131/ipfs/QmUDS3nsBD1X4XK5Jo836fed7SErTyTuQzRqWaiQAyBYMP"
{
  "Strings": [
    "disconnect QmUDS3nsBD1X4XK5Jo836fed7SErTyTuQzRqWaiQAyBYMP success",
  ]
}
`+"```"+`

Note that it can be used multiple times to signify multiple arguments.

### Flags

Flags are added through the query string. For example, the %s
flag is the %s query parameter below:

`+"```"+`
> curl "http://127.0.0.1:5001/api/v0/object/get?arg=QmaaqrHyAQm7gALkRW8DcfGX3u8q9rWKnxEMmf7m9z515w&encoding=json"
{
  "Links": [
    {
      "Name": "index.html",
      "Hash": "QmYftndCvcEiuSZRX7njywX2AGSeHY2ASa7VryCq1mKwEw",
      "Size": 1700
    },
    {
      "Name": "static",
      "Hash": "QmdtWFiasJeh2ymW3TD2cLHYxn1ryTuWoNpwieFyJriGTS",
      "Size": 2428803
    }
  ],
  "Data": "CAE="
}
`+"```\n",
		time.Now().Format("2006-01-02"),
		FilecoinVersion(),
		"`--encoding=json`",
		"`&encoding=json`")

	return buf.String()
}

func (md *MarkdownFormatter) GenerateIndex(endps []*Endpoint) string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "## Index\n\n")

	for _, endp := range endps {
		fmt.Fprintf(buf, "  *  [%s](#%s)\n",
			strings.TrimPrefix(endp.Name, "/api/v0"),
			strings.Replace(strings.TrimPrefix(endp.Name, "/"), "/", "-", -1))
	}

	buf.WriteString("\n\n## Endpoints\n\n")
	return buf.String()
}

func (md *MarkdownFormatter) GenerateEndpointBlock(endp *Endpoint) string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, `
### %s

%s


`, endp.Name, endp.Description)
	return buf.String()
}

func (md *MarkdownFormatter) GenerateArgumentsBlock(args []*Argument, opts []*Argument) string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "#### Arguments\n\n")

	if len(args)+len(opts) == 0 {
		fmt.Fprintf(buf, "This endpoint takes no arguments.\n")
	}

	for _, arg := range args {
		fmt.Fprintf(buf, genArgument(arg, true))
	}
	for _, opt := range opts {
		fmt.Fprintf(buf, genArgument(opt, false))
	}

	fmt.Fprintf(buf, "\n")
	return buf.String()
}

func genArgument(arg *Argument, aliasToArg bool) string {
	buf := new(bytes.Buffer)
	alias := arg.Name
	if aliasToArg {
		alias = "arg"
	}
	fixDesc, _ := regexp.Compile(" Default: [a-zA-z0-9-_]+ ?\\.")
	fmt.Fprintf(buf, "  - `%s` [%s]: %s", alias, arg.Type, fixDesc.ReplaceAll([]byte(arg.Description), []byte("")))
	if len(arg.Default) > 0 {
		fmt.Fprintf(buf, ` Default: "%s".`, arg.Default)
	}
	if arg.Required {
		fmt.Fprintf(buf, ` Required: **yes**.`)
	} else {
		fmt.Fprintf(buf, ` Required: no.`)
	}
	fmt.Fprintln(buf)
	return buf.String()
}

func (md *MarkdownFormatter) GenerateBodyBlock(args []*Argument) string {
	var bodyArg *Argument
	for _, arg := range args {
		if arg.Type == "file" {
			bodyArg = arg
			break
		}
	}

	if bodyArg != nil {
		buf := new(bytes.Buffer)
		fmt.Fprintf(buf, `
#### Request Body

Argument "%s" is of file type. This endpoint expects a file in the body of the request as 'multipart/form-data'.

`, bodyArg.Name)
		return buf.String()
	}
	return ""
}

func (md *MarkdownFormatter) GenerateResponseBlock(response string) string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, `
#### Response

On success, the call to this endpoint will return with 200 and the following body:

`)

	buf.WriteString("```json\n" + response + "\n```\n\n")

	return buf.String()
}

func (md *MarkdownFormatter) GenerateExampleBlock(endp *Endpoint) string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "#### cURL Example\n\n")
	fmt.Fprintf(buf, "`")
	fmt.Fprintf(buf, "curl ")

	// Assemble arguments which are not of type file
	var queryargs []string
	hasFileArg := false
	for _, arg := range endp.Arguments {
		q := "arg="
		if arg.Type != "file" {
			q += "<" + arg.Name + ">"
			queryargs = append(queryargs, q)
		} else {
			hasFileArg = true
		}
	}

	// Assemble options
	for _, opt := range endp.Options {
		q := opt.Name + "="
		//if !opt.Required { // Omit non required options
		//	continue
		//}
		if len(opt.Default) > 0 {
			q += opt.Default
		} else {
			q += "<value>"
		}
		queryargs = append(queryargs, q)
	}

	if hasFileArg {
		fmt.Fprintf(buf, "-F file=@myfile ")
	}

	fmt.Fprintf(buf, "\"http://localhost:5001%s", endp.Name)
	if len(queryargs) > 0 {
		fmt.Fprintf(buf, "?%s\"", strings.Join(queryargs, "&"))
	} else {
		fmt.Fprintf(buf, "\"")
	}

	fmt.Fprintf(buf, "`\n\n***\n")
	return buf.String()
}
