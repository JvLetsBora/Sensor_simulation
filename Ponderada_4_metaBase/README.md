# Integração do Subscriber com SQLite e Metabase

## Descrição
O objetivo é integrar o subscriber a uma ferramenta de Business Intelligence (BI) usando dados persistentes.

## Setup do Ambiente 

Siga as instruções abaixo para realizar a simulação dos testes:

1. Certifique-se de ter instalado as seguintes tecnologias: 
   - [Go](https://rmnicola.github.io/m9-ec-encontros/go)
   - [Mosquitto](https://mosquitto.org)
   - [Docker](https://www.docker.com/get-started/)

## Testes 

Siga a sequência de comandos abaixo para realizar os testes:

Observação: É necessário um arquivo `.env` nos diretórios `./publisher` e `./subscribe` com as seguintes variáveis de ambiente:
- `BROKER_ADDR=<ENDEREÇO DO SEU BROKER>`
- `HIVE_USER=<USUÁRIO DO SEU BROKER>`
- `HIVE_PSWD=<SENHA DO SEU BROKER>`

**Testando o Publisher**

1. Navegue até o diretório usando o comando:
```bash
cd Sensor_simulation\Ponderada_2\publisher

```

2. Rode o comando de teste:
```
cmd go test
```

Executando os comandos acima, o resultado esperado deve ser semelhante ao que está demonstrado neste vídeo:[Link](https://drive.google.com/file/d/1QueFLvk9FpLdzFm7fInNkbERy7pIfkiU/view?usp=sharing)
