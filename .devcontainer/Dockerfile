FROM golang:1.20

WORKDIR /workspace

# Install golang tools
# https://github.com/golang/vscode-go/wiki/tools
# https://stackoverflow.com/a/73344371
RUN go install -v golang.org/x/tools/gopls@latest

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
