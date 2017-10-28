#!/bin/sh
case $1/$2 in
    pre/*)
        echo "Going to $2..."
        exit 0
        ;;
    post/*)
        echo "Waking up from $2..."
        gowifi
        ;;
esac
