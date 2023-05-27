package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// >>> REPOSITORIOS É A CAMADA QUE VAI INTERAGIR COM O BANCO DE DADOS <<< //

// Usuarios representa um repositorio de usuarios
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um novo repositorio de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuario no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statment, erro := repositorio.db.Prepare(
		"INSERT INTO usuarios (nome, nick, email, senha) VALUES(?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statment.Close()

	resultado, erro := statment.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

/*
	A função NovoRepositorioDeUsuarios retorna um ponteiro para uma estrutura usuarios
	que contém um objeto *sql.DB como campo. Aqui está uma explicação passo a passo do código:

    func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios: Esta linha define a função NovoRepositorioDeUsuarios
	que recebe um ponteiro para um objeto *sql.DB como argumento e retorna um ponteiro para uma estrutura usuarios.

    return &usuarios{db}: Nesta linha, a função retorna um ponteiro para a estrutura usuarios.
	A estrutura usuarios é criada usando um literal de estrutura (usuarios{}), e o campo db é inicializado
	com o valor do objeto db passado como argumento. O & antes de usuarios{db} retorna o endereço de memória
	da estrutura usuarios, criando um ponteiro para essa estrutura. Isso permite que o chamador da função receba
	um ponteiro para a estrutura usuarios que contém o objeto *sql.DB passado como argumento.
*/
