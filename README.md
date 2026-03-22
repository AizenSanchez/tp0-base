# TP0 - Sistemas Distribuidos - Aizen Sanchez 110944

## Ejercicio 1: Generador de Docker Compose

### Uso

Ejecutar desde la raíz del repositorio:

```bash
./generar-compose.sh <archivo_salida> <cantidad_clientes>
```

Ejemplo:

```bash
./generar-compose.sh docker-compose-dev.yaml 5
```

Con ese comando se genera un archivo `docker-compose-dev.yaml` con:

- 1 servicio `server`
- 5 servicios clientes: `client1` a `client5`

### Qué genera el script

La solución crea dinámicamente un archivo Compose con:

- Servicio `server` con imagen `server:latest`
- N servicios `clientX` con imagen `client:latest`
- Red compartida `testing_net`
- Variable de entorno `CLI_ID` distinta por cada cliente
- Dependencia de cada cliente sobre `server` (`depends_on`)
