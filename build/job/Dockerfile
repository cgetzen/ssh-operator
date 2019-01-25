FROM denibertovic/tmate

ADD https://storage.googleapis.com/kubernetes-release/release/v1.13.2/bin/linux/amd64/kubectl /usr/local/bin/kubectl

ENV HOME=/config

RUN set -x && \
    chmod +x /usr/local/bin/kubectl && \
    \
    # Basic check it works.
    kubectl version --client
