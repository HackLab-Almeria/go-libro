package main

// https://yourbasic.org/golang/format-parse-string-time-date-example/
import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func restarMeses(num_mes_a_restar int, fecha_introducida string) string {

	fecha_parseada, _ := time.Parse("2006-01-02", fecha_introducida)

	fecha_final := fecha_parseada.Add(time.Duration(num_mes_a_restar) * -30 * 24 * time.Hour).Format("2006-01-02")

	return fecha_final
}

func compararAnos(fecha_introducida string) float64 {
	fecha_parseada, _ := time.Parse("2006-01-02", fecha_introducida)
	nacimiento_autor, _ := time.Parse("2006-01-02", "1994-11-08")
	diferencia := fecha_parseada.Sub(nacimiento_autor)
	anos_diferencia := diferencia.Hours() / 8760
	return anos_diferencia

}

func main() {
	// Fecha inmediata
	fecha := time.Now()
	fmt.Println("La fecha es:", fecha)

	// Obtención de año, mes y día de la semana
	fmt.Println("El año es ", fecha.Year())
	fmt.Println("El mes es ", fecha.Month())
	fmt.Println("El día es ", fecha.Day())

	// Cambios de formato

	// RFC1123
	fmt.Println("En formato RFC1123 ", fecha.Format(time.RFC1123))

	// Kitchen
	fmt.Println("La hora es: ", fecha.Format(time.Kitchen))

	// Stamp con milisegundos
	fmt.Println("Fecha de ejecución ", fecha.Format(time.StampMilli))

	// Cálculos con fechas https://golangbyexample.com/understanding-duration-go/

	// Restarle  X número de meses a una fecha introducida por parámetro
	cantidad_a_restar := os.Args[1]
	cantidad_convertida, _ := strconv.Atoi(cantidad_a_restar)

	fecha_primera := os.Args[2]
	fmt.Println("La fecha pasada es: ", restarMeses(cantidad_convertida, fecha_primera))

	fmt.Println("Han pasado estos años desde el nacimiento del autor ", compararAnos(fecha_primera))

}
