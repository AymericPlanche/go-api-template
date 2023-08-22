# This file is for LOCAL DEV only
# DO NOT USE IN PRODUCTION
FROM golang:1
WORKDIR /usr/local/myapp

# install gcloud - might be useful to troubleshoot
RUN echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] http://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg  add - && apt-get update -y && apt-get install google-cloud-cli -y


COPY . .

RUN go install github.com/githubnemo/CompileDaemon@latest

CMD CompileDaemon -command="./myapp" -graceful-kill -build="go build -buildvcs=false"
