#!/bin/sh

/bin/app &
sleep 2
chmod 777 /var/run/appgo.sock
nginx -g "daemon off;"
