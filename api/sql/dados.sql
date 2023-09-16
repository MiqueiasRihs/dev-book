INSERT INTO usuarios (nome, nick, email, senha) VALUES ('Usuário 1', 'user1', 'user1@example.com', '$2a$10$T3JlOSM3HIc.GjNpLOPOeu9yghftvkpd/65csT8ri7O8bU0oeOgLO');
INSERT INTO usuarios (nome, nick, email, senha) VALUES ('Usuário 2', 'user2', 'user2@example.com', '$2a$10$T3JlOSM3HIc.GjNpLOPOeu9yghftvkpd/65csT8ri7O8bU0oeOgLO');
INSERT INTO usuarios (nome, nick, email, senha) VALUES ('Usuário 3', 'user3', 'user3@example.com', '$2a$10$T3JlOSM3HIc.GjNpLOPOeu9yghftvkpd/65csT8ri7O8bU0oeOgLO');

INSERT INTO seguidores (usuario_id, seguidor_id) VALUES (1, 2), (3, 1), (1, 3);


INSERT INTO publicacoes(titulo, conteudo, autor_id)
VALUES
    ("Publicação do usuário 1", "Essa é a publicação do usuário 1! Oba!", 1),
    ("Publicação do usuário 2", "Essa é a publicação do usuário 2! Oba!", 2),
    ("Publicação do usuário 3", "Essa é a publicação do usuário 3! Oba!", 3);