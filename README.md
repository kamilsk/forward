> # 🎳 forward
>
> `forward` - extended `kubectl port-forward` - reliable multiple port forwarding.

[![Build][build.icon]][build.page]
[![Template][template.icon]][template.page]

## 💡 Idea

```bash
$ forward postgresql 5432 redis 6379:6379
```

Full description of the idea is available [here][design.page].

## 🏆 Motivation

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
$ forward postgresql 5432 redis 6379:6379
which redis?
> catalog-redis-76bbdf658b-4zdwc
  site-redis-b654f56d4-55kvk
```

## 🤼‍♂️ How to

[![asciicast](https://asciinema.org/a/217993.svg)](https://asciinema.org/a/217993)

## 🧩 Installation

### Homebrew

```bash
$ brew install kamilsk/tap/forward
```

### Binary

```bash
$ curl -sSL https://bit.ly/install-forward | sh
# or
$ wget -qO- https://bit.ly/install-forward | sh
```

### Source

```bash
# use standard go tools
$ go get -u github.com/kamilsk/forward
# or use egg tool
$ egg tools add github.com/kamilsk/forward
```

> [egg][egg.page]<sup id="anchor-egg">[1](#egg)</sup> is an `extended go get`.

### Bash and Zsh completions

```bash
$ forward completion bash > /path/to/bash_completion.d/forward.sh
$ forward completion zsh  > /path/to/zsh-completions/_forward.zsh
```

<sup id="egg">1</sup> The project is still in prototyping.[↩](#anchor-egg)

---

made with ❤️ for everyone

[build.icon]:       https://travis-ci.org/kamilsk/forward.svg?branch=master
[build.page]:       https://travis-ci.org/kamilsk/forward

[design.page]:      https://www.notion.so/octolab/forward-94a09f0b2f6143d1b71d08edf3e52771?r=0b753cbf767346f5a6fd51194829a2f3

[promo.page]:       https://github.com/kamilsk/forward

[template.page]:    https://github.com/octomation/go-tool
[template.icon]:    https://img.shields.io/badge/template-go--tool-blue

[egg.page]:         https://github.com/kamilsk/egg
