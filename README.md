HTTP server
===========

This is Go (Golang) sample code for a minimal HTTP server. Learning purposes.

jsonserver-raw
--------------

Barebones http json server, it just uses a `switch` based HTTP router.

```
$ ab -v 9 http://localhost:8080/sample1
$ ab -v 9 http://localhost:8080/sample2
```

jsonserver-pkgjson
------------------

Barebones http json server, it just uses a `switch` based HTTP router.
Codifies JSON with `encoding/json` Go package

```
$ ab -v 9 http://localhost:8080/sample1
$ ab -v 9 http://localhost:8080/sample2
```

jsonserver-jsonwithhandler
--------------------------

Uses two server muxes

```
$ ab -v 9 http://localhost:8080/sample/1
$ ab -v 9 http://localhost:8080/sample/2
$ ab -v 9 http://localhost:8080/tmp/hola
```

jsonserver-gocraft 
------------------

Uses [gocraft/web](https://github.com/gocraft/web). Sample from its README.md.



