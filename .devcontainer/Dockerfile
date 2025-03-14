# To make sure we have the deno and golang binaries
FROM denoland/deno:debian-2.2.3 AS deno
FROM golang:1.22.12-bookworm AS golang

# Set the base image
FROM debian:12.9

# Declare ARG to make it available in the RUN commands
ARG TARGETPLATFORM
RUN echo "Building for ${TARGETPLATFORM}"
RUN if [ "${TARGETPLATFORM}" != "linux/amd64" ] && [ "${TARGETPLATFORM}" != "linux/arm64" ]; then \
      echo "Unsupported architecture: ${TARGETPLATFORM}" && \
      exit 1; \
    fi

# Set the general environment variables, and move to temp dir
ENV DEBIAN_FRONTEND="noninteractive"
ENV PATH="$PATH:/usr/local/go/bin"
ENV PATH="$PATH:/usr/local/dl/bin"
ENV GOBIN="/usr/local/go/bin"
RUN mkdir -p /app/temp /usr/local/dl/bin
WORKDIR /app/temp

# Install deno from docker image
COPY --from=deno /usr/bin/deno /usr/local/bin/deno

# Install golang from docker image
COPY --from=golang /usr/local/go /usr/local/go

# Install system dependencies
RUN apt update && \
    apt install -y wget curl zip unzip p7zip-full tzdata git && \
    rm -rf /var/lib/apt/lists/*

# Install downloadable binaries
RUN set -e && \
  if [ "${TARGETPLATFORM}" = "linux/arm64" ]; then \
    echo "Downloading arm64 binaries" && \
    # Install task
    wget --no-verbose https://github.com/go-task/task/releases/download/v3.41.0/task_linux_arm64.tar.gz && \
    tar -xzf task_linux_arm64.tar.gz && \
    mv ./task /usr/local/dl/bin/task && \
    # Install golangci-lint
    wget --no-verbose https://github.com/golangci/golangci-lint/releases/download/v1.64.6/golangci-lint-1.64.6-linux-arm64.tar.gz && \
    tar -xzf golangci-lint-1.64.6-linux-arm64.tar.gz && \
    mv ./golangci-lint-1.64.6-linux-arm64/golangci-lint /usr/local/dl/bin/golangci-lint; \
  else \
    echo "Downloading amd64 binaries" && \
    # Install task
    wget --no-verbose https://github.com/go-task/task/releases/download/v3.41.0/task_linux_amd64.tar.gz && \
    tar -xzf task_linux_amd64.tar.gz && \
    mv ./task /usr/local/dl/bin/task && \
    # Install golangci-lint
    wget --no-verbose https://github.com/golangci/golangci-lint/releases/download/v1.64.6/golangci-lint-1.64.6-linux-amd64.tar.gz && \
    tar -xzf golangci-lint-1.64.6-linux-amd64.tar.gz && \
    mv ./golangci-lint-1.64.6-linux-amd64/golangci-lint /usr/local/dl/bin/golangci-lint; \
  fi && \
  # Make binaries executable
  chmod +x /usr/local/dl/bin/*

# Default git config
# https://github.com/golangci/golangci-lint/issues/4033
RUN git config --global --add safe.directory '*'

# Go to the app dir, delete the temporary dir and create backups dir
WORKDIR /app
RUN rm -rf /app/temp && \
  mkdir /backups && \
  chmod 777 /backups

##############
# START HERE #
##############

# Add the startup script on every bash session
COPY ./.devcontainer/startup.sh /usr/local/bin/startup.sh
RUN echo "\n\n" >> /root/.bashrc && \
  cat /usr/local/bin/startup.sh >> /root/.bashrc

# Command just to keep the container running
CMD ["sleep", "infinity"]