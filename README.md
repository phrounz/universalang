# universalang

A silly attempt to do an universal language, which would use for each word, an "average" of the letters/sounds of the same word in many languages around the world.

Early development.

To run:
```
go build -o universalang ./src && ./universalang
```
Use -v for verbose mode.

Example of current translations (english: universalang):
```
i: we
hello: hola
this: ke
be: ji
future: futar
past: pas
eat: take
soccer: sutba
language: langa
speak: habla
do: fur
tea: te
yes: yei
no: no
love: aiea
```

Because of the way the algorithm currently works, each execution may return different values.

You can use Google Translate by setting constUseExamples=false but it's not tested and requires a Google Translate API token.
