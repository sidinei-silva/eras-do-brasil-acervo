package main
import ("fmt";"os";"strconv";"time";"runtime")
type Ator struct{HP,AC int32}
type Ev struct{Tipo string;Alvo,Val int32}
type Estado struct{RNG uint32;Turno int32;Heroi Ator;Mobs []Ator}
func rnd(s uint32) uint32 { s^=s<<13; s^=s>>17; s^=s<<5; return s }
func aplicar(st Estado) (Estado, []Ev) {
 r:=rnd(st.RNG); d20:=int32(r%20+1); ev:=[]Ev{{"ataque",0,d20}}
 mobs:=append([]Ator(nil), st.Mobs...); alvo:=0; for i,m:=range mobs{ if m.HP>0 {alvo=i;break} }
 if d20+5>=mobs[alvo].AC { r=rnd(r); d:=int32(r%8+3); mobs[alvo].HP-=d; ev=append(ev,Ev{"dano",int32(alvo),d}) }
 h:=st.Heroi
 for _,m:=range mobs{ if m.HP<=0 {continue}; r=rnd(r); if int32(r%20+1)+3>=h.AC { r=rnd(r); d:=int32(r%6+1); h.HP-=d; ev=append(ev,Ev{"dano",-1,d}) } }
 return Estado{r,st.Turno+1,h,mobs}, ev
}
func vivo(ms []Ator) bool { for _,m:=range ms { if m.HP>0 {return true} }; return false }
func main(){ N:=1000; if len(os.Args)>1 { N,_=strconv.Atoi(os.Args[1]) }
 t:=time.Now(); var tt,ee int
 for i:=1;i<=N;i++{ st:=Estado{uint32(i)*2654435761,0,Ator{60,16},[]Ator{{20,13},{20,13}}}
  for st.Heroi.HP>0 && vivo(st.Mobs) { var ev []Ev; st,ev=aplicar(st); ee+=len(ev) }; tt+=int(st.Turno) }
 var m runtime.MemStats; runtime.ReadMemStats(&m)
 fmt.Printf("Go N=%d %.1fms turnos=%d eventos=%d heapSys=%dMB\n",N,float64(time.Since(t).Microseconds())/1000,tt,ee,m.Sys/1e6) }
