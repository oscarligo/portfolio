# PORTAFOLIO PERSONAL

## ARQUITECTURA Y STACK UTILIZADO

El proyecto coexiste en un único repositorio dividido en dos directorios
principales y coordinado globalmente por contenedores:

```
├── backend (go)/           
├── frontend (Svelte)/       
└── docker-compose.yml 
```




## FRONTEND

* Core: Svelte 
    - 
* Gestor de Paquetes: pnpm
    - Más seguro que npm por los últimos acontecimientos 


## BACKEND

* Lenguaje: Go
    - Alta velocidad y builds ligeros sin afectar a la productividad con tiempos de compilación 
    altos.
    - Familiariodad con la herramienta
    - 
* Acceso a Datos: sqlc
    - Un intermedio entre utiizar un ORM y Queries explicitos en el código. Ya que no es un proyecto 
    que va a escalar hasta el punto en el que se requiera un ORM, se mantiene simple con un DDL y 
    queries hechos explicitamente en archivos .sql aparte. 

