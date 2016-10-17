# s3-signed-redirect

A very small webserver that returns a 302 redirect with a Presigned S3 URL for a specific bucket, with a configurable Expires time.
Designed for use with nginx `proxy_pass` or similar.

Any path requested will be presigned as a S3 request and a 302 redirect issued for the configured bucket.

## Config

Config is yaml based and loaded from the following directories:

* `/etc/s3-signed-redirect/s3-signed-redirect.yaml`
* `~/.s3-signed-redirect/s3-signed-redirect.yaml`
* `./s3-signed-redirect.yaml`

Config is also loaded from ENV variables, the same names as the yaml keys prefixed with `S3SR` e.g. `S3SR_PORT`.
`AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` should also work.

```yaml
aws_access_key_id: foo
aws_secret_access_key: bar
s3_host: rjoc-bbc2.s3-eu-west-1.amazonaws.com

address: 127.0.0.1
port: 3000
timeout: 7200s
debug: false
```