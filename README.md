# O que é isso?

Um testador de estresse escrito em Golang.

É também um dos desafios finais do pos-goexpert da Full Cycle.

# Como usá-lo?

Crie a imagem do Docker:

```bash
docker build -t stress_go .
```

Execute com docker:

```bash
docker run --rm stress_go --url=https://www.google.com/ --requests=20 --concurrency=5
```

## Argumentos

|Valor|Abreviação|Descrição|
|---|---|---|
|url|u|URL que será testado contra estresse.|
|requests|r|Número total de solicitações que serão feitas.|
|concurrency|c|Número de threads que farão solicitações simultâneas.|


