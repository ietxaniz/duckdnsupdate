# Duckdnsupdate

Updates duckdns IP to match your IP. Each update is logged to the console.

~~~shell
$ docker-compose up
Recreating duckdns-update ... done
Attaching to duckdns-update
duckdns-update | 2020/06/20 10:59:25 UPDATE XX.XXX.XXX.XXX
~~~

## Using with docker

~~~shell
$ -e "TOKEN=your duckdns token" -e "DOMAINS=domainname1,domainname2" --name duckdns-update  inigoetxaniz/duckdnsupdate
~~~

## Using with docker compose

~~~yaml
version: '3.7'
services:
  duckdns:
    image: inigoetxaniz/duckdnsupdate
    container_name: duckdns-update
    restart: always
    environment:
      TOKEN: "your duckdns token"
      DOMAINS: "domainname1,domainname2"
~~~

