FROM kalilinux/kali-rolling

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        john  \
        wordlists \
        unzip \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /root

CMD ["/bin/bash"]