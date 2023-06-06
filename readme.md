# Bike Soft Maintenance

Este repositorio contiene una aplicación de Flutter y una API REST escrita en Go para la gestión de mantenimiento de bicicletas.

## Configuración del proyecto

Sigue los siguientes pasos para configurar el proyecto en tu máquina local.

### 1. Clonación del repositorio

Primero, clona el repositorio en tu máquina local utilizando Git. Puedes hacer esto con el siguiente comando en tu terminal:

```shell
git clone https://github.com/ucaldas-tech-titans/bike-soft-maintenance.git
```

### 2. Configuración de la API REST de Go

La API REST está escrita en Go y utiliza un servidor Gin y una base de datos MySQL. Para configurar la API, necesitas tener Go instalado en tu máquina.

Instala las dependencias del proyecto utilizando el comando go get.

Configura la cadena de conexión a la base de datos en un archivo .env en la raíz del proyecto. La cadena de conexión se carga en el código a través de la función VariablesEnv(key string) que busca variables en el archivo .env.

El archivo main.go inicializa el servidor y las rutas y se conecta a la base de datos. El esquema de la base de datos se migra automáticamente durante la inicialización.

### 3. Configuración de la aplicación Flutter

La aplicación de Flutter depende de varias bibliotecas, como http y flutter_local_notifications. Para configurar la aplicación, necesitas tener Flutter y Dart instalados en tu máquina.

Instala las dependencias del proyecto utilizando el comando flutter pub get.

Modifica las URLs de la aplicación en flutter para que coincidan con las del backend local

Si estás corriendo en el navegador debes usar una extensión de habilitación de CORS.
<https://chrome.google.com/webstore/detail/allow-cors-access-control/lhobafahddgcelffkeicbaginigeejlf>

Puedes iniciar la aplicación utilizando el comando flutter run.

## Ejecución de las aplicaciones

Sigue los siguientes pasos para ejecutar las aplicaciones en tu máquina local.

### Ejecución de la API REST de Go

Para ejecutar la API REST de Go, navega al directorio de la API y ejecuta el siguiente comando:

```shell
go run main.go
```

Este comando iniciará el servidor en el puerto 8080. Puedes interactuar con la API a través de esta dirección: <http://localhost:8080>.

Si te da error porque el puerto está ocupado modifícalo en el main.go. Recuerda también modificar las URLs de flutter.

### Ejecución de la aplicación Flutter

Para ejecutar la aplicación Flutter, navega al directorio de la aplicación y ejecuta el siguiente comando:

```shell
flutter run
```

Este comando iniciará la aplicación Flutter en un emulador o dispositivo físico si está conectado. Asegúrate de tener un emulador en ejecución o un dispositivo físico conectado antes de ejecutar este comando.
