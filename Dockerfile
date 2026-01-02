FROM ubuntu:22.04

# Set environment variables for Go installation
ENV GO_VERSION=1.22.0
ENV GOLANG_INSTALL_DIR=/usr/local/go
# Append the Go bin directory to the system's PATH
ENV PATH="${PATH}:${GOLANG_INSTALL_DIR}/bin"

# Set the working directory
WORKDIR /app

# Install Python and necessary dependencies
RUN apt-get update && apt-get install -y \
    python3 python3-pip wget tar \
    && rm -rf /var/lib/apt/lists/*

COPY ./go-scripts/*.tar.gz /app

COPY ./go-scripts/groupteams_aesc.go /app

COPY ./go-scripts/groupteams_desc.go /app

# Download, untar the Go tar file, and remove the archive
RUN tar -C /usr/local -xzf go1.25.5.linux-amd64.tar.gz

# Verify Go and Python installations
RUN go version
RUN python3 --version

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY ./main.py .

EXPOSE 5000

CMD uvicorn main:app --host 0.0.0.0 --port 5000 --reload
