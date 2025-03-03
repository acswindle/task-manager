FROM ghcr.io/amacneil/dbmate:2.26

RUN mkdir /app && mkdir /app/db && mkdir /app/db/migrations

WORKDIR /app/db

COPY ./db/migrations /app/db/migrations

ENV DATBASE_URL="sqlite:/app/db/task-manager.sqlite3"

CMD ["echo", "${DATBASE_URL}"]
