function rnd(s){ s ^= s<<13; s>>>=0; s ^= s>>>17; s ^= s<<5; return s>>>0; }
function aplicar(st, acao){
  let r = rnd(st.rng); const d20 = r%20+1; const ev=[];
  const alvo = st.mobs.findIndex(m=>m.hp>0);
  const mobs = st.mobs.map(m=>({...m}));
  ev.push({tipo:'ataque',rolagem:d20});
  if (d20+5 >= mobs[alvo].ac){ r=rnd(r); const dano=r%8+3; mobs[alvo].hp-=dano; ev.push({tipo:'dano',alvo,dano}); }
  let heroi={...st.heroi};
  for(const m of mobs){ if(m.hp<=0) continue; r=rnd(r); if(r%20+1+3>=heroi.ac){ r=rnd(r); const d=r%6+1; heroi.hp-=d; ev.push({tipo:'dano',alvo:-1,dano:d}); } }
  return {estado:{rng:r,heroi,mobs,turno:st.turno+1},eventos:ev};
}
function luta(seed){
  let st={rng:seed||1,turno:0,heroi:{hp:60,ac:16},mobs:[{hp:20,ac:13},{hp:20,ac:13}]}; let n=0;
  while(st.heroi.hp>0 && st.mobs.some(m=>m.hp>0)){ const o=aplicar(st,'atacar'); st=o.estado; n+=o.eventos.length; }
  return [st.turno,n];
}
const N=+process.argv[2]||1000; const t=process.hrtime.bigint(); let tt=0,ee=0;
for(let i=1;i<=N;i++){const [a,b]=luta(i*2654435761>>>0);tt+=a;ee+=b;}
const ms=Number(process.hrtime.bigint()-t)/1e6;
console.log(`TS/Node N=${N} ${ms.toFixed(1)}ms turnos=${tt} eventos=${ee} rss=${(process.memoryUsage().rss/1e6).toFixed(0)}MB`);
