> # ♻️ retry
>
> Functional mechanism to perform actions repetitively until successful.

[![Awesome][icon_awesome]][awesome]
[![Patreon][icon_patreon]][support]
[![GoDoc][icon_docs]][docs]
[![Research][icon_research]][research]
[![License][icon_license]][license]

## Important news

The **[master][legacy]** is a feature frozen branch for versions **3.3.x** and no longer maintained.

```bash
$ dep ensure -add github.com/kamilsk/retry@3.3.3
```

The **[v3][]** branch is a continuation of the **[master][legacy]** branch for versions **v3.4.x**
to better integration with [Go Modules][gomod].

```bash
$ go get -u github.com/kamilsk/retry/v3@v3.4.4
```

The **[v4][]** branch is an actual development branch.

```bash
$ go get -u github.com/kamilsk/retry/v4

$ dep ensure -add github.com/kamilsk/retry@v4.0.0-rc5
```

Version **v4.x.y** focused on integration with the 🚧 [breaker][] and the 🧰 [platform][] packages.

## Usage

### Quick start

#### retry.Retry

```go
var response *http.Response

action := func(uint) error {
	var err error
	response, err = http.Get("https://github.com/kamilsk/retry")
	return err
}

if err := retry.Retry(breaker.BreakByTimeout(time.Minute), action, strategy.Limit(3)); err != nil {
	// handle error
}
// work with response
```

#### retry.Try

```go
var response *http.Response

action := func(uint) error {
	var err error
	response, err = http.Get("https://github.com/kamilsk/retry")
	return err
}
interrupter := breaker.MultiplexTwo(
	breaker.BreakByTimeout(time.Minute),
	breaker.BreakBySignal(os.Interrupt),
)
defer interrupter.Close()

if err := retry.Try(interrupter, action, strategy.Limit(3)); err != nil {
	// handle error
}
// work with response
```

Or use Context

```go
ctx, cancel := context.WithTimeout(request.Context(), time.Minute)
defer cancel()

if err := retry.Try(ctx, action, strategy.Limit(3)); err != nil {
	// handle error
}
// work with response
```

#### retry.TryContext

```go
var response *http.Response

action := func(ctx context.Context, _ uint) error {
	req, err := http.NewRequest(http.MethodGet, "https://github.com/kamilsk/retry", nil)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	response, err = http.DefaultClient.Do(req)
	return err
}
ctx, cancel := context.WithTimeout(request.Context(), time.Minute)
br, ctx := breaker.WithContext(ctx)
defer func() {
	// they do the same thing
	br.Close()
	close()
}()

if err := retry.TryContext(ctx, action, strategy.Limit(3)); err != nil {
	// handle error
}
// work with response
```

### Console tool for command execution with retries

This example shows how to repeat console command until successful.

```bash
$ retry --infinite -timeout 10m -backoff=lin:500ms -- /bin/sh -c 'echo "trying..."; exit $((1 + RANDOM % 10 > 5))'
```

[![asciicast](https://asciinema.org/a/150367.png)](https://asciinema.org/a/150367)

See more details [here][cli].

## Installation

```bash
$ go get github.com/kamilsk/retry
$ # or use mirror
$ egg bitbucket.org/kamilsk/retry
```

> [egg][]<sup id="anchor-egg">[1](#egg)</sup> is an `extended go get`.

## Update

This library is using [SemVer](https://semver.org/) for versioning, and it is not
[BC](https://en.wikipedia.org/wiki/Backward_compatibility)-safe.

<sup id="egg">1</sup> The project is still in prototyping. [↩](#anchor-egg)

---

[![Gitter][icon_gitter]][gitter]
[![@kamilsk][icon_tw_author]][author]
[![@octolab][icon_tw_sponsor]][sponsor]

made with ❤️ by [OctoLab][octolab]

[awesome]:         https://github.com/avelino/awesome-go#utilities
[build]:           https://travis-ci.org/kamilsk/retry
[cli]:             https://github.com/kamilsk/retry.cli
[docs]:            https://godoc.org/github.com/kamilsk/retry
[gitter]:          https://gitter.im/kamilsk/retry
[license]:         LICENSE
[promo]:           https://github.com/kamilsk/retry
[quality]:         https://scrutinizer-ci.com/g/kamilsk/retry/?branch=v4
[research]:        https://github.com/kamilsk/go-research/tree/master/projects/retry
[legacy]:          https://github.com/kamilsk/retry/tree/master
[v3]:              https://github.com/kamilsk/retry/tree/v3
[v4]:              https://github.com/kamilsk/retry/projects/4

[egg]:             https://github.com/kamilsk/egg
[breaker]:         https://github.com/kamilsk/breaker
[gomod]:           https://github.com/golang/go/wiki/Modules
[platform]:        https://github.com/kamilsk/platform

[author]:          https://twitter.com/ikamilsk
[octolab]:         https://www.octolab.org/
[sponsor]:         https://twitter.com/octolab_inc
[support]:         https://www.patreon.com/octolab

[analytics]:       https://ga-beacon.appspot.com/UA-109817251-1/retry/v4?pixel
[tweet]:           https://twitter.com/intent/tweet?text=Functional%20mechanism%20to%20perform%20actions%20repetitively%20until%20successful&url=https://github.com/kamilsk/retry&via=ikamilsk&hashtags=go,repeat,retry,backoff,jitter

[icon_awesome]:    https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg
[icon_build]:      https://travis-ci.org/kamilsk/retry.svg?branch=v4
[icon_coverage]:   https://scrutinizer-ci.com/g/kamilsk/retry/badges/coverage.png?b=v4
[icon_docs]:       https://godoc.org/github.com/kamilsk/retry?status.svg
[icon_gitter]:     https://badges.gitter.im/Join%20Chat.svg
[icon_license]:    https://img.shields.io/badge/license-MIT-blue.svg
[icon_patreon]:    https://img.shields.io/badge/patreon-donate-orange.svg
[icon_quality]:    https://scrutinizer-ci.com/g/kamilsk/retry/badges/quality-score.png?b=v4
[icon_research]:   https://img.shields.io/badge/research-in%20progress-yellow.svg
[icon_tw_author]:  https://img.shields.io/badge/author-%40kamilsk-blue.svg
[icon_tw_sponsor]: https://img.shields.io/badge/sponsor-%40octolab-blue.svg
[icon_twitter]:    https://img.shields.io/twitter/url/http/shields.io.svg?style=social
