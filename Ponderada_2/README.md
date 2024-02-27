# Simulador de ataques MQTT

| Aluno                | Curso                  | Módulo | Grupo |
|----------------------|------------------------|--------|-------|
| João Vitor Oliveira  | Engenharia da Computação | 9      | 5     |

Atividade feita em conjunto com os membros do grupo: Caio e Kil

## Descrição
Avaliar a tríade CIA em uma conexão com um broker MQTT.

## Setup do ambiente 

Siga as instruções abaixo para realizar a simulação dos testes:

1. Certifique-se de ter instalado as seguintes tecnologias: 
   - [Python](https://www.python.org)
   - [Mosquitto](https://mosquitto.org)
   - [Docker](https://www.docker.com/get-started/)

2. Execute os passos do seguinte diretório: [Link para o diretório](https://github.com/rafaelmatsuyama/Inteli-T2-EC-M09)

## Perguntas - Roteiro
1. O que acontece se você utilizar o mesmo ClientID em outra máquina ou sessão do navegador? Algum pilar do CIA Triad é violado com isso?
    - R: Ele é desconectado da sessão anterior e conectado na nova sessão. O pilar de Disponibilidade é violado, pois o cliente é desconectado sem aviso prévio.

2. Com os parâmetros de recursos, algum pilar do CIA Triad pode ser facilmente violado?
    - R: Sim, como nesse exemplo não há a criação de nível de acesso aos tópicos, é possível ferir a Confiabilidade do sistema, já que todo usuário tem acesso a qualquer tópico do sistema. Outra forma de ferir a CIA Triad é pela Disponibilidade, pois o contêiner pode ser facilmente derrubado por falta de recursos em uma eventual sobrecarga de mensagens.

3. Já tentou fazer o Subscribe no tópico #? (sim, apenas a hashtag). O que acontece?
    - R: O '#' é um wildcard multi-level, um método coringa que permite inscrever-se em todos os tópicos disponíveis. Isso pode ser perigoso, pois pode violar a confidencialidade das mensagens.

4. Sem autenticação (repare que a variável allow_anonymous está como true), como a parte de confidencialidade pode ser violada?
    - R: Como qualquer pessoa pode se conectar ao broker e ler mensagens de qualquer tópico, acaba por violar a Confidencialidade destas.

## Perguntas - Desenvolvimento

### 1. Como você faria para violar a confidencialidade?
    R: Para violar a confidencialidade, eu poderia me inscrever em tópicos que transitam informações sensíveis, uma vez que o broker não possui uma lista de controle de acesso (ACL) para limitar o acesso a tópicos específicos. Sendo assim tendo quaisquer credenciais de acesso, poderia publicar e ler mensagens em tópicos que não deveria ter acesso.
Código para simular ataque:
```
python ./ataques/confiabilidade.py
```

Sensor_simulation\Ponderada_3\ataques\confiabilidade.py


![img alt](static/confiabilidade.png)

### 2. Como você faria para garantir a integridade do broker MQTT?
    R: Para garantir a integridade dos dados, além de credenciais de acesso, criaria uma lista de controle de acesso (ACL) para tópicos específicos, de forma a garantir que apenas usuários autorizados possam publicar mensagens em tópicos específicos, evitando o comprometimento da integridade dos dados. Além disso, criaria uma lógica para manter de forma criptografada um histórico de mensagens, juntamente com um protocolo de backups mensais das mesmas. 


### 3. Tente simular uma violação do pilar de Disponibilidade.

    R: Tendo acesso ao host, porta e credenciais de protocolo mqtt, é possível fazer um script que faça uma cadência de requisições superior a memória do servidor, assim fazendo com que novas mensagens não sejam registradas e posteriormente o server venha a cair.
Código para simular ataque:
```
python ./ataques/disponibilidade.py
```

Sensor_simulation\Ponderada_3\ataques\disponibilidade.py

![img alt](static/disponibilidade.png)

Me basiei nessa referência [( CIA Triad Introduction )](https://informationsecurity.wustl.edu/items/confidentiality-integrity-and-availability-the-cia-triad/Introdução) para responder a esse autoestudo.


