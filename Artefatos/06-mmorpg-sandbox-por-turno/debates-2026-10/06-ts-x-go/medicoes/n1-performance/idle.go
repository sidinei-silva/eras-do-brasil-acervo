package main
import ("net/http";"time";"os")
func main(){ go http.ListenAndServe("127.0.0.1:0",nil); time.Sleep(2000*time.Millisecond); os.Exit(0) }
