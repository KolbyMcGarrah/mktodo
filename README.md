# mktodo
Mktodo is a small cli tool written in go that allows you to make Github issues from your command line. The aim is to make an easy way to keep track of all the things that need to be done later on a code base. 

## Usage
A github personal access token is required to be able to use the service. I reccommend storing that value in your bash config and passing it in the provided flag when running the command

To list all commands run
```
ghtodo -h
```

To make a github issue run 
```
ghtodo mk
```
or 
```
ghtodo m
```
Additional flags can be passed to supply known info (I reccommend using the --git-token flag and setting it equal to the user access token from your bashrc/zshrc file).

To see the list of all flags run 
```
ghtodo mk help
``` 
Any information not passed in flags will be need to be input from the terminal.

## Installation
Currently this can be installed by copying the repo or downloading the `ghtodo` binary.
After aquiring the binary, mv it to your `/usr/local/bin` folder and then ensure you have the proper permissions by running
```
chmod + /usr/local/bin/ghtodo
```
then you can run the command via the terminal with the `ghtodo` command.