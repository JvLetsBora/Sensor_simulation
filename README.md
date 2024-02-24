# Sensor_simulation
Entregas referentes ao 9 trimestre da minha formação em Engenharia da computação. 


## Setup do ambiente 

Este tutorial pressupõe que você já tenha instalado as seguintes tecnologias: 
- [Python](https://www.python.org)
- [Mosquitto](https://mosquitto.org)

Para o uso dessa funcionalidade siga as etapas abaixo:

Entre no diretório 'Sensor_simulation\mqtt-paho':
```
cd mqtt-paho
```

É recomendavel o uso de virtual environment. Para isso, rode:
```
python -m venv venv
```

Para ativar o venv em MAC e UBUNTU, use:
```
source venv/bin/activate
```

Para ativar o venv no WINDOWS, use:
```
venv\Scripts\activate
```

Agora, instale as dependências do projeto:
```
pip install -r requirements.txt 
```

Com as dependências instaladas basta rodar o projeto com o comando abaixo:
```
python MQTT_Publisher.py
```

 ## Funcionamento do módulo 
Caso todas as etapas acima estejam satisfeitas, o resultado esperado encontra-se neste vídeo de demonstração: [Click Para Ver](https://drive.google.com/drive/folders/1mXoaW-fK-zhebGrNRBtqksd4BWKMuF36?usp=sharing)