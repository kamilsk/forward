#!/usr/bin/env bash

echo I am gonna sleep

sleep 0.1

echo Zzz...

sleep ${1:-0.1}

exit 1
