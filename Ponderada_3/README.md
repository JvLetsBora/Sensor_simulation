# Simulador de ataques MQTT

| Aluno                    | Curso                  | Módulo | Grupo |
|--------------------------|------------------------|--------|-------|
| João Vitor Oliveira    | Engenharia da Computação | 9      | 5     |

## Descrição
Avaliar a tríade CIA em uma conexão com um broker MQTT.
  

![img alt](Ponderada_3/static/confiabilidade.png)
![img alt](Ponderada_3/static/disponibilidade.png)

## Setup do ambiente 

Para reproduzir os tests a seguir é necessário a instação das seguintes tecnologias: 
- [Python](https://www.python.org)
- [Mosquitto](https://mosquitto.org)
- [Docker](https://www.docker.com/get-started/)

Perguntas
1. O que acontece se você utilizar o mesmo ClientID em outra máquina ou sessão do browser? Algum pilar do CIA Triad é violado com isso?
Quando você utiliza o mesmo ClienteID, o broker vai entender que o Cliente atual trocou de conexão. Logo, ele vai deixar de mandar as mensagens que estão sendo publicadas para o primeiro e vai mandar apenas para o último que se conectou.

Impacto no CIA Triad:

Disponibilidade: Impacto potencial. As mensagens podem ser perdidas se o cliente original se desconectar antes de receber todas as mensagens.
2. Com os parâmetros de resources, algum pilar do CIA Triad pode ser facilmente violado?
Sim, os parâmetros de resources podem ser usados para violar os pilares do CIA Triad:

Confiabilidade: Um invasor pode usar recursos para negar o serviço a clientes legítimos, consumindo todos os recursos disponíveis do broker.
Integridade: Um invasor pode usar recursos para modificar mensagens em trânsito, alterando o conteúdo das mensagens.
Disponibilidade: Um invasor pode usar recursos para tornar o broker indisponível, consumindo todos os recursos do sistema.
3. Já tentou fazer o Subscribe no tópico #? (sim, apenas a hashtag). O que acontece?
Sim, é possível se inscrever no tópico "#". Isso faz com que o cliente receba todas as mensagens publicadas em qualquer tópico.

Impacto no CIA Triad:

Confiabilidade: Se existir um tópico com conteúdos sensíveis, pode ter vazamento de informações.
Disponibilidade: Impacto potencial. O cliente pode receber um grande volume de mensagens, o que pode afetar seu desempenho.
4. Sem autenticação (repare que a variável allow_anonymous está como true), como a parte de confidencialidade pode ser violada?
Sem autenticação, qualquer um pode se conectar ao broker e publicar mensagens. Isso significa que um invasor pode:

Ler mensagens confidenciais publicadas no broker.
Publicar mensagens falsas em nome de usuários legítimos.
Impacto no CIA Triad:

Confiabilidade: A existência de mensagens falsas, perdendo a confiabilidade das informações.
Integridade: Impacto alto. As mensagens podem ser modificadas ou falsificadas.
5. Simulações de Violações
Tente simular uma violação do pilar de Confidencialidade.

Conecte-se ao broker sem autenticação.
Publique uma mensagem confidencial.
Use um sniffer de rede para capturar a mensagem em trânsito.
Tente simular uma violação do pilar de Integridade.

Conecte-se ao broker sem autenticação.
Publique uma mensagem falsa em nome de um usuário legítimo.
Tente simular uma violação do pilar de Disponibilidade.

Conecte-se ao broker com vários clientes e publique um grande volume de mensagens.
Observe o desempenho do broker e dos clientes.
Conclusões
O MQTT é um protocolo de comunicação leve e eficiente, mas é importante estar ciente dos riscos de segurança associados ao seu uso. A implementação de medidas de segurança adequadas, como autenticação e autorização, é essencial para proteger a confidencialidade, integridade e disponibilidade das mensagens.

Referenicia: https://informationsecurity.wustl.edu/items/confidentiality-integrity-and-availability-the-cia-triad/Introdução