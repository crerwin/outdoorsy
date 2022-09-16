# Outdoor.sy

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
If you are using docker, add `docker run` before any example commands.  Note that the input files are build into the docker image under `/data`.  If you want to supply your own data, you'll have to map it to a volume in the container.  Example:

```
$ docker run -v $(pwd)/data:/data outdoorsy import data/commas.txt
```

## The Outdoor.sy CLI

`outdoorsy` has subcommands, and the `import` command fulfills the requirements of the challenge.  The CLI executes in a single run - the `import` command imports the specified files and outputs the data as requested.  State is not saved between runs, so you must specify all applicable input files each time the command is run.  The provided input files are checked in under `data/` and can be specified as arguments:

```
$ outdoorsy import data/commas.txt

or

$ outdoorsy import data/pipes.txt
```

Multiple files can be specified at once:
```
$ outdoorsy import data/commas.txt data/pipes.txt
```

Import any other files not provided with the challenge by name:
```
$ outdoorsy import much_longer_test_data.txt
```

## Sorting data

In order to sort the data, add the `--sort` flag (or `-s`).  By default, the customers are sorted by full name (see assumptions).  To sort by vehicle type, add `--sort-by=vehicle-type` (or `-b=vehicle-type`).
```
outdoorsy import data/commas.txt --sort --sort-by=vehicle-type
```

# Assumptions
I've made the following assumptions about the requirements of the challenge:
- The data is always in the same order (first name, last name, email, vehicle type, etc)
- A row will always use a consistent delimiter (all commas or all pipes)
- A file can contain rows that use either delimiter and won't necessary be consistent throughout the file.  I decide on the delimiter for each row independently.  
- The challenge requirements says "sort the data by Vehicle type or by Full name."  I take full name to mean that Abe Zed will sort before Val Ace (instead of sorting by last name then first).  