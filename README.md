# Benchmark

Benchmarking concurrent HTTP requests on a linux box

## Launch server

Firstly, setup your .env file

```
GOOGLE_CLOUD_PROJECT=...
```

**Build and deploy to GCP**

```bash
./build.sh
./deploy.sh
```

You can now run a benchmark

```bash
go build src/benchmark
ulimit -n 65535
./benchmark -url=http://localhost -n=10000 -c 10000
```
