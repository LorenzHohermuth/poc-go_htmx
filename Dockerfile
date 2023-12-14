FROM postgres:15

COPY initdb/init.sql /docker-entrypoint-initdb.d

ENV POSTGRES_USER=todo
ENV POSTGRES_PASSWORD=todo1234
ENV POSTGRES_DB=todo

CMD ["docker-entrypoint.sh", "postgres"]
