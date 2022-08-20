# go-rabbitmq
Golang implementation for messaging service using Rabbit MQ where Publisher and Subscriber component is added.

steps involved:
--------------
- try to connect to rabbitmq
- start listening for messages
- create consumere
- watch the queue and consume events from the topic


NOTES:
-----
- write some funcs that allow us to interact with the queue
- put rmq related func & code in their own package
- broker is going to have to push something onto the queue
- now this listener service right now at least is not going to put anything onto the queue
- instead, the queue will push to this service which listens for things
