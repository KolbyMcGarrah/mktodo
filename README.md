# mktodo
Mktodo is a small cli tool written in go that allows you to make Github issues from your command line. The aim is to make an easy way to keep track of all the things that need to be done later on a code base. 

## Usage
A github personal access token is required to be able to use the service. I reccommend storing that value in your bash config and passing it in the provided flag when running the command

To list all commands run
```
mktodo -h
```

To make a github issue run 
```
mktodo make
```
or 
```
mktodo m
```
Additional flags can be passed to supply known info (I reccommend using the --git-token flag and setting it equal to the user access token from your bashrc/zshrc file).

To see the list of all flags run 
```
mktodo make help
``` 
Any information not passed in flags will be need to be input from the terminal.

## Installation

This is available to install via [homebrew](https://brew.sh/).
After installing homebrew, you will need to connect to the tap with the following command
```
brew tap kolbymcgarrah/kolbymcgarrah
```

After that succeeds, you can install the cli tool with
```
brew install kolbymcgarrah/kolbymcgarrah/mktodo
```

You may need to have go installed on your machine which can be done by following [these instructions](https://go.dev/doc/install)