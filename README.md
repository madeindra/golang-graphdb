# Graph Database with Neo4j

Example of CRUD operation in Golang with Neo4j

## Preparation
1. Run Neo4j database

```
docker run -d -p 7474:7474 -p 127.0.0.1:7687:7687 --env NEO4J_AUTH=username/password neo4j:latest
```

2. Run application

```
go run main.go
```