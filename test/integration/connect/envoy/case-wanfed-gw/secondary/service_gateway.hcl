# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: BUSL-1.1

services {
  name = "mesh-gateway"
  kind = "mesh-gateway"
  port = 4432
  meta {
    consul-wan-federation = "1"
  }
}
