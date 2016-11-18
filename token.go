package chipres

// Token represents a lexical token.
type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS

	// Literals
	IDENT // main

	// Misc characters
	ASTERISK // *
	//COMMA    // ,
	PUNTOYCOMA
	DPUNTOS
	PARDERE
	PARIZQI
	CORCHETIZQUI
	CORCHETDERE

	//Pregunta
	cicloif
	celseif
	cicloelse

	// Keywords
	//SELECT
	//FROM
)
