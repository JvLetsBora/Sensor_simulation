import requests

import paho.mqtt.client as mqtt


url = "https://httpbin.org/post"
# Callback quando uma mensagem é recebida do servidor.
def on_message(client, userdata, message):
    response = requests.post(url, json=message.payload.decode())
    print(f"Recebido: {response.json()} no tópico {message.topic}")

# Callback para quando o cliente recebe uma resposta CONNACK do servidor.
def on_connect(client, userdata, flags, rc):
    print("Conectado com código de resultado "+str(rc))
    # Inscreva no tópico aqui, ou se perder a conexão e se reconectar, então as
    # subscrições serão renovadas.
    client.subscribe("test/topic")

# Configuração do cliente
client = mqtt.Client("python_subscriber")
client.on_connect = on_connect
client.on_message = on_message

# Conecte ao broker
client.connect("broker.hivemq.com", 1883, 60)

# Loop para manter o cliente executando e escutando por mensagens
client.loop_forever()