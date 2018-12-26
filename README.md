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

<!-- Then add at the a before

```bash
openssl aes-256-cbc -K $encrypted_8301fcd250ef_key -iv $encrypted_8301fcd250ef_iv -in client-secret.json.enc -out client-secret.json -d
``` -->