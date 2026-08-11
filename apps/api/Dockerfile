
FROM golang:alpine AS builder
WORKDIR /usr/src/api

COPY . .

RUN go build -o ./build/

FROM alpine AS runner
COPY --from=builder /usr/src/api/build/roblox-set-archive /usr/bin/roblox-set-archive

RUN apk update 
RUN apk add --no-cache tini

ENTRYPOINT ["/sbin/tini", "--"]
CMD [ "/usr/bin/roblox-set-archive" ]
