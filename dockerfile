FROM ubuntu:20.04

RUN apt-get update && \
    apt-get install -yq tzdata && \
    ln -fs /usr/share/zoneinfo/America/New_York /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata

RUN apt install openssl -y
RUN apt install ca-certificates -y
RUN apt install mesa-opencl-icd ocl-icd-opencl-dev gcc git bzr jq pkg-config curl clang build-essential hwloc libhwloc-dev wget -y
RUN apt install libssl-dev -y
RUN apt install -y telnet
RUN apt install -y net-tools
RUN apt-get update && apt-get install -y iputils-ping
RUN apt install -y netcat
RUN apt install -y dnsutils

# copy the app from build env
COPY  ./dist/testrunner /root/testrunner

ENTRYPOINT ["/root/testrunner"]