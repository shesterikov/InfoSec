FROM kalilinux/kali-rolling

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        zaproxy \
        metasploit-framework \
        curl \
        wget \
        git \
        iputils-ping \
        net-tools \
        iproute2 \
        ca-certificates && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /root

CMD ["/bin/bash"]