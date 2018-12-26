# insult-jmk

A small micro service which can produce insults in french. Used to insult JMK with DankFaceBot.

## Use

```bash
go run main.go
```

## Create Travis secret

Log in to travis

```bash
travis login --com
```

Then encrypt

```bash
travis encrypt-file client-secret.json --add
```

## GRPC

### Compile proto

Inside the proto folder :

```bash
protoc --go_out=plugins=grpc:. *.proto
```

### Test GRPC

Use npm `grpcc`.

```bash
grpcc -p grpc/insult-jmk.proto -a localhost:50051 -i
```

Then, in the REPL

```javascript
client.getInsult({name: "dixneuf19"}, printReply)
```