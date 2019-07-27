# srcclr-cli
SourceClear CLI

## Development

1. Install go-swagger
    ```bash
    $ sudo curl -L -o /usr/local/bin/swagger https://github.com/go-swagger/go-swagger/releases/download/v0.19.0/swagger_linux_amd64
    $ chmod 755 /usr/local/bin/swagger
    $ swagger version
    version: v0.19.0
    commit: 312366608bbf17dd219190b66ab63bdc8b4d0db0
    ```
1. Download SourceClear API Swagger specification file ([sca_v1_spec_sca.json](https://help.veracode.com/viewer/book-attachment/LMv_dtSHyb7iIxAQznC~9w/TJ9aAyVCY8DBEwCSH89EQQ)).
1. Generate an API client.
    ```bash
    # NOTE: go-swagger does not support Go module yet.
    # Your source tree must be in GOPATH.
    # You must hit `direnv deny` and deactivate Go module if you use direnv. 
    $ swagger generate client -f ./swagger/sca_v1_spec_sca.json
    ...
    2019/01/01 12:00:00 Generation completed!

    For this generation to compile you need to have some packages in your GOPATH:

            * github.com/go-openapi/errors
            * github.com/go-openapi/runtime
            * github.com/go-openapi/runtime/client
            * github.com/go-openapi/strfmt

    You can get these now with: go get -u -f ./...
    
    # activate Go module
    $ direnv allow
    $ go get -u -f ./...
    $ go mod tidy
    ```

## Known Issues

### Invalid swagger specs

```bash
$ swagger validate swagger/sca_v1_spec_sca.json 
...
The swagger spec at "swagger/sca_v1_spec_sca.json" is invalid against swagger specification 2.0.
See errors below:
- default value for status in query does not validate its schema
- status in query must be of type array: "string"
```
