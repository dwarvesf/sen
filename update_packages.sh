#!/bin/bash

set -euo pipefail

glide up --delete --force --update-vendored --strip-vcs --strip-vendor
