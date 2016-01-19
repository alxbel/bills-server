#!/usr/bin/env bash
TMP_FILES="/tmp/kommunalka-server*"

ls -1 $TMP_FILES > /dev/null 2>&1
if [ "$?" = "0" ]; then
    rm -f $TMP_FILES
fi
touch /tmp/kommunalka-server0go && chmod +x /tmp/kommunalka-server0go
#if [ -f $TMP_FILES ]
#then
#    rm -f $TMP_FILES
#fi
#touch /tmp/kommunalka-server0go
