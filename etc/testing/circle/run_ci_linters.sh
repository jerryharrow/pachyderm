#!/bin/bash

set -exo pipefail

make lint
make enterprise-code-checkin-test
make test-proto-static