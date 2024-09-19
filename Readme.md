# Temporal learning


download 
```
wget https://temporal.download/cli/archive/latest\?platform\=linux\&arch\=amd64  
```

untar 
```
tar -xvf temporal_cli_1.1.0_linux_amd64.tar.gz
```

environment up

```
make local-up
```

environment down
```
make local-down
```

run worker 

```
go run ./src/worker/worker.go  
```

run starter

```
go run ./src/starter/starter.go

```