const mqtt = require('mqtt');

const client = mqtt.connect('mqtt://localhost:1883'); // Replace with your broker URL

client.on('connect', () => {
    console.log('Connected to MQTT broker');
    
    client.subscribe('test/topic', (err) => {
        if (err) {
            console.error('Failed to subscribe:', err);
        } else {
            console.log('Subscribed to topic: test/topic');
        }
    });
});

client.on('message', (topic, message) => {
    console.log(`Received message: ${message.toString()} on topic: ${topic}`);
});

client.on('error', (err) => {
    console.error('MQTT Client Error:', err);
});
