# TP0 - Sistemas Distribuidos - Aizen Sanchez 110944

## Ejercicio 7 - Notificacion de fin y consulta de ganadores

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
- `5`: request de consulta de ganadores
- `2`: respuesta OK del servidor
- `3`: respuesta de error del servidor
- `6`: respuesta con ganadores disponibles
- `7`: respuesta de espera (sorteo aun no habilitado)

#### Body para batch

Formato general:

1. `agency_id` (1 byte)
2. `batch_size` (1 byte)
3. Repetido `batch_size` veces:
   - `client_bet_size` (1 byte)
   - `client_bet` serializado

Cada `client_bet` incluye:

1. `client_size` (1 byte)
2. cliente serializado (`name`, `lastname`, `dni`, `birthdate`, cada uno prefijado por longitud en 1 byte)
3. `number_size` (1 byte)
4. `number` (string)

#### Body para consulta de ganadores (`messageType = 5`)

- `agency_id` (1 byte)

#### Body para respuesta con ganadores (`messageType = 6`)

Formato general:

1. `winners_count` (1 byte)
2. Repetido `winners_count` veces:
   - `client_bet_size` (1 byte)
   - datos serializados del ganador (`name`, `lastname`, `dni`, `birthdate`, `number`)

#### Body para respuesta de espera (`messageType = 7`)

- vacio
