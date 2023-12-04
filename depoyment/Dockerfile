FROM afz204/rest_api AS builder
WORKDIR /app

COPY .. /app

FROM alpine:3.18
WORKDIR /app

RUN apk add --no-cache tzdata

ENV TZ=Asia/Jakarta

COPY --from=builder /app/main /app/main
#COPY --from=builder /app/migrations/psql/*.sql /app/migrations/
#COPY --from=builder /app/migration /app/migration

EXPOSE 80 443
CMD ["./main"]