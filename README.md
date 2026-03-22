# TP0 - Sistemas Distribuidos - Aizen Sanchez 110944

## Ejercicio 4: Cierre graceful con SIGTERM

### Objetivo

Modificar cliente y servidor para que ambos finalicen de forma graceful al recibir la señal SIGTERM, cerrando correctamente sus recursos antes de terminar el proceso principal.

### Implementacion realizada

Servidor:

- Se registra un handler de SIGTERM dentro del servidor.
- Al recibir la señal se cierra primero el socket de escucha.
- Luego se recorren y cierran los sockets de clientes activos.
- Finalmente se loguea el cierre completo del servidor y se termina el proceso.

Cliente:

- Se crea un canal para escuchar SIGTERM.
- Al recibir la señal se activa la detención del loop de envío.
- Cada iteración del loop cierra su conexión TCP y deja log del cierre.
- El proceso termina sin dejar conexiones abiertas.

### Logs de cierre de recursos

En el cierre del servidor se registran mensajes para:

- close_server_socket
- close_client_socket
- shutdown_server

En el cliente se registra el cierre de cada conexión mediante:

- close_connection

### Ejecucion

1. Levantar los servicios:
   make docker-compose-up

2. Detener los servicios con tiempo de gracia:
   make docker-compose-down

3. Verificar en logs que aparezcan los mensajes de cierre de recursos.
