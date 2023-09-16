package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Publicacoes representa um repositorio de publicações
type Publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes cria um repositorio de publicações
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare("INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES (?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query(`
		SELECT p.*, u.nick FROM
		publicacoes p INNER JOIN usuarios u
		on u.id = p.autor_id WHERE p.id = ?`,
		publicacaoID)

	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacao modelos.Publicacao

	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

// func (repositorio Publicacoes) BuscarPuplicacoes() (modelos.Publicacao, error) {
// 	linha, erro := repositorio.db.Query(`
// 		SELECT p.*, u.nick FROM
// 		publicacoes p INNER JOIN usuarios u
// 		on u.id = p.autor_id`)

// 	if erro != nil {
// 		return modelos.Publicacao{}, erro
// 	}
// 	defer linha.Close()

// 	var publicacao []modelos.Publicacao
// 	if linha.Next() {
// 		if erro = linha.Scan(
// 			&publicacao.ID,
// 			&publicacao.Titulo,
// 			&publicacao.Conteudo,
// 			&publicacao.AutorID,
// 			&publicacao.Curtidas,
// 			&publicacao.CriadaEm,
// 			&publicacao.AutorNick,
// 		); erro != nil {
// 			return modelos.Publicacao{}, erro
// 		}
// 	}

// 	return publicacao, nil

// }
