btsync-cli
============

Console Application for BitTorrent SyncApp API written in Go
Beta version. Tested against 1.2.82.

Features
------
  * add folder to sync folders (with empty or predefined secret)
  * remove folder from sync folders
  * list all sync folders with read_only secrets
  * generate secret (with readonly secret)

Example Usage
------

    Usage of ./cli:
      -a="":                absolute path to add for index (-r for relative path support)
      -d="":                delete folder by secret
      -g=false:             get new secret
      -host="127.0.0.1":    btsync hostname
      -l=false:             list folders (secret, read-only secret, type, path)
      -p="123456":          password
      -port="8888":         btsync port
      -r=false:             resolve relative path (for -a)
      -s="":                secret, if empty will be autogenerated
      -u="admin":           username
      -v=false:             verbose mode on


Known Issues
------
  * Btsync api is not fully implemented
  * Bad Russian support
  * You need special API Key http://www.bittorrent.com/intl/ru/sync/developers


How-to build on Synology DS210j
------
    cd ~btsync-cli
    export GOPATH=`pwd`
    export GOARM=5
    /opt/go/bin/go build -o btsync-cli btsync/cli
    ./btsync-cli -h

Changelog
------
v0.3
 * support official API instead of webui
 * delete folder by secret without path
 * generate readonly secret for folder

v0.2
  * -r flag for relative path on server
  * tests
