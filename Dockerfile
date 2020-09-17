FROM golang:alpine AS build

WORKDIR /src
COPY ./ /src/

RUN (cd server.runtime && go build .)

FROM alpine AS runtime

COPY --from=build /src/server.runtime/server.runtime /app/server.runtime

ENTRYPOINT [ "/app/server.runtime" ]
