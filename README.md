# Outdoor.sy

## Running the code
If you have go >= 1.18 installed, you can run the code like so:
```
$ go install ./...
$ outdoorsy
```

`outdoorsy` has subcommands, and the `import` command fulfills the requirements 
of the challenge.  The provided input files are checked in under `data/` and
can be specified as arguments:

```
$ outdoorsy import data/commas.txt

or

$ outdoorsy import data/pipes.txt
```

Multiple files can be specified at once:
```
$ outdoorsy import data/commas.txt data/pipes.txt
```