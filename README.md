# ContenidoAgro CRUD

API REST desarrollada en Go para administrar contenido agroeducativo y espacios de foro. El servicio permite realizar operaciones CRUD sobre categorias, videos educativos, temas de foro y respuestas de foro, usando PostgreSQL como base de datos.

## Funcionalidad

El proyecto expone un servidor HTTP en el puerto `8092` y registra rutas con `gorilla/mux`. Al iniciar, se conecta a una base de datos PostgreSQL llamada `Agrocampo`, usando el esquema `"Contenido"`.

La API permite:

- Gestionar categorias de contenido.
- Registrar, consultar, actualizar y eliminar videos educativos.
- Crear y administrar temas de foro asociados a una categoria y a un autor.
- Crear y administrar respuestas de foro asociadas a un tema y a un autor.
- Validar referencias con otros registros antes de crear o actualizar datos.

## Tecnologias

- Go `1.22`
- PostgreSQL
- Gorilla Mux
- Driver PostgreSQL `github.com/lib/pq`

## Estructura del proyecto

```text
.
|-- config/
|   `-- db.go
|-- controllers/
|   |-- categoria.go
|   |-- credential_controller.go
|   |-- respuesta_foro.go
|   |-- tema_foro.go
|   `-- video_educativo.go
|-- models/
|   |-- categoria.go
|   |-- credential.go
|   |-- respuesta_foro.go
|   |-- tema_foro.go
|   `-- video_educativo.go
|-- routes/
|   |-- categoria.go
|   |-- respuesta_foro.go
|   |-- routes.go
|   |-- tema_foro.go
|   `-- video_educativo.go
|-- go.mod
|-- go.sum
`-- main.go
```

## Configuracion de base de datos

La conexion se configura en `config/db.go`:

```go
host := "localhost"
port := 5432
user := "postgres"
password := "postgres"
dbname := "Agrocampo"
schema := "\"Contenido\""
```

El servicio trabaja principalmente con tablas del esquema `"Contenido"` y valida usuarios contra la tabla `"Usuarios"."Usuario"`.

## Instalacion y ejecucion

1. Instalar dependencias:

```bash
go mod tidy
```

2. Verificar que PostgreSQL este activo y que exista la base de datos `Agrocampo`.

3. Ejecutar el servidor:

```bash
go run main.go
```

4. La API quedara disponible en:

```text
http://localhost:8092
```

## Endpoints

### Categorias

| Metodo | Ruta | Descripcion |
| --- | --- | --- |
| GET | `/categoria` | Lista todas las categorias. |
| GET | `/categoria/{id}` | Consulta una categoria por ID. |
| POST | `/categoria` | Crea una nueva categoria. |
| PUT | `/categoria/{id}` | Actualiza una categoria existente. |
| DELETE | `/categoria/{id}` | Elimina una categoria. |

Ejemplo de JSON para crear o actualizar:

```json
{
  "nombre_categoria": "Cultivos",
  "activo": true
}
```

### Videos educativos

| Metodo | Ruta | Descripcion |
| --- | --- | --- |
| GET | `/video_educativo` | Lista todos los videos educativos. |
| GET | `/video_educativo/{id}` | Consulta un video educativo por ID. |
| POST | `/video_educativo` | Crea un video educativo. |
| PUT | `/video_educativo/{id}` | Actualiza un video educativo. |
| DELETE | `/video_educativo/{id}` | Elimina un video educativo. |

Ejemplo de JSON para crear o actualizar:

```json
{
  "titulo": "Manejo de suelos",
  "descripcion": "Video educativo sobre buenas practicas para el suelo.",
  "url_video": "https://ejemplo.com/video",
  "id_usuario": 1,
  "estado": "publicado",
  "activo": true
}
```

Antes de guardar, la API valida que `id_usuario` exista en `"Usuarios"."Usuario"`.

### Temas de foro

| Metodo | Ruta | Descripcion |
| --- | --- | --- |
| GET | `/tema_foro` | Lista todos los temas del foro. |
| GET | `/tema_foro/{id}` | Consulta un tema de foro por ID. |
| POST | `/tema_foro` | Crea un tema de foro. |
| PUT | `/tema_foro/{id}` | Actualiza un tema de foro. |
| DELETE | `/tema_foro/{id}` | Elimina un tema de foro. |

Ejemplo de JSON para crear o actualizar:

```json
{
  "titulo": "Control de plagas",
  "descripcion": "Consejos para identificar plagas en cultivos.",
  "id_autor": 1,
  "id_categoria": 2,
  "estado": "abierto",
  "activo": true
}
```

Antes de guardar, la API valida que:

- `id_autor` exista en `"Usuarios"."Usuario"`.
- `id_categoria` exista en `"Categoria"`.

### Respuestas de foro

| Metodo | Ruta | Descripcion |
| --- | --- | --- |
| GET | `/respuesta_foro` | Lista todas las respuestas del foro. |
| GET | `/respuesta_foro/{id}` | Consulta una respuesta por ID. |
| POST | `/respuesta_foro` | Crea una respuesta de foro. |
| PUT | `/respuesta_foro/{id}` | Actualiza una respuesta de foro. |
| DELETE | `/respuesta_foro/{id}` | Elimina una respuesta de foro. |

Ejemplo de JSON para crear o actualizar:

```json
{
  "id_tema": 1,
  "id_autor": 2,
  "descripcion": "Una opcion recomendada es revisar las hojas afectadas.",
  "activo": true
}
```

Antes de guardar, la API valida que:

- `id_tema` exista en `"TemaForo"`.
- `id_autor` exista en `"Usuarios"."Usuario"`.

## Respuestas de la API

Las respuestas exitosas se devuelven en formato JSON. Para errores, la API responde con un objeto como:

```json
{
  "error": "mensaje del error"
}
```

Codigos comunes:

- `200 OK`: consulta, actualizacion o eliminacion exitosa.
- `201 Created`: registro creado correctamente.
- `400 Bad Request`: JSON invalido, ID invalido o referencia inexistente.
- `404 Not Found`: registro no encontrado.
- `409 Conflict`: registro duplicado, por ejemplo `nombre_categoria`.
- `500 Internal Server Error`: error interno o problema al consultar la base de datos.

## Notas

- Los campos `fecha_creacion` y `fecha_modificacion` son devueltos desde la base de datos.
- Las eliminaciones se realizan con `DELETE` fisico sobre la tabla correspondiente.
- Los valores permitidos para `estado` dependen de las restricciones configuradas en PostgreSQL.
