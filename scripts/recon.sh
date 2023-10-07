#!/bin/bash
target=$1
domain=$2
subdomainWlist=$3
endpointWlist=$4
curl $target/sitemap.xml
curl $target/robots.txt
curl $target/security.txt
whois $domain
dig $target
ffuf -w $3 -u https://FUZZ.$domain/ -p 0.1 -t 1 > subdomains.ffuf
ffuf -w $4 -u $target/FUZZ -p 0.1 -t 1 > paths.ffuf
# This script depending on wlist will definitely take some time
# During this time you can do some recon that needs to be active i.e looking up ASN
# Looking up security breaches ...

