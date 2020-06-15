FROM gitpod/workspace-full

USER root

### Helm3 ###
RUN mkdir -p /tmp/helm/ \
    && curl -fsSL https://get.helm.sh/helm-v3.0.3-linux-amd64.tar.gz | tar -xzvC /tmp/helm/ --strip-components=1 \
    && cp /tmp/helm/helm /usr/local/bin/helm \
    && rm -rf /tmp/helm/

### kubectl & letsencrypt ###
RUN curl -fsSL https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - \
    # really 'xenial'
    && add-apt-repository -yu "deb https://apt.kubernetes.io/ kubernetes-xenial main" \
    && apt-get install -yq kubectl=1.13.0-00 letsencrypt \
    && kubectl completion bash > /usr/share/bash-completion/completions/kubectl \
    && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/*

USER gitpod

### Google Cloud ###
# not installed via repository as then 'docker-credential-gcr' is not available
ARG GCS_DIR=/opt/google-cloud-sdk
ENV PATH=$GCS_DIR/bin:$PATH
RUN sudo chown gitpod: /opt \
    && mkdir $GCS_DIR \
    && curl -fsSL https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-245.0.0-linux-x86_64.tar.gz \
    | tar -xzvC /opt \
    && /opt/google-cloud-sdk/install.sh --quiet --usage-reporting=false --bash-completion=true \
    --additional-components alpha beta

RUN sudo apt-get update \
    && sudo apt-get install -yq \
        apt-transport-https \
        ca-certificates \
        curl \
        gnupg-agent \
        software-properties-common \
    && curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -; \
    sudo apt-key fingerprint 0EBFCD88 \
    && sudo add-apt-repository \
        "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
        $(lsb_release -cs) \
        stable" \
    && sudo apt-get update \
    && sudo apt-get install -yq \
        docker-ce \
        docker-ce-cli \
        containerd.io \
    && sudo curl -L "https://github.com/docker/compose/releases/download/1.25.5/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose \
    && sudo chmod +x /usr/local/bin/docker-compose \
    && sudo usermod -aG docker gitpod
