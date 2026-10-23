export type Evento =
  | { tipo: "dano"; seq: number; alvo: string; valor: number }
  | { tipo: "movimento"; seq: number; de: number; para: number };
