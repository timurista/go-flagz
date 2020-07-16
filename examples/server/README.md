# Simple CLI demo

This demonstrates how dynamic values are being updated in a simple CLI app. Uses legacy V2 not newest v3.

## Quick set up V2 (etcd V3 not supported):

Download [etcd](https://github.com/coreos/etcd/releases), extract and make it available on your `$PATH`.

Launch `etcd ` server serving from a `default.data` in `/tmp`:

```sh
cd /tmp
etcd --enable-v2=true 
```

Set up a set of flags:

```sh
ETCDCTL_API=2 etcdctl mkdir /example/flagz
ETCDCTL_API=2 etcdctl set /example/flagz/staticint 9090
ETCDCTL_API=2 etcdctl set /example/flagz/dynstring foo
```

`go run .`

Play around by launching the server and visitng [http://localhost:8080](http://localhost:8080):

V2
```sh
ETCDCTL_API=2 etcdctl set /example/flagz/example_my_dynamic_string "I'm santient"
ETCDCTL_API=2 etcdctl set /example/flagz/example_my_dynamic_int 12345
```

V3
```sh
etcdctl put /example/flagz/example_my_dynamic_string "I'm santient"
etcdctl put /example/flagz/example_my_dynamic_int 12345
```

Marvel at the [flagz endpoint](http://localhost:8080/debug/flagz)).
