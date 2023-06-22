# Microservice_In_Golang

In this project, two microservices has been implemented.
First Microservice is JWT_CREATOR.
Second Microservice is API.

In the first microservice, user will be able to generate a JWT token and in the second microservice, user will be able to parse that token. If the token is parsed correctly, passing all the claims then the success message will be displayed.

Both the microservices are independent of each other, first one is running on port :8080, second one on :9001.
