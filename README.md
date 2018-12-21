> # 🎳 forward
>
> **forward** - extended kubectl port-forward - multiple port forwarding at the same time.

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

(ง̀-́)ง boring...
```

I want to

```bash
$ forward postgresql 5432:5432 redis 6379:6379
which redis?
> catalog-redis-76bbdf658b-4zdwc
  site-redis-b654f56d4-55kvk
```

## TODO

- [ ] pod name suggestion (autocomplete)
- [ ] ports suggestion based on pod description
- [ ] pass arguments
- [ ] detach mode
- [ ] [gops](https://github.com/google/gops) integration
- [ ] better process management
  - stop forwarding for part ports/pods
  - signal handling
- [ ] better kubernetes integration (API instead CLI)

---

[![@kamilsk][icon_tw_author]](https://twitter.com/ikamilsk)
[![@octolab][icon_tw_sponsor]](https://twitter.com/octolab_inc)

made with ❤️ by [OctoLab](https://www.octolab.org/)

[icon_license]:    https://img.shields.io/badge/license-MIT-blue.svg
[icon_patreon]:    https://img.shields.io/badge/patreon-donate-orange.svg
[icon_tw_author]:  https://img.shields.io/badge/author-%40kamilsk-blue.svg
[icon_tw_sponsor]: https://img.shields.io/badge/sponsor-%40octolab-blue.svg
[icon_twitter]:    https://img.shields.io/twitter/url/http/shields.io.svg?style=social
