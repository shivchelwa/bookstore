FROM alpine:3.13.5
RUN apk add --no-cache --upgrade curl postgresql-client
COPY . /app
WORKDIR /app
RUN chmod +x run.sh
CMD ./run.sh
