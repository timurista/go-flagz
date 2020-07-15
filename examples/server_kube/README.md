# Kubernetes ConfigMap example

- First download [etcd](https://github.com/coreos/etcd/releases), extract and make it available on your `$PATH`. (might need to move to `/user/local/bin`)
- run it by invoking binary `etcd`
- initialize empty tmp mapping `mkdir -p /tmp/foobar`
- run the server `go run .`
- visit the debug `http://localhost:8080/debug/flagz`

Should see this if successful:
![](flagz-debug-view.png)
