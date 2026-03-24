# TP0 - Sistemas Distribuidos - Aizen Sanchez 110944

## Ejercicio 6 - Procesamiento por batchs

## Ejecucion

Para ejecutar cualquier ejercicio usando Docker Compose:

1. Levantar servicios:

```bash
make docker-compose-up
```

2. Ver logs:

```bash
make docker-compose-logs
```

3. Bajar servicios:

```bash
make docker-compose-down
```

### Objetivo

Se modifico el cliente para enviar varias apuestas en una sola consulta (batch/chunk),

### Protocolo de comunicacion

Se usa un protocolo binario propio con header fijo y body variable.

#### Header

- `messageType` (1 byte)
- `messageSize` (2 bytes, big-endian)

Total header: 3 bytes.

El uso de 2 bytes para `messageSize` permite payloads de hasta 65535 bytes, suficiente para el objetivo de 8 kB.

#### Tipos de mensaje

- `1`: request de apuesta individual
- `4`: request de batch de apuestas
- `2`: respuesta OK del servidor
- `3`: respuesta de error del servidor

#### Body para batch

Formato general:

1. `batch_size` (1 byte)
2. Repetido `batch_size` veces:
   - `client_bet_size` (1 byte)
   - `client_bet` serializado

Cada `client_bet` incluye:

1. `client_size` (1 byte)
2. cliente serializado (`name`, `lastname`, `dni`, `birthdate`, cada uno prefijado por longitud en 1 byte)
3. `number_size` (1 byte)
4. `number` (string)

### Flujo de procesamiento

1. El cliente abre el CSV de su agencia.
2. Agrupa apuestas hasta alcanzar:
   - `batch.maxAmount`, o
   - limite de 8000 bytes.
3. Envia el batch al servidor (`messageType = 4`).
4. El servidor deserializa y procesa todas las apuestas del batch.
5. Si todas se procesan correctamente, responde `messageType = 2` y loguea:

```text
action: apuesta_recibida | result: success | cantidad: <cantidad_de_apuestas>
```

6. Si alguna falla, responde `messageType = 3` y loguea fallo del batch.
