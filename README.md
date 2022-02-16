# dwlr

A downloader written in golang, created with the aim of quickly downloading files in batch concurrently and composing a security automation recognition process.

## Example usages

Single URL:

```
echo https://google.github.io/styleguide/include/styleguide.js | dwlr
```

Multiple URLs:

```
cat urls.txt | dwlr
```

## Installation

First, you'll need to [install go](https://golang.org/doc/install).

Then run this command to download + compile dwlr:
```
go install github.com/edivangalindo/dwlr@latest
```

You can now run `~/go/bin/dwlr`. If you'd like to just run `dwlr` without the full path, you'll need to `export PATH="/go/bin/:$PATH"`. You can also add this line to your `~/.bashrc` file if you'd like this to persist.