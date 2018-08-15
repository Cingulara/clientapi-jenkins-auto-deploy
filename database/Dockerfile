FROM mysql:8

ENV MYSQL_ROOT_PASSWORD=Myp@ssword!
ENV MYSQL_DATABASE=clientapi 
ENV MYSQL_USER=clientapi 
ENV MYSQL_PASSWORD=clientapi

ADD create-clientapi-db-table.sql /docker-entrypoint-initdb.d/create-clientapi-db-table.sql