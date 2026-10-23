import http from 'node:http'; http.createServer((q,s)=>s.end(JSON.stringify({seq:1,tipo:'dano'}))).listen(7779,'127.0.0.1')
