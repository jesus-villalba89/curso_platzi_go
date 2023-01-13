package mypackage

/*
	Los nombres de clase, atributos y metodos cuando empiece en Mayus son publicos

	go mod init => seguido del nombre que tomará como la ruta base para nuestro paquete, creará un archivo de nombre go. mod en el directorio donde lo ejecutemos. Por ejemplo, si le pasamos como nombre mimodulo, todas las carpetas que estén al mismo nivel que el archivo go.
*/

//CarroPublico  carro con acceso publico
type CarroPublico struct {
	Marca  string
	Modelo int
}

func ImprimirMensaje() {

}
