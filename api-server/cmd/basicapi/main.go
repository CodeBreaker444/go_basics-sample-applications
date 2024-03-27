package main
import(
	"fmt"
	"net/http"
	"apiserver/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)
func main(){
	log.SetReportCaller(true)
	fmt.Printf("API Server - v1")
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler_call(r)
	err:=http.ListenAndServe("localhost:8000",r)
	if err!=nil{
		log.Error(err)

	}
}