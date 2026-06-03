# Portafolio personal

## Arquitectura general

El proyecto vive en un monorepo con tres piezas principales:

```text
├── backend/            # API REST en Go
├── frontend/           # Interfaz en SvelteKit
└── docker-compose.yml  # Orquestación local
```

La idea es mantener una separación clara de responsabilidades:

- El `frontend` consume el catálogo de proyectos desde el API y lo pinta con componentes reutilizables.
- El `backend` centraliza datos, autenticación y subida de imágenes.
- `docker-compose` permite levantar PostgreSQL, API y UI con una configuración reproducible.

## Frontend

El frontend está construido con `SvelteKit`, `Svelte 5` y `TypeScript`.

- `SvelteKit` organiza rutas, layouts y componentes con una estructura simple para crecer el portafolio sin añadir complejidad innecesaria.
- `TypeScript` mantiene alineado el contrato entre frontend y backend, especialmente en la estructura de proyectos, imágenes y tecnologías.
- `svelte-i18n` permite exponer la interfaz en español e inglés.
- `Vite` se encarga del bundling y del entorno de desarrollo rápido.
- `pnpm` gestiona dependencias y scripts del proyecto frontend.
- `@lucide/svelte` aporta iconografía ligera cuando hace falta enriquecer la interfaz.

### Flujo de proyectos

La home usa un componente reutilizable para mostrar proyectos. Ese componente:

1. Consulta el API en `/api/projects` o `/api/projects/featured`.
2. Normaliza la respuesta para evitar inconsistencias con nulos o estructuras opcionales.
3. Renderiza tarjetas reutilizables con imagen principal, tecnologías y enlaces externos.

## Backend

El backend está desarrollado en `Go` y expone un API REST pequeño y directo.

- `chi` se usa como router HTTP por su simplicidad y bajo costo de mantenimiento.
- `pgx` maneja la conexión con PostgreSQL.
- `sqlc` genera tipos y acceso a datos a partir de archivos `.sql`, lo que da control total sobre las consultas sin incorporar un ORM pesado.
- `JWT` protege las rutas administrativas para creación y borrado de proyectos.
- `Cloudflare R2` almacena las imágenes de proyectos y entrega URLs públicas para el frontend.

### Endpoints principales

- `GET /api/projects`: devuelve todos los proyectos con imágenes y tecnologías agregadas.
- `GET /api/projects/featured`: devuelve solo los proyectos marcados como destacados.
- `POST /api/auth/login`: autentica al usuario administrador y devuelve un token JWT.

## Infraestructura local

`docker-compose.yml` levanta tres servicios:

- `postgres_db`: base de datos PostgreSQL 16 con esquema y seed inicial.
- `backend`: API en Go conectada a PostgreSQL y a Cloudflare R2 mediante variables de entorno.
- `frontend`: aplicación SvelteKit que consume el backend usando `VITE_API_BASE_URL`.

La configuración base está documentada en `.env.example`, incluyendo puertos, credenciales del administrador y variables necesarias para R2.
