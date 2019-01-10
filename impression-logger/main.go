package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/bigquery"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

func main() {
	ctx := context.Background()

	projectID := "madhive-central"

	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	fmt.Println("Client created!")

	query := client.Query("SELECT madhive_id, sim.url FROM `madhive-central.jelly.logs` WHERE datehour = TIMESTAMP('2018-12-16')")

	fmt.Println("Querying for data...")

	it, err := query.Read(ctx)
	if err != nil {
		fmt.Printf("Error: %v\r\n", err)
		os.Exit(1)
	}

	counter := 1
	var leaves [][]byte

	for {
		var record []bigquery.Value

		err := it.Next(&record)
		if err == iterator.Done {
			break
		}

		if err != nil {
			fmt.Printf("Error: %v\r\n", err)
			os.Exit(1)
		}

		leaves = append(leaves, hashSHA256(fmt.Sprintf("%v%v", record[0], record[1])))

		counter = counter + 1
	}

	if len(leaves) == 0 {
		return
	}

	fmt.Printf("Number of nodes: %v\r\n", len(leaves))

	root := createRow(leaves)

	fmt.Printf("Merkle Root: %X\r\n", root)
}

func createRow(r [][]byte) []byte {
	var row [][]byte

	for i := 0; i < len(r); i = i + 2 { // Iterate as pairs
		var a []byte
		var b []byte

		a = r[i]

		if i+1 == len(r) { // If these are the same, we are at the end of our slice and dont have a pair to hash with; we must clone
			b = r[i]
		} else { // If they dont match, we have a pair we can use...
			b = r[i+1]
		}

		row = append(row, hashSHA256(fmt.Sprintf("%X%X", a, b)))
	}

	if len(row) == 1 {
		return row[0]
	}

	return createRow(row)
}

func hashSHA256(r string) []byte {
	h := sha256.New()
	h.Write([]byte(r))

	return h.Sum(nil)
}
