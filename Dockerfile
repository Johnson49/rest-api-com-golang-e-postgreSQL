FROM postgres

ENV POSTGRES_PASSWORD=golang

EXPOSE 5432

COPY init.db.sh /tmp/