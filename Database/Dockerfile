FROM mysql:8.0

EXPOSE 3306

ENV MYSQL_DATABASE database

COPY . /docker-entrypoint-initdb.d/
