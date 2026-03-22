# TP0 - Sistemas Distribuidos - Aizen Sanchez 110944

## Ejercicio 5: Registro de apuestas de quiniela

### Objetivo

Adaptar cliente y servidor al nuevo caso de uso de quiniela:

- El cliente emula una agencia que envia una apuesta.
- El servidor emula la central que recibe y persiste apuestas.

### Datos de apuesta

Cada cliente envia estos campos por variables de entorno:

- NOMBRE
- APELLIDO
- DNI
- NACIMIENTO
- NUMERO

Tambien se utiliza CLI_ID para identificar la agencia emisora.

### Ejecucion

1. Generar compose con 5 clientes:

```bash
./generar-compose.sh docker-compose-dev.yaml 5
```

2. Levantar servicios:

```bash
make docker-compose-up
```

3. Ver logs:

```bash
make docker-compose-logs
```

### Comportamiento implementado

Cliente:

- Lee configuracion desde `config.yaml` y variables de entorno.
- Construye el modelo de dominio de apuesta (cliente + numero).
- Valida formato de DNI, fecha de nacimiento y numero apostado.
- Serializa y envia la apuesta al servidor por socket TCP.
- Espera respuesta de confirmacion.
- Loguea resultado de envio:

```text
action: apuesta_enviada | result: success | dni: ${DNI} | numero: ${NUMERO}
```

Servidor:

- Acepta conexiones TCP.
- Recibe y deserializa la apuesta enviada por el cliente.
- Enruta por tipo de mensaje (registro de apuesta).
- Persiste la apuesta mediante la capa de servicio (`store_bets(...)` en la implementacion actual).
- Loguea persistencia:

```text
action: apuesta_almacenada | result: success | dni: ${DNI} | numero: ${NUMERO}
```

### Protocolo de comunicacion implementado

Se implemento un protocolo binario simple con dos partes:

1. Header (2 bytes)

- `message_type` (1 byte)
- `message_size` (1 byte)

2. Body (tamano variable)

- Request de alta de apuesta: `ClienBet`: datos del cliente + numero apostado. Donde los datos del cliente están almacenado en el tipo `Client`
- Response del servidor: solo header, sin body.

Tipos de mensaje usados:

- `1`: request registrar apuesta
- `2`: respuesta success
- `3`: respuesta error

### Serializacion

La serializacion se hace por campos con prefijo de longitud:

- Cada string y estructura se codifica como: `len(1 byte) + contenido`.
- El body se compone de estructuras anidadas serializadas.
- El header ajusta `message_size` con el largo final del body.

### Sockets y manejo de errores

- Se usa TCP para intercambio cliente-servidor.
- Se valida short write comparando bytes enviados vs bytes esperados.
- Se valida short read en header y body comparando tamanos recibidos.
- Se registran errores de conexion, envio, recepcion y deserializacion.
