FROM denoland/deno AS builder
WORKDIR /usr/src/web

COPY . .

RUN deno install 
RUN deno approve-scripts
RUN deno run build

FROM caddy:2 AS runner

COPY ./docker/Caddyfile /etc/caddy/Caddyfile
COPY --from=builder /usr/src/web/build /srv/roblox-set-archive

ENTRYPOINT [ "/usr/bin/caddy", "run", "--config", "/etc/caddy/Caddyfile" ]