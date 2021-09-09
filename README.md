# libprotoconf

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