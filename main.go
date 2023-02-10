package main

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func main() {
	// Connect to Neo4j
	driver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "password", ""))
	if err != nil {
		fmt.Println("Error connecting to Neo4j:", err)
		return
	}
	defer driver.Close()

	// Create a session
	session, err := driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}
	defer session.Close()

	// Create a node
	createResult, err := session.Run("CREATE (p:Person {name: $name})", map[string]interface{}{"name": "John Doe"})
	if err != nil {
		fmt.Println("Error creating node:", err)
		return
	}
	if createResult.Err() != nil {
		fmt.Println("Error in create result:", createResult.Err())
		return
	}

	// Read a node
	readResult, err := session.Run("MATCH (p:Person {name: $name}) RETURN p", map[string]interface{}{"name": "John Doe"})
	if err != nil {
		fmt.Println("Error reading node:", err)
		return
	}
	if readResult.Err() != nil {
		fmt.Println("Error in read result:", readResult.Err())
		return
	}
	if readResult.Next() {
		fmt.Println("Read node:", readResult.Record().GetByIndex(0).(neo4j.Node).Props())
	} else {
		fmt.Println("Node not found")
	}

	// Update a node
	updateResult, err := session.Run("MATCH (p:Person {name: $name}) SET p.age = $age", map[string]interface{}{"name": "John Doe", "age": 30})
	if err != nil {
		fmt.Println("Error updating node:", err)
		return
	}
	if updateResult.Err() != nil {
		fmt.Println("Error in update result:", updateResult.Err())
		return
	}

	// Delete a node
	deleteResult, err := session.Run("MATCH (p:Person {name: $name}) DELETE p", map[string]interface{}{"name": "John Doe"})
	if err != nil {
		fmt.Println("Error deleting node:", err)
		return
	}
	if deleteResult.Err() != nil {
		fmt.Println("Error in update result:", deleteResult.Err())
		return
	}
}
