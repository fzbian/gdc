package models

type SaldosPlataforma struct {
	PlataformaID  int    `db:"plataforma_id" json:"plataforma_id"`
	Saldo         int    `db:"saldo" json:"saldo"`
	FechaCreacion string `db:"fecha_creacion" json:"fecha_creacion"`
}
