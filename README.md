> # 🎳 forward
>
> `forward` - extended `kubectl port-forward` - reliable multiple port forwarding.

[![Build Status][icon_build]][page_build]

## 💡 Idea

```bash
$ forward postgresql 5432 redis 6379:6379
```

Full description of the idea is available
[here](https://www.notion.so/octolab/forward-94a09f0b2f6143d1b71d08edf3e52771?r=0b753cbf767346f5a6fd51194829a2f3).

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
$ REQ_VER=0.1.0  # all available versions are on https://github.com/kamilsk/forward/releases/
$ REQ_OS=Linux   # macOS is also available
$ REQ_ARCH=64bit # 32bit is also available
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

> [egg][page_egg]<sup id="anchor-egg">[1](#egg)</sup> is an `extended go get`.

### Bash and Zsh completions

```bash
$ forward completion -f bash > /path/to/bash_completion.d/forward.sh
$ forward completion -f zsh  > /path/to/zsh-completions/_forward.zsh
```

<sup id="egg">1</sup> The project is still in prototyping.[↩](#anchor-egg)

---

made with ❤️ for everyone

[icon_build]:      https://travis-ci.org/kamilsk/forward.svg?branch=master

[page_build]:      https://travis-ci.org/kamilsk/forward
[page_promo]:      https://github.com/kamilsk/forward
[page_egg]:        https://github.com/kamilsk/egg
