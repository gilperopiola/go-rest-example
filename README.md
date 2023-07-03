# **go-rest-example (Go Rest Example)**

##### This project was created in 2018-2019, now we're on 2023 and i want to rip my eyes off reading the code 
##### UPDATE: Finally refactored it into something decent. Still needs sum luv <3

## **Architecture**

### **cmd/main.go** contains the main file, which creates all the dependencies and starts the application
### **pkg/codec** holds the transformations between the entities and the models
### **pkg/config** holds the configuration of the app
### **pkg/entities** holds the data structures used in the transport/services layers
### **pkg/models** holds the data structures used for the database
### **pkg/repository** holds all the database code, including setup and queries
### **pkg/service** holds all the business logic code, it accepts and returns entities
### **pkg/transport** holds the different endpoints and handles requests/responses
### **pkg/utils** holds the utils