// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    evento, err := UnmarshalEvento(bytes)
//    bytes, err = evento.Marshal()

package main

import "encoding/json"

func UnmarshalEvento(data []byte) (Evento, error) {
	var r Evento
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Evento) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Evento struct {
	Alvo  *string `json:"alvo,omitempty"`
	Seq   int64   `json:"seq"`
	Tipo  Tipo    `json:"tipo"`
	Valor *int64  `json:"valor,omitempty"`
	De    *int64  `json:"de,omitempty"`
	Para  *int64  `json:"para,omitempty"`
}

type Tipo string

const (
	Dano      Tipo = "dano"
	Movimento Tipo = "movimento"
)
