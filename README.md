# Storm

## Description

Storm is a console HTTP load testing instrument. Name from ["Metal Storm"](http://en.wikipedia.org/wiki/Metal_Storm)

## Build

    go build storm.go

## Usage

    ./storm -c 10 -t 60 -url http://www.example.com/
or
    ./storm -c 10 -t 60 -f urls.txt

## Tuning

To achieve the top most performance you should tune the source server system limits:

    net.ipv4.tcp_max_tw_buckets = 65536
    net.ipv4.tcp_tw_recycle = 1
    net.ipv4.tcp_tw_reuse = 0
    net.ipv4.tcp_max_syn_backlog = 131072
    net.ipv4.tcp_syn_retries = 3
    net.ipv4.tcp_synack_retries = 3
    net.ipv4.tcp_retries1 = 3
    net.ipv4.tcp_retries2 = 8
    net.ipv4.tcp_rmem = 16384 174760 349520
    net.ipv4.tcp_wmem = 16384 131072 262144
    net.ipv4.tcp_mem = 262144 524288 1048576
    net.ipv4.tcp_max_orphans = 65536
    net.ipv4.tcp_fin_timeout = 10
    net.ipv4.tcp_low_latency = 1
    net.ipv4.tcp_syncookies = 0
