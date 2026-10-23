package main
import ("encoding/json";"net/http";"fmt")
type Ev struct{ Seq int `json:"seq"`; Tipo string `json:"tipo"`}
func main(){ http.HandleFunc("/", func(w http.ResponseWriter,r *http.Request){ json.NewEncoder(w).Encode(Ev{1,"dano"})}); fmt.Println("ok"); http.ListenAndServe("127.0.0.1:7777",nil)}
