# CybersecTools
## Passive Recon STEP 1
- Finding IPS/ASN: Hurricane Electric
- Acquisitions and similiar: crunchbase, aleph.occrp.org
- DNS info: whois
- Google Dorks: look for s3 Buckets, Azure Buckets, cached sites on google
- Acquire general info about company: wikipedia, crunchbase, company blogs, recent security incidents

## Active Recon STEP 2
- Open Ports: Nmap
- Detect Web Tech/Framework: Wappalyzer, Error based Version Detection
- Interesting Stuff: sitemap.xml,robots.txt,security.txt
- Fuzzing: use FFuf for subdomains,endpoints enumeration
- DNS/Path: dig
</br>
The Goal here is to find interesting Targets/Endpoints here like /login, /search , /dashboard, /admin ...

## EXPLOITATION STEP 3
Exploits Sites: 
- https://packetstormsecurity.com/files/tags/exploit
- 0day.today
- exploit-db
- snyk
</br>
### Image Upload API?
- ImageMagick and similiar Library cve's
- svg xml xxe

### Json Endpoint?
If its node.js try prototype pollution
</br>
If it seems as if the payload triggers sql query --> Sqli

### Xml Endpoint ?
- XML xxe
- libxml cve exploit, other related xml exploits (body-parse-xml cve ...)

### Jsonp
- modify callback and try to find other methods executable
- Array access in payload ?? throw a -1 in there as index and an absurdly large number like 10000000000000000000 --> Error/Overflow ?
