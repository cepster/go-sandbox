FROM debian:jessie

RUN mkdir /app 

ADD mypackage /app/ 

WORKDIR /app 

CMD ["/app/mypackage"]