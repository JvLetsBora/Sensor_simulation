# Simulação de dispositivos MQTT (Prova)




## Setup do ambiente 

Siga as instruções abaixo para realizar a simulação dos testes:

1. Certifique-se de ter instalado as seguintes tecnologias: 
   - [go](https://rmnicola.github.io/m9-ec-encontros/go)
   - [Mosquitto](https://mosquitto.org)
   - [Docker](https://www.docker.com/get-started/)



## Testes 

Siga a sequência de comandos abaixo para realizar os testes:

Observação: É necessário um arquivo '.env' nos diretórios './publisher' e './subscribe' com as seguintes variáveis de ambiente:
- BROKER_ADDR=<ENDEREÇO DO SEU BROKER>
- HIVE_USER=<USUÁRIO DO SEU BROKER>
- HIVE_PSWD=<SENHA DO SEU BROKER>

**Testando o publisher**

1. Entre no diretório usando o comando:
```
   cmd cd \publisher
```

2. Rode o comando de teste:
```
cmd go test
```

Resulatdo esperado:



**Testando o subscriber**
Para realizar este test, é necessário que haja um publicador enviando uma mensagem de teste para o tópico no qual a inscrição está registrada.

1. Entre no diretório usando o comando:
```
cmd cd \subscribe
```

2. Rode o comando de teste:
```
cmd go test
```
Video de funcionamento:
<br>

[link](https://drive.google.com/drive/folders/1SexpJoBAbcSdW7PNzwS3-qXH-RLEknLD?usp=sharing)
