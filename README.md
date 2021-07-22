# Ambiente Go

Um ambiente de _"live reload"_, onde as alterações no código são observadas e
re-executadas automaticamente, com Docker e Docker Compose.

O recurso de _"live reload"_ é possível através da lib
`github.com/cespare/reflex`, que observa as alterações de arquivos expecíficos e
permite a re-execução segura do projeto. Mais informações podem ser lidas pela
[documentação do projeto](https://github.com/cespare/reflex/blob/master/README.md).

## Requisitos

- Docker
- Docker Compose

## Execução

1. Faça o clone deste repositório.
2. Vá até a pasta raiz e execute o seguinte comando:

```shell
docker-compose up
```

3. Crie e/ou atualize quaisquer arquivos Go (`.go`, `go.mod`) e observe a
execução.
4. Para interromper, execute o seguinte comando:

```shell
docker-compose down
```

