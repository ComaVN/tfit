Find out what a piece of text represents.

tfit tries to detect what a particular piece of text represents: It's intended
to automatically decode and/or interpret any of a large number of encodings,
formats, etc.

Check the [Dockerfile](Dockerfile) to see what's required to build.

The inspiration for this project came during a sleepless night in a lodge at 5,150m in Gorak Shep, Nepal.

```
        __/|
   |\_--    \
   |      (O \
   \  (O ^   |
 /  \        \
|\   |        \
 \\  /         \
  \\/          |
   V           \
```

## Dev instructions

To build the main executable and run it:
```
docker run -i --rm $(docker build -q .) < test/data/foo.base64
```
