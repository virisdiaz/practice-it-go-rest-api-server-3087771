package practice

import (
	"fmt"
	"net/http"
)

func Test(){
	fmt.Print("sucessful called the practice module")
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Word form the Practice module")

}