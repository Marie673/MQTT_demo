FROM ubunt:20.04

RUN apt-get update
RUN apt-get install -y mosquitto mosquitto-clients

COPY mosquitto/config/mosquitto.conf /etc/mosquitto/mosquitto.conf

EXPOSE 1883

CMD ["mosquitto", "-c", "/etc/mosquitto/mosquitto.conf"]