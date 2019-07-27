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
    $ swagger generate client -f ./swagger/sca_v1_spec_sca.json
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
