![Zeppelin](assets/Zeppelin_header.png)

Zeppelin is a tiny web server for uploading and downloading files able to run on close to any system without dependencies and with some extra pentester toys.

## Build:

For compiling _zeppelin_ you need to have a working go instalation, please visit [golang.org](https://golang.org) for detailed instructions on how to do this.

After, its as easy as issuing the following commands:
```
git clone https://github.com/hcninja/zeppelin.git
cd zeppelin
go build
```

Or just download the latest release for your architecture directly from [releases](https://github.com/hcninja/zeppelin/releases/).

## Usage:
With `./zeppelin -h` you will get the available flags.

```
Usage of ./zeppelin:
  -host string
    	Server host (default "127.0.0.1")
  -path string
    	Path to serve and upload files to (default "./")
  -port string
    	Server port (default "8080")
  -tls
    	[NYI] Enables TLS. Cert and key must be in root 'cert.pem and key.pem'
  -unsafe
    	Removes the file upload limit of 8MB (default true)
```

The standard running mode, is to serve the path where `zeppelin` is running on. For serving the `/etc` directory on all interfaces on port 8443 you can use `./zeppelin -host 0.0.0.0 -pots 8443 -path /etc`.

## ToDo:
- [x] Navigate the served directory
- [x] File upload
- [ ] HTTPS
- [ ] System command execution
- [ ] File navigator and uploader authentication
- [ ] Web interface for logging request with headers
- [ ] Process injection/migration
