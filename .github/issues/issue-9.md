---
code: null
id: MDU6SXNzdWUzOTM4MTAzMTg=
databaseId: 393810318
number: 9
url: https://github.com/kamilsk/forward/issues/9
title: "support all options of kubectl port-forward"
issueType: null
labels:
  - "type:task"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: OPEN
stateReason: null
duplicateOf: null
createdAt: 2018-12-24T04:55:09Z
updatedAt: 2019-05-14T07:37:30Z
lastEditedAt: 2018-12-24T04:56:11Z
closedAt: null
---

# support all options of kubectl port-forward

- [x] kubectl port-forward mypod 5000 6000
Listen on ports 5000 and 6000 locally, forwarding data to/from ports 5000 and 6000 in the pod
- [x] kubectl port-forward mypod 8888:5000
Listen on port 8888 locally, forwarding to 5000 in the pod
- [ ] kubectl port-forward mypod :5000
Listen on a random port locally, forwarding to 5000 in the pod

<!-- 2019-05-14T07:37Z https://github.com/kamilsk/forward/issues/9#issuecomment-492115425
do I need to generate some port or use default os behavior? how I can receive port number in last case?
-->
