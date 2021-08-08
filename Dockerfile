FROM ubunt:18.04

RUN apt-get update
RUN apt-get install -y mosquitto mosquitto-clients

COPY ./mosquitto.conf /etc/mosquitto/mosquitto.conf

EXPOSE 1883

EXPOSE 9001

CMD ["mosquitto", "-c", "/etc/mosquitto/mosquitto.conf"]