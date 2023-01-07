<div align="right">

[![golangci-lint](https://github.com/sheepla/whois-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/sheepla/whois-cli/actions/workflows/ci.yml)

[![Release](https://github.com/sheepla/whois-cli/actions/workflows/release.yml/badge.svg)](https://github.com/sheepla/whois-cli/actions/workflows/release.yml)

</div>


<div align="center">

# whois-cli

A simple command line [whois](https://en.wikipedia.org/wiki/Whois) client

</div>


## Usage

```
NAME:
   whois - whois CLI

USAGE:
   whois [global options] command [command options] QUERY

DESCRIPTION:
   whois CLI

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --json, -j  Output in JSON format (default: false)
   --help, -h  show help (default: false)
```

Specifying a domain name as an argument, owner information for that domain will be displayed.

```
$ whois github.com
::
:: DOMAIN
::
       ITEM      |             VALUE
-----------------+---------------------------------
  ID             | 1264983250_DOMAIN_COM-VRSN
  Domain         | github.com
  Punycode       | github.com
  Name           | github
  Extension      | com
  WhoisServer    | whois.markmonitor.com
  Status         | clientdeleteprohibited
                 | clienttransferprohibited
                 | clientupdateprohibited
  NameServers    | dns1.p08.nsone.net
                 | dns2.p08.nsone.net
                 | dns3.p08.nsone.net
                 | dns4.p08.nsone.net
                 | ns-1283.awsdns-32.org
                 | ns-1707.awsdns-21.co.uk
                 | ns-421.awsdns-52.com
                 | ns-520.awsdns-01.net
  DNSSec         | false
  CreatedDate    | 2007-10-09T18:20:50Z
  UpdatedDate    | 2022-09-07T09:10:44Z
  ExpirationDate | 2024-10-09T18:20:50Z

::
:: REGISTRAR
::
     ITEM     |              VALUE
--------------+----------------------------------
  ID          |                             292
  Name        | MarkMonitor Inc.
  Phone       | +1.2086851750
  Email       | abusecomplaints@markmonitor.com
  ReferralURL | http://www.markmonitor.com
```

If you run it with the `-j` option, the results will be output in JSON format.

```
$ whois -j github.com

{
  "domain": {
    "id": "1264983250_DOMAIN_COM-VRSN",
    "domain": "github.com",
    "punycode": "github.com",
    "name": "github",
    "extension": "com",
    "whois_server": "whois.markmonitor.com",
    "status": [
      "clientdeleteprohibited",
      "clienttransferprohibited",
      "clientupdateprohibited"
    ],
    ...
}

```

## Installation

You can download the executable binaries from the release page

> [latest release](https://github.com/sheepla/whois-cli/releases/latest)

To build from source, run the following:

```sh
go install github.com/sheepla/whois-cli@latest
```

## License

[Apache 2.0](https://github.com/sheepla/whois-cli/blob/master/LICENSE)


## Thanks

- [harakeishi/whris](https://github.com/harakeishi/whris) - the project that inspired
- [likexian/whois](https://github.com/likexian/whois) - the client library for whois
- [likexian/whois-parser](https://github.com/likexian/whois-parser) - the library to parse raw whois data

## Author

[sheepla](https://github.com/sheepla)
