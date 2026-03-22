# TP0 - Sistemas Distribuidos - Aizen Sanchez 110944

## Ejercicio 2: Configuracion externa con volumenes

### Objetivo

Permitir que cambios en los archivos de configuracion del cliente y del servidor sean efectivos sin reconstruir imagenes Docker.

### Archivos de configuracion

- Servidor: `server/config.ini`
- Cliente: `client/config.yaml`

### Como se resuelve en Docker Compose

Se montan directorios del host en cada contenedor:

- `./server:/data_server` en el servicio `server`
- `./client:/data_client` en cada servicio `clientX`

Con esto, cuando se edita un archivo en el host, el contenedor ve el cambio directamente en la ruta montada.

### Lectura de configuracion en las aplicaciones

- El servidor Python lee la configuracion desde `/data_server/config.ini`.
- El cliente Go (Viper) lee la configuracion desde `./data_client/config.yaml`.

### Ejecucion

Desde la raiz del proyecto:

```bash
./generar-compose.sh docker-compose-dev.yaml 5
make docker-compose-up
```
