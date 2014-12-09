# go-tugboat
go-tugboat was heavily inspired by the go-github and tugboat library.

[![Build Status](https://drone.io/github.com/masayukioguni/go-tugboat/status.png)](https://drone.io/github.com/masayukioguni/go-tugboat/latest)
[![Build Status](https://travis-ci.org/masayukioguni/go-tugboat.svg?branch=master)](https://travis-ci.org/masayukioguni/go-tugboat)

DigitalOcean API v2 command line tool for interacting with your [DigitalOcean](https://www.digitalocean.com/) droplets.

## Installation

    go get github.com/masayukioguni/go-tugboat  

## Configuration

Run the configuration utility, `go-tugboat authorize`. You can grab your keys
[here](https://cloud.digitalocean.com/settings/applications).

    $ go-tugboat authorize
    Entser your API key:foo
    To retrieve region, image, size and key ID's, you can use the corresponding go-tugboat command,
    such as go-tugboat images. Defaults can be changed at any time in your ~/.go-tugboat/config.yaml configuration     file.
    Enter your default region (optional, defaults to nyc3 (New York)):
    Enter your default image id(optional, defaults to 6918990(Ubuntu 14.04 x64):
    Enter your default size (optional, defaults to 512MB:

    Authentication with DigitalOcean was successful!

## Usage

### Retrieve a list of your droplets

    $ go-tugboat droplets
    test (ip: xxx.xxx.xxx.xxx, status: active, region :nyc1, id: 3395705)
    test1(ip: xxx.xxx.xxx.xxx, status: active, region :nyc1, id: 3395706)

### Create a droplet

    $ go-tugboat create -name=test 
    Queueing creation of droplet 'test1' ...done

### Destroy a droplet

    $ go-tugboat destroy -id=3402715
    Queuing destroy for 3402715 ...done

### List Available Images

You can list images that you have created.

list images provided by DigitalOcean as well.

    $ go-tugboat images
    My Images:
    test (id: 7979948, distro: Ubuntu)
     ....
    Global Images:
    Ruby on Rails on 14.04 (Nginx + Unicorn) (id: 6376601, distro: Ubuntu)
    ...

list all Distribution Images

    $ go-tugboat images -type=dist
    20 x64 (id: 6370882, distro: Fedora)
    ....

List all application images
   $ go-tugboat images -type=app 
   Ruby on Rails on 14.04 (Nginx + Unicorn) (id: 6376601, distro: Ubuntu)
   ....
   
### List Available Sizes

    $ go-tugboat sizes
    slug:512mb memory: 512mb vcpus: 1 disk: 20gb transfer:1.0tb monthly:5.0$
    slug:  1gb memory: 1024mb vcpus: 1 disk: 30gb transfer:2.0tb monthly:10.0$
    slug:  2gb memory: 2048mb vcpus: 2 disk: 40gb transfer:3.0tb monthly:20.0$
    ...

### List Available Regions

    $ go-tugboat regions
    Regions:
    New York 1 (id: 1) (slug: nyc1)
    Amsterdam 1 (id: 2) (slug: ams1)
    San Francisco 1 (id: 3) (slug: sfo1)

### List SSH Keys

    $ go-tugboat keys
    id:xxxxxx name:masayuki.oguni
    ...

## Help

If you're curious about command flags for a specific command, you can
ask go-tugboat about it.

    $ go-tugboat --help create

For a complete overview of all of the available commands, run:

    $ go-tugboat help

## Reporting Bugs

Yes, please!

You can create a new issue [here](https://github.com/masayukioguni/go-tugboat/issues/new). Thank you!
