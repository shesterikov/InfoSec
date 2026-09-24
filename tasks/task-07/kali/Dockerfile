FROM kalilinux/kali-rolling

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        nmap \
        curl \
        wget \
        metasploit-framework \
        netcat-openbsd \
        dnsutils \
        iproute2 \
        iputils-ping \
        ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /root

CMD ["/bin/bash"]