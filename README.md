# zeppelin
Zeppelin is a tiny web server for uploading and downloading files able to run on close to any system without dependencies and with some extra pentester toys.

## Build:

```
git clone https://github.com/hcninja/zeppelin.git
cd zeppelin
cargo build --release
```

You will find your compiled binary under `target/release/zeppelin`.

Or just download the latest release for your architecture directly from [releases](https://github.com/hcninja/zeppelin/releases/).

## Usage:
```
OPTIONS:
    -h, --host <HOST>    Host address for the file-server [127.0.0.1]
    -d, --dir <PATH>     Directory to serve [./]
    -p, --port <PORT>    Port for the file-server [8080]
```

## ToDo:
- [x] Navigate the served directory
- [ ] File upload
- [ ] System command execution
- [ ] Web interface for logging request with headers
- [ ] Process injection/migration
