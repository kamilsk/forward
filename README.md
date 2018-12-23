> # 🎳 forward
>
> `forward` - extended `kubectl port-forward` - multiple port forwarding simultaneously.

[![Patreon][icon_patreon]](https://www.patreon.com/octolab)
[![Build Status][icon_build]][page_build]
[![License][icon_license]](LICENSE)

## Motivation

Instead of

```bash
$ kubectl get pod
NAME                                  READY     STATUS    RESTARTS   AGE
site-5d7f49cf95-zsct2                 4/4       Running   0          1d
catalog-79c558d96-zg6cg               1/1       Running   0          1d
catalog-postgresql-7595dd6b9c-fkrbz   1/1       Running   0          1d
catalog-redis-76bbdf658b-4zdwc        1/1       Running   0          1d
site-redis-b654f56d4-55kvk            1/1       Running   0          1d
site-rabbitmq-7677fdf798-flswj        1/1       Running   0          1d
$ kubectl port-forward catalog-postgresql-7595dd6b9c-fkrbz 5432:5432 &
$ kubectl port-forward catalog-redis-76bbdf658b-4zdwc      6379:6379 &
$ ps x | fgrep 'kubectl port-forward ...' | xargs kill -SIGKILL

It's so boring... （╯°□°）╯︵┻━┻
```

I want to

```bash
$ forward postgresql 5432:5432 redis 6379:6379
which redis?
> catalog-redis-76bbdf658b-4zdwc
  site-redis-b654f56d4-55kvk
```

## Roadmap

- [ ] v1: [MVP][project_v1]
  - [**Someday, 20xx**][project_v1_dl]
  - Main concepts and working prototype.
- [ ] v2: [Rate limiting][project_v2]
  - [**Somehow, 20xx**][project_v2_dl]
  - Better integration with [Kubernetes](https://kubernetes.io/).

## Demo

[![asciicast](https://asciinema.org/a/217993.svg)](https://asciinema.org/a/217993)

## Installation

### Homebrew

```bash
$ brew install kamilsk/tap/forward
```

### Binary

```bash
$ export REQ_VER=0.1.0  # all available versions are on https://github.com/kamilsk/forward/releases/
$ export REQ_OS=Linux   # macOS is also available
$ export REQ_ARCH=64bit # 32bit is also available
# wget -q -O forward.tar.gz
$ curl -sL -o forward.tar.gz \
       https://github.com/kamilsk/forward/releases/download/"${REQ_VER}/forward_${REQ_VER}_${REQ_OS}-${REQ_ARCH}".tar.gz
$ tar xf forward.tar.gz -C "${GOPATH}"/bin/ && rm forward.tar.gz
```

### From source code

```bash
# using standard go tools
$ go get -u github.com/kamilsk/forward
# or using egg tool
$ egg github.com/kamilsk/forward -- go install .
# with mirror
$ egg bitbucket.org/kamilsk/forward -- go install .
```

> [egg](https://github.com/kamilsk/egg)<sup id="anchor-egg">[1](#egg)</sup> is an `extended go get`.

<sup id="egg">1</sup> The project is still in prototyping.[↩](#anchor-egg)

---

[![@kamilsk][icon_tw_author]](https://twitter.com/ikamilsk)
[![@octolab][icon_tw_sponsor]](https://twitter.com/octolab_inc)

made with ❤️ by [OctoLab](https://www.octolab.org/)

[icon_build]:      https://travis-ci.org/kamilsk/forward.svg?branch=master
[icon_license]:    https://img.shields.io/badge/license-MIT-blue.svg
[icon_patreon]:    https://img.shields.io/badge/patreon-donate-orange.svg
[icon_tw_author]:  https://img.shields.io/badge/author-%40kamilsk-blue.svg
[icon_tw_sponsor]: https://img.shields.io/badge/sponsor-%40octolab-blue.svg
[icon_twitter]:    https://img.shields.io/twitter/url/http/shields.io.svg?style=social

[page_build]:      https://travis-ci.org/kamilsk/forward
[page_promo]:      https://github.com/kamilsk/forward

[project_v1]:      https://github.com/kamilsk/forward/projects/1
[project_v1_dl]:   https://github.com/kamilsk/forward/milestone/1
[project_v2]:      https://github.com/kamilsk/forward/projects/2
[project_v2_dl]:   https://github.com/kamilsk/forward/milestone/2
