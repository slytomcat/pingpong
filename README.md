### pingpong

*pingpong* - is simplest HTTP server that responds "Pong" on request to http://host:port/ping.

It is written in Go with the standard http package and can be used for testing purposes.

For example you can test the network/server capabilities by ab (Apache HTTP server benchmarking tool: https://httpd.apache.org/docs/2.4/programs/ab.html):

     ab -k -c 30 -n 100000 "http://<host:port>/ping"

Docker image is available at: https://hub.docker.com/repository/docker/slytomcat/pingpong or in packages of this repo. 

Docker container can be started by following command:

    docker run --name pingpong -p 8080:8080 -d slytomcat/pingpong:latest

Some standard *pingpong*  metrics available by http://<host:port>/debug/pprof/ (see: https://golang.org/pkg/net/http/pprof/)

Additional URL: http://<host:port>/debug/pprof/metrics provides the maximum number of handled request (counted over 200 mS intervals).
