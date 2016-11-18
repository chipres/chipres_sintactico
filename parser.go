package chipres

import (
	"fmt"
	"io"
)

// SelectStatement represents a SQL SELECT statement.
type SelectStatement struct{}
type FromStatement struct{}
type Statement struct{}
type Statements struct{}

// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (p *Parser) Parse() (*Statement, error) {
	stmt := &Statement{}
	// First token should be a "SELECT" keyword.

	// First token should be a "SELECT" keyword.
	if tok, lit := p.scanIgnoreWhitespace(); tok != cicloif {
		return nil, fmt.Errorf("encontro %q, esperaba cicloif", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CORCHETIZQUI {
		return nil, fmt.Errorf("encontro %q, esperaba [", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba IDENTIFICADOR", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CORCHETDERE {
		return nil, fmt.Errorf("encontro %q, esperaba ]", lit)
	}	
	if tok, lit := p.scanIgnoreWhitespace(); tok != DPUNTOS {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PARIZQI {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PARDERE {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}	
	if tok, lit := p.scanIgnoreWhitespace(); tok != celseif {
		return nil, fmt.Errorf("encontro %q, esperaba celseif", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CORCHETIZQUI {
		return nil, fmt.Errorf("encontro %q, esperaba [", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba IDENTIFICADOR", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CORCHETDERE {
		return nil, fmt.Errorf("encontro %q, esperaba ]", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DPUNTOS {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PARIZQI {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PARDERE {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != cicloelse {
		return nil, fmt.Errorf("encontro %q, esperaba cicloelse", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PARIZQI {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PARDERE {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PUNTOYCOMA{
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	return stmt, nil

}


func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	
	p.buf.tok, p.buf.lit = tok, lit

	return
}

func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}

