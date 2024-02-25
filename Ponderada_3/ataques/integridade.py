import time
import paho.mqtt.client as mqtt



pub_broker = {
    "link":"172.27.112.1",
    "port":1883
}

def on_publish(client, userdata, mid, reason_code, properties):
    # reason_code and properties will only be present in MQTTv5. It's always unset in MQTTv3
    try:
        userdata.remove(mid)
    except KeyError:
        print("on_publish() is called with a mid not present in unacked_publish")
        print("This is due to an unavoidable race-condition:")
        print("* publish() return the mid of the message sent.")
        print("* mid from publish() is added to unacked_publish by the main thread")
        print("* on_publish() is called by the loop_start thread")
        print("While unlikely (because on_publish() will be called after a network round-trip),")
        print(" this is a race-condition that COULD happen")
        print("")
        print("The best solution to avoid race-condition is using the msg_info from publish()")
        print("We could also try using a list of acknowledged mid rather than removing from pending list,")
        print("but remember that mid could be re-used !")

unacked_publish = set()

mqttc = mqtt.Client(mqtt.CallbackAPIVersion.VERSION2)
mqttc.on_publish = on_publish
mqttc.username_pw_set(username="jv",password="jv")
mqttc.user_data_set(unacked_publish)
mqttc.connect(host=pub_broker["link"],port=pub_broker["port"])
mqttc.loop_start()

msg_info2 = mqttc.publish("hello/topic", "crash start", qos=1)
unacked_publish.add(msg_info2.mid)

# Wait for all message to be published
while True:
    # Our application produce some messages
    msg_info = mqttc.publish("hello/topic", "crash", qos=1)
    unacked_publish.add(msg_info.mid)
    

    # Due to race-condition described above, the following way to wait for all publish is safer
    msg_info.wait_for_publish()


mqttc.disconnect()
mqttc.loop_stop()