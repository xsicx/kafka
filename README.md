# KAFKA

### Install 
`make install` 

### Produce
`docker exec -ti cli /go/bin/cli produce --topic=topic1 --server=kafka:9092`

### Consume
`docker exec -ti cli /go/bin/cli consume --topic=topic1 --server=kafka:9092`

### Help Services
- kafdrop http://localhost:9000/
