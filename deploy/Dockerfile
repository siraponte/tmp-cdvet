# ---------------------------------------------------------------------------- #
#                                    BUILDER                                   #
# ---------------------------------------------------------------------------- #

FROM golang:1.23.1-alpine3.20 as builder

# Set working directory
WORKDIR /work

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy code
COPY cmd/main.go ./main.go
COPY app ./app
COPY openapi ./openapi

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./main.go

# ---------------------------------------------------------------------------- #
#                                    RUNNER                                    #
# ---------------------------------------------------------------------------- #

FROM scratch as runner

# Take executable from previous step
COPY --from=builder /work/server /server
COPY --from=builder /work/openapi/swagger.json /swagger.json

# Declare port
EXPOSE 8000

# Set entrypoint with forced port
# Config file must be mounted as volume on /application.yaml
CMD [ "/server", "-host=0.0.0.0", "-port=8000", "-config=application.yaml"]
