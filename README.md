# ¿Qué es esto?
Aquí trataré de hacer la intro de fiber según la documentación y algunos ejemplos de los del [github de fiber](https://github.com/gofiber/recipes)

# Recursos:
* [fiber](https://docs.gofiber.io) framework web que tiene como mayor ventaja la velocidad de su enrutador
* [ejemplos de fiber](https://github.com/gofiber/recipes) sobre implementación de varias cosas comunes en desarrollo web
* [ayudó a entender el proceso con los jwt](https://dev.to/koddr/build-a-restful-api-on-go-fiber-postgresql-jwt-and-swagger-docs-in-isolated-docker-containers-475j#a-middleware-functions)
* [conector para Postgres](https://github.com/jackc/pgx) y el [tutorial](https://github.com/jackc/pgx/wiki/Getting-started-with-pgx), además se usó [el subpaquete pgxpool](https://pkg.go.dev/github.com/jackc/pgx/v4@v4.11.0/pgxpool)
* Se usó el sitio web [mockaroo.com](https://mockaroo.com/) para generar dummy info 
* [Paquete JWT para Fiber](https://pkg.go.dev/github.com/gofiber/jwt/v2@v2.2.7) la documentación deja muy buenas ejemplos de cómo funciona y con diferentes codificadores
* [¿Por qué GORM es una mala idea?](https://www.reddit.com/r/golang/comments/umkgk3/gorm_is_a_bad_idea/) y aquí un [artículo de complemento](https://alanilling.com/exiting-the-vietnam-of-programming-our-journey-in-dropping-the-orm-in-golang-3ce7dff24a0f)
* [Squirrel](https://github.com/Masterminds/squirrel) para generación de código SQL dinámicamente
* [Structurable (Struct-Table Mapping for Go)](https://github.com/Masterminds/structable) no me acaba de convencer pero le echaré un ojo
* [Swagger para Fiber](https://github.com/gofiber/swagger): Solo hay que tener en cuenta que si no tienes el path al directorio en que se genera los ejecutables de Go no puedes ejecutar el ````$ swag snit````. Por debajo de este paquete, se usa swaggo y aquí está una (referencia a los marcadores de swaggo)[https://pkg.go.dev/github.com/swaggo/swag]

# TODO:
- [X] Configurar servidor de Fiber
- [X] Configurar router
- [X] Probar que el pool de pgx funciona bien mientras esté abierto el server (No comprueba que la conexión es correcta hasta que pides que abra una conexión y es así por diseño)
- [X] Hacer un CRUD con Persons (Necesito cargar el schema)
- [X] Probar Squirrel
- [X] Implementar Middleware JWT Auth
- [X] Implementar Swagger
- [X] Implementar Middleware CORS
- [X] Implementar Middleware Recover
- [ ] Implementar carga de archivos
- [ ] Implementar endpoints para websockets
