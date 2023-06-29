package GenerarArchivos

import (
	"io/ioutil"
	"os"
	"os/exec"
)

func CrearArchivo(nombre_archivo string) {
	var _, err = os.Stat(nombre_archivo)

	if os.IsNotExist(err) {
		var file, err = os.Create(nombre_archivo)
		if err != nil {
			return
		}
		defer file.Close()
	}
	//fmt.Println("Archivo generado exitosamente")
}

func EscribirArchivo(contenido string, nombre_archivo string) {
	var file, err = os.OpenFile(nombre_archivo, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	_, err = file.WriteString(contenido)
	if err != nil {
		return
	}
	err = file.Sync()
	if err != nil {
		return
	}
	//fmt.Println("Archivo guardado correctamente")
}

func Ejecutar(nombre_imagen string, archivo string) {
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tjpg", archivo).Output()
	mode := 0777
	_ = ioutil.WriteFile(nombre_imagen, cmd, os.FileMode(mode))
}
