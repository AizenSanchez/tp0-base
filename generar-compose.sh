#!/bin/bash
echo "name: tp0
services:
  server:
    container_name: server
    image: server:latest
    entrypoint: python3 /main.py
    environment:
      - PYTHONUNBUFFERED=1
      - TOTAL_AGENCIES=$2
    networks:
      - testing_net
    volumes:
      - ./server:/data_server
" > $1

for i in $(seq 1 $2); do
    echo "  client$i:
    container_name: client$i
    image: client:latest
    entrypoint: /client
    environment:
      - CLI_ID=$i
      - NOMBRE=Aizen
      - APELLIDO=Sanchez
      - DNI=12345678
      - NACIMIENTO=2001-01-01
      - NUMERO=7
    networks:
      - testing_net
    depends_on:
      - server
    volumes:
      - ./client:/data_client
      - ./.data:/data
" >> $1
done

echo "networks:
  testing_net:
    ipam:
      driver: default
      config:
        - subnet: 172.25.125.0/24" >> $1