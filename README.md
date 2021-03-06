# .//libprotoconf

[![Test and coverage](https://github.com/protoconf/libprotoconf/actions/workflows/ci.yml/badge.svg)](https://github.com/protoconf/libprotoconf/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/protoconf/libprotoconf/branch/main/graph/badge.svg?token=4YsYeTWEYa)](https://codecov.io/gh/protoconf/libprotoconf)
[![Go Reference](https://pkg.go.dev/badge/github.com/protoconf/libprotoconf.svg)](https://pkg.go.dev/github.com/protoconf/libprotoconf)

Runtime configuration made easy with `protocol-buffers`.

## Goals

Developers can define their desired configuration struct in a proto file, then get all sorts of ways to load that configuration into their software runtime:

- [x] Populate config from environment variables
- [x] Get a `flag.FlagSet` instance
- [ ] Read and merge configs from global, user and workspace locations (`/etc/mypackage/`, `$USER/.mypackage/`, `$(pwd)/.mypackage/`)
- Allow configuration changes dynamically
  - [ ] Reload on Signals
  - [ ] Reload on file change (`inotify`)
  - [ ] `gRPC`/`rest` configuration endpoints
  - [ ] key-value stores (`consul`, `etcd`, `zookeeper`, `redis`)
  - [ ] protoconf gRPC streaming agent
- Formats:
  - [x] JSON (`.json`)
  - [x] YAML (`.yaml`, `.yml`)
  - [x] pbtext (`.pb`)
  - [x] pb binary (`.data`)
  - [x] pb binary with base64 encoding (`.base64`, `.b64`)
  - [x] jsonnet (`.jsonnet`)
  - [x] toml (`.toml`)
  - [ ] hcl (`.hcl`)
  - [ ] cue (`.cue`)
- [ ] Pre/Post hooks for dynamic changes
- Debugging features
  - [ ] `ShowConfig()` method (available via gRPC/rest)
  - [ ] `ConfigMetadata()` method (available via gRPC/rest)
    - [ ] Where a specific variable was set (env/flag/which file/key value store/rpc)
    - [ ] When was the change applied to runtime
    - [ ] When was the change made (if made before the process was starting)
    - [ ] Who made the change (if data available)
  - [ ] `/varz` http endpoint
  - [ ] `/metrics` http endpoint for prometheus (with `ConfigMetadata()` information available)

## Implementations

- First implementation will be in Go. This implementation will also include the `protoconf` gRPC streaming agent, so developers of any language can use the dynamic reload features as a sidecar agent.
- Future implementations (please vote!)
  - Rust (will probably replace Go for the agent implementation)
  - Python
  - JavaScript
  - Java
  - C++
  - Suggest other languages
