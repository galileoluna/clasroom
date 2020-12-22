# Clasroom
Clasroom es una aplicación CLI, construida con el framework  [Cobra](https://github.com/spf13/cobra),  que nos permite generar una aplicación de linea de comandos y agregar los comandos que yo necesite. Luego para persistir todos los datos, utilizare una base de datos Postgres.

![main](https://raw.githubusercontent.com/galileoluna/clasroom/main/images/main.png)
## Creación del Proyecto
Para crear el proyecto , utilizaremos el go modules para gestionar las  dependencias, **cobra** y **postgres**.

Vamos a  contar con tres paquetes, **db**, el que posee toda la lógica de negocios, que estará conformado por tres clases:

 - alumnos.go
 - clase.go
 - asistencia.go

Cada una de estas  tendrá sus respectivos objetos, como también las operaciones de escritura y lectura de la base de datos,  como también tendremos el método **NewNombreClase**, para facilitar la creación de objetos.

Por otra parte tenemos la clase **conexionPG.go** que es la encargada de instanciar la base de datos, que luego utilizaremos en las otras clases, para ser utilizado en otra computadora hay que cambiar el password correspondiente del cliente postgres que tenga.

Luego tenemos la el paquete cmd que es el que genera automáticamente, donde nos proporciona una clase root.go que es el main de cobra, y luego mediante los comandos cobra add nombreComando, nos genera las clases donde invocamos las funciones de db.

Y por ultimo tenemos el paquete test que probamos el paquete db, donde realizo unos test simples para validar el funcionamiento correcto.  

## Modelo de Datos

En el script que generamos para crear la base de datos tenemos unos datos precargados de alumnos y clases.

![Modelo de datos](https://raw.githubusercontent.com/galileoluna/clasroom/main/images/diagrama.png)

## Ejecutandolo 

Para ejecutarlo tenemos que ir a la consola y luego llamarla con la palabra clasroom seguido del comando que deseemos utilizar,

Con el comando getAlumnos chequeo los alumnos existentes

![getAlumnos](https://raw.githubusercontent.com/galileoluna/clasroom/main/images/getAlumnos.png)
Con el comando getClases chequeo las clases existentes
![getClases](https://raw.githubusercontent.com/galileoluna/clasroom/main/images/getClases.png)

Para ver los participantes de una clase debemos ejecutar el comando clasroom getAsistencias
seguidos del nombre de la  clase que deseamos conocer.

![getAsistencias](https://raw.githubusercontent.com/galileoluna/clasroom/main/images/getAsistencia.png)

Para agregar un asistente a una clase tenemos que ingresar el comando clasroom addAsistencia
seguido del nombre del alumno y la clase que va a participar.
![addAsistencia](https://raw.githubusercontent.com/galileoluna/clasroom/main/images/addAsistencia.png)

Por ultimo chequeamos que el alumno haya sido 
![chequeo asistencias](https://raw.githubusercontent.com/galileoluna/clasroom/main/images/getAsistenciaPost.png)
## Problemas 

Uno de los principales problemas que tuve fue que en un principio lo desarrolle, con sqllite pero al tener Windows, por falta de GCC, no se podía ejecutar. Y el mayor problema que tuve fue con cuestiones de la modularidad del proyecto, ya que por cuestiones de go, hay practicas que tengo acostumbre, no me lo permite, otro ejemplo es los imports, como a la hora de crear test, que si o si tienen que ser **nombreClase_test.go** de otra forma nos sugiere que no existen tests. 

## Para mejorar

Se podría agregar validaciones extras para no causar problemas, a la hora de escribir en la base de datos, o controlar lo que el usuario ingresa atreves de la línea de comandos.