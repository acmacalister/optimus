optimus
=======

Buildin iOS enterprise apps yo.

## Install

`go get github.com/acmacalister/optimus`

`go install github.com/acmacalister/optimus`

## Example

optimus is utility you install, not a library. Maybe in the future, but for now head over to the releases section and download yourself a compiled version or checkout and build/install it yourself. Below is the example configuration that optimus needs to build an app.

```yaml
development:
  aws_key: your_s3_key
  aws_secret: your_s3_secret
  bucket: mybucket
  path: folderinmybucket
  signing_identity: iPhone Distribution: Your Team
  project_name: yourprojectname
staging:
  aws_key: your_s3_key
  aws_secret: your_s3_secret
  bucket: mybucket
  path: folderinmybucket
  signing_identity: iPhone Distribution: Your Team
  project_name: yourprojectname
production:
  aws_key: your_s3_key
  aws_secret: your_s3_secret
  bucket: mybucket
  path: folderinmybucket
  signing_identity: iPhone Distribution: Your Team
  project_name: yourprojectname
```

## Supported Platforms

OS X only of course. Still very much alpha software. We are using it internally to build and upload our apps. If you have any issues with your app, please open an issue.

## Help

* [Github](https://github.com/acmacalister)
* [Twitter](http://twitter.com/acmacalister)

## ToDo

* We be needin some test