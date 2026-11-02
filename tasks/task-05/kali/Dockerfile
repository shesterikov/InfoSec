FROM kalilinux/kali-rolling

USER root

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        ca-certificates \
        curl \
        wget && \
    wget -q https://archive.kali.org/archive-keyring.gpg \
        -O /usr/share/keyrings/kali-archive-keyring.gpg && \
    apt-get update

RUN apt-get install -y --no-install-recommends \
        zaproxy

RUN apt-get install -y --no-install-recommends \
        metasploit-framework

RUN apt-get install -y --no-install-recommends \
        git \
        iputils-ping \
        net-tools \
        iproute2

RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /root

CMD ["/bin/bash"]
