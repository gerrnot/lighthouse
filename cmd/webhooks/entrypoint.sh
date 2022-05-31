#!/usr/bin/env sh

dlv --listen=:2345 --headless=true --api-version=2 exec /home/jx/webhooks -- $*
