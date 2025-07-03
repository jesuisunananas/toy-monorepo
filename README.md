# toy-monorepo
toy monorepo to experiment with exposing prometheus metrics and pprof

# install
brew install prometheus # to locally run prometheus server
brew install graphviz # to download/view graphs from pprof

# commands
go run main.go client.go
prometheus --config.file=prometheus.yaml # open up http://localhost:9090/query and go to status/target-health, confirm state is up
curl http://localhost:6060/debug/pprof/heap > heap.pb.gz # heap snapshot
go tool pprof -http=:8080 heap.pb.gz # should open up pprof in web browser to view flame graph/call graph etc.
watch -n 1 curl localhost:6060/debug/vars # to watch memory live
