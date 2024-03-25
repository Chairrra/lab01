/*Problema: Scrivere un programma Go **angolo_triangolo.go** che, date in input le ampiezze in
gradi (float64) di due angoli di un triangolo, determini l’ampiezza del terzo angolo.

Nota. La somma degli angoli di un triangolo è sempre 180 gradi, è quindi il caso di utilizzare una costante per questo valore.
*/
package main
import (
"fmt"
)

func main (){
const TOT = 180
var a1, a2 float64
fmt.Scan(&a1,&a2)
a3:= 180 -(a1+a2)
fmt.Println(a3)
}
