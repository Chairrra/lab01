/*crivere un programma Go **quoziente_resto.go** che legge da standard input un
dividendo e un divisore (interi), calcola il quoziente e il resto e li stampa.*/

package main
import(
  "fmt"
)
func main (){
  var x1, x2 int
  fmt.Println("due numeri interi:")
  fmt.Scanf("%d %d", &x1, &x2)
  quoziente:= x1/x2
  resto:= x1%x2
  fmt.Printf("quoziente: %d, resto: %d", quoziente, resto)
}
