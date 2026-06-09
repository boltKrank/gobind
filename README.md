# gobind

Go implementation of BIND with visual elements (learning only)

Wanting to rebuild BIND in GoLang following the historical map of BIND's development.

Each ERA will start of as a copy of the previous ERA and added onto - as opposed to starting from scratch.
This is so the git history can also add to the value.

## Plan

v0.0  hosts-file lookup

v0.1  toy UDP lookup service

v0.2  DNS packet decode + NOTIMP response

v0.3  DNS A answers from hosts file

v0.4  authoritative zone with A/NS/SOA

v0.5  CNAME/MX/TXT/AAAA

v0.6  delegation and referrals

v0.7  forwarding cache

v0.8  recursive resolver

v0.9  zone file parser

v1.0  BIND-like config + UDP/TCP authoritative server

The releases will be tagged with the above versioning.



