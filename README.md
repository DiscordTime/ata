# Android Test Automation

[![N|Solid](https://cldup.com/dTxpPi9lDf.thumb.png)](https://nodesource.com/products/nsolid)

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)


[comment]: <> (TODO Description)
Project description


## Clone Project

```bash
$ git clone git@github.com:DiscordTime/ata.git
$ cd ata
```

## Compiling

```bash
$ make
```
This will create a binary(ata) on current folder and set a version variable to it.

## Running

```bash
$ make run
```

This command will run ./ata

## Testing all possible args

```bash
$ make test
```

This will run all possible args in ata binary.

## Installing

```bash
$ make install
```

This will install ata binary on ~/.local/bin/ and install completion on /etc/bash\_completion.d/
> :warning: **To change binary installation path you can set the variable BIN_PATH when do installor on environment as the example below.**

```bash
$ make install BIN_PATH=/usr/bin/
```

## Uninstalling

```bash
$ make uninstall
```

This will remove completion.bash file and ata binary.
> :warning: **If you change the BIN_PATH on installation you may set the correct BIN_PATH to uninstall.**

```bash
$ make uninstall BIN_PATH=/usr/bin
```
