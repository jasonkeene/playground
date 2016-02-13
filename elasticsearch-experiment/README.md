
## Example

Starting elastic search:

```bash
./docker-up.sh
```

Inserting document:

```bash
go run write.go -url "http://192.168.99.100:32773" -index debug -type test -message "hello world"
```

Reading documents:

```bash
go run read.go -url "http://192.168.99.100:32773" -index debug -type test
```
