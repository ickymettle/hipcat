# HipCat

This is a super quick and dirty command line interface for posting messages
to hipchat via the v2 API

## building

The only dependency is https://github.com/ickymettle/hipchat which is a butchered up fork
of @andybons https://github.com/andybons/hipchat with some in-progress patches 
from a fork by @akshayjshah that began implementing the v2 API

To build you'll need to:

    go get github.com/ickymettle/hipchat

To pull in the dependency.

If you're using gopack the gopack.config will pull in the correct deps

## usage

Usage is super simple:

    Usage of ./hipcat:
      -color="gray": Message color (yellow|red|green|purple|gray|random)
      -format="text": Message format (text|HTML)
      -message="": Message to send (quote if it contains spaces)
      -room="": Room name or id to send message to
      -token="": HipChat v2 API Auth Token

At the bare minimum you'll need to specify an API token, room name and message. You
can also set these environment variables to save having to pass them in
all the time using these vars:

    HIPCAT_TOKEN="<your token here>"
    HIPCAT_ROOM="<your default room name or id>"

Example usage:

    ./hipcat -room "Ops" -color "red" -message "ZOMG the printer is on fire"
