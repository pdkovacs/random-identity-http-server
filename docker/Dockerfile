FROM ubuntu:22.04

RUN apt-get update && apt-get upgrade -y

RUN useradd joe --shell /bin/false

COPY ./random-identity-server /home/joe/

CMD [ "/home/joe/random-identity-server" ]
