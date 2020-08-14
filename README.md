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
  -nocmd
    	Disables CMD endpoints
  -noupload
    	Disables the upload endpoint
  -path string
    	Path to serve and upload files to (default "./")
  -port string
    	Server port (default "8080")
  -tls
    	Enables TLS. Cert and key must be in root 'cert.pem and key.pem'
  -unsafe
    	Removes the file upload limit of 8MB
```

The standard running mode, is to serve the path where `zeppelin` is running on. For serving the `/etc` directory on all interfaces on port 8443 you can use `./zeppelin -host 0.0.0.0 -pots 8443 -path /etc`.

### Generating a testing self-signed certificate
In order to enable TLS on _Zeppelin_ you must provide a certificate and a key in pem format, _Zeppelin_ will search for this two files under the name "cert.pem" and "key.pem".

If you have no available certificate (you can use a Let's Encrypt one for free), you can generate a self-signed certificate with openssl:

```bash
openssl req -x509 -out cert.pem -keyout key.pem \
  -newkey rsa:2048 -nodes -sha256 \
  -subj '/CN=localhost' -extensions EXT -config <( \
   printf "[dn]\nCN=localhost\n[req]\ndistinguished_name = dn\n[EXT]\nsubjectAltName=DNS:127.0.0.1\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth")
```

## ToDo:
- [x] Navigate the served directory
- [x] File upload
- [x] HTTPS
- [x] System command execution
- [ ] File navigator and uploader authentication
- [ ] Web interface for logging request with headers
- [ ] Process injection/migration
