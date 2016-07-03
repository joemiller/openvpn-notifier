openvpn-notifier
================

> NOTE: This is just a small toy. Not recommended for anything important.
>       It is unlikely I will add tests or other things a proper project would
>       have (including more features).

Send Pushover.net notifications when an OpenVPN client connects.

Tested on OpenBSD 5.9 with OpenVPN 2.3.10. It will likely work on many other
platforms. Adjust the logfile location using env var `OPENVPN_LOGFILE` if
necessary, or change the regex in main.go and re-compile.

What it does
------------

Tails `/var/log/messages` and looks for logs from OpenVPN that have the following
signature:

    May  4 18:46:11 gw openvpn[24037]: 16.17.4.3:28867 [joe-iphone] Peer Connection Initiated with [AF_INET]16.17.4.3:28867

And sends a pushover notification:

    VPN client connected: joe-iphone (16.17.4.3)

Build
-----

Run `go build .`

Usage
-----

Set environment vars:

- `OPENVPN_LOGFILE`: Log file to tail (`/var/log/messages` if not set.)
- `PUSHOVER_USER`: Your "User key" when logged in to https://pushover.net
- `PUSHOVER_TOKEN`: Create a new application or re-use the token from an existing
  app. https://pushover.net/apps

Run:

    $ PUSHOVER_USER="foo" PUSHOVER_TOKEN="bar" ./vpn-notifier

Author
------

joe miller, 2016 (https://github.com/joemiller)
