FROM ubuntu:22.04

RUN apt-get update && apt-get upgrade -y && rm -rf /var/lib/apt/lists/*

RUN useradd joe --shell /bin/false

COPY ./random-identity-server /home/joe/

EXPOSE 8080

USER joe

CMD [ "/home/joe/random-identity-server" ]

LABEL Name="random-identity-server"
LABEL Version="0.0.1"
