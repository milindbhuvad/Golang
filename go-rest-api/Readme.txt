Get all people:
    http://localhost:8080/people

Get one person:
    http://localhost:8080/people/1

Create a new person:
    http://localhost:8080/people

    curl -X POST -H "Content-Type: application/json" \
    -d '{"id":"3","firstName":"Alice","lastName":"Smith","age":22}' \

Update a person:
    http://localhost:8080/people/1
    
    curl -X PUT -H "Content-Type: application/json" \
    -d '{"id":"1","firstName":"John","lastName":"Doe","age":35}' \

Delete a person:

    curl -X DELETE http://localhost:8080/people/2