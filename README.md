In this project, we will implement a database from scratch in Go, step by step.
## Go version
```
go version go1.24.1 
```

## Project Structure (Final Layout)
```
/BeamDB
    /cmd
        main.go
    /internal
        /storage
            btree.go           // B+tree implementation (Ch 3-5)
            node.go            // B+tree node operations (Ch 4)
            iterator.go        // B+tree iterator for range queries (Ch 9)
            freelist.go        // Free list for space management (Ch 7)
            kv_store.go        // KV store implementation (Ch 6)
        /table
            table.go           // Table operations (Ch 8)
            index.go           // Secondary indexes (Ch 10)
            schema.go          // Schema definitions and management (Ch 8)
        /transaction
            tx.go              // Transaction management (Ch 11-12)
            concurrency.go     // Concurrency control (Ch 12)
        /sql
            parser.go          // SQL parser (Ch 13)
            executor.go        // Query execution (Ch 14)
            ast.go             // Abstract syntax tree definitions (Ch 13)
        /util
            encoding.go        // Order-preserving encoding (Ch 9)
    /pkg
        /db
            db.go              // Public database interface
            tx.go              // Public transaction interface
    go.mod
    README.md


