# Outdoor.sy

![Build](https://github.com/crerwin/outdoorsy/actions/workflows/build.yaml/badge.svg)

## Running the code
If you have go >= 1.18 installed, you can run the code like so:
```
$ go install ./...
$ outdoorsy
```

If you don't have Go installed you can use Docker instead:
```
$ docker build . -t outdoorsy
$ docker run outdoorsy
```
If you are using docker, add `docker run` before any example commands.  Note that the input files are built into the docker image under `/data`.  If you want to supply your own data, you'll have to map it to a volume in the container.  Example:

```
$ docker run -v $(pwd)/data:/data outdoorsy import data/commas.txt
```

## The Outdoor.sy CLI

For help on running the `outdoorsy` cli, simply run it:
```
$ outdoorsy

or 

$ outdoorsy help
```

`outdoorsy` has subcommands, and the `import` command fulfills the requirements of the challenge.  For help on the `import` command, run it:
```
$ outdoorsy import

or 

$ outdoorsy help import
```

The CLI executes in a single run - the `import` command imports the specified files and outputs the data as requested.  State is not saved between runs, so you must specify all applicable input files each time the command is run.  The provided input files are checked in under `data/` and can be specified as arguments:

```
$ outdoorsy import data/commas.txt

or

$ outdoorsy import data/pipes.txt
```

Multiple files can be specified at once:
```
$ outdoorsy import data/commas.txt data/pipes.txt
```
Two larger csvs have been checked in for fun:
```
$ outdoorsy import data/mock_data.csv
$ outdoorsy import data/mock_data2.csv
```

You can use wildcards when specifying files:
```
$ outdoorsy import data/*.txt
$ outdoorsy import data/*.csv
$ outdoorsy import data/*
```


Import any other files not provided with the challenge by name:
```
$ outdoorsy import much_longer_test_data.txt
```

## Sorting data

In order to sort the data, add the `--sort` flag (or `-s`).  
```
$ outdoorsy import data/commas.txt --sort
```
By default, the customers are sorted by full name (see assumptions).  To sort by vehicle type, add `--sort-by=vehicle-type` (or `-b=vehicle-type`).
```
$ outdoorsy import data/commas.txt --sort --sort-by=vehicle-type
$ outdoorsy import data/commas.txt data/pipes.txt --sort --sort-by=vehicle-type
$ outdoorsy import data/* --sort --sort-by=vehicle-type
```

# Assumptions
I've made the following assumptions about the requirements of the challenge:
- The data is always in the same order (first name, last name, email, vehicle type, etc)
- A row will always use a consistent delimiter (all commas or all pipes)
- A file can contain rows that use either delimiter and won't necessary be consistent throughout the file.  I decide on the delimiter for each row independently.  
- The challenge requirements says "sort the data by Vehicle type or by Full name."  I take full name to mean that Abe Zed will sort before Val Ace (instead of sorting by last name then first).  

# Pipeline
A simple pipeline has been set up in Github Actions.  The configuration can be found in `.github/workflows`.  The pipeline does the following:

- lint the code
- run unit tests with both Go 1.18 and Go 1.19 on Ubuntu, MacOS, and Windows
- Perform multiarch builds of `outdoorsy` and save them as build artifacts:
```
outdoorsy-darwin-amd64
outdoorsy-darwin-arm64
outdoorsy-linux-amd64
outdoorsy-linux-arm64
outdoorsy-windows-amd64.exe
```
![pipeline](doc/pipeline.png?raw=true "Pipeline")

# Ruby version
I've added a very basic dirty Ruby script as a way of refreshing myself with Ruby syntax and RSpec.  It is found inside the `ruby/` folder. 

```
bundle install
bundle exec ruby outdoorsy.rb ../data/commas.txt
bundle exec ruby outdoorsy.rb ../data/commas.txt -n # (sort by name)
bundle exec ruby outdoorsy.rb ../data/commas.txt -t # (sort by vehicle type)
```

# What I'd do with more time!

## A few low-hanging fruits:
- Publish binaries when tagging releases on Github
- Integration tests - build the binary, then run the actual binary on the given platform with specific input files and inspect the output.  Inspect the output with various flags (`--sort`, `--sort-by`, `--debug`).  Even with this simple CLI program there was quite a bit of manual QA.

## Stretch Goals
- State - keep customers in a database so we don't have to keep loading files every time
- Outdoor.sy API - `/customers` to get customers, maybe with a `sort` parameter.  `/upload` enpoint allowing uploading of files.
- Front end - Display customers and at that point the sorting might as well happen on the front end.
- Deployments - build containers, deploy to Kubernetes or wherever.  Smoke test in between environments.  Deploy MRs to their own namespace.