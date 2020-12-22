# blog-terratest-example
A starter template for running Terraform and Terratest in Azure. 

## How to run the example

Login to your azure subscription:
```shell
$ az login
```

On the `test` folder, install the go dependencies:
```shell
$ go get -t -v ./...
```

Finally, run the tests:
```shell
$ go test -timeout 30m
```
