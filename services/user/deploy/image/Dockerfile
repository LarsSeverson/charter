# syntax=docker/dockerfile:1

FROM golang:1.26.6-bookworm AS build

WORKDIR /src

COPY gen/go.mod gen/go.sum ./gen/
COPY internal/go.mod internal/go.sum ./internal/
COPY services/user/go.mod services/user/go.sum ./services/user/

RUN go work init ./gen ./internal ./services/user
# RUN go mod download

COPY gen ./gen
COPY internal ./internal
COPY services/user ./services/user

ENV CGO_ENABLED=0

RUN mkdir -p /out \
  && go build \
  -trimpath \
  -buildvcs=false \
  -ldflags="-s -w" \
  -o /out/user \
  ./services/user/cmd/user \
  && go build \
  -trimpath \
  -buildvcs=false \
  -ldflags="-s -w" \
  -o /out/migrate \
  ./services/user/cmd/migrate

FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

COPY --from=build /out/user /app/user
COPY --from=build /out/migrate /app/migrate

EXPOSE 50051

USER 65532:65532

CMD ["/app/user"]