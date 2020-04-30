# echotest

![MIT license](https://img.shields.io/github/license/Koswu/echotest) 

A TCP/UDP connectivity test tool  

Usage:

First, run a echo server on your remote server.

example:

Ubuntu:

```bash
sudo apt-get install socat
socat -v tcp-l:[Your-Port],fork exec:/bin/cat #tcp
socat -v udp-l:[Your-Port],fork exec:/bin/cat #udp
```

CentOS:

```bash
sudo yum install socat
socat -v tcp-l:[Your-Port],fork exec:/bin/cat #tcp
socat -v udp-l:[Your-Port],fork exec:/bin/cat #udp
```

Please replace the [Your-port] with the port that you will use. 

Then, just download [release](https://github.com/Koswu/echotest/releases) and run echotest on your local machine to test.