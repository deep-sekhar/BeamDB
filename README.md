In this project, we will implement a database from scratch in Go, step by step.
## Go version
```
go version go1.24.1 linux/amd64
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

```

For Chapter 1 (From Files To Databases):

The atomic file update concepts would be implemented in internal/storage/kv_store.go as part of the KV store implementation

The fsync for durability would also be in internal/storage/kv_store.go

For Chapter 2 (Indexing Data Structures):

The concepts about B+trees and other indexing structures inform the design of internal/storage/btree.go

The understanding of different query types (point, range) influences the overall architecture

If you wanted to explicitly represent these chapters in the file structure, you could add:

text
/internal
    /storage
        /file
            atomic_io.go       // Atomic file operations (Ch 1)
        /index
            index_types.go     // Index structure concepts (Ch 2)

## Goals
### 1: From Files to Databases
- Implement atomic file updates using the rename technique.
- Use `fsync` for durability.
- Create a basic file structure for the database.

### 2: Indexing Data Structures
- Understand different query types (scan, point, range).
- Compare indexing structures (hash tables, sorted arrays, B-trees).
- Prepare for B+tree implementation.

### 3: B-Tree & Crash Recovery
- Implement B-tree as a balanced n-ary tree.
- Design B-tree as nested arrays.
- Implement copy-on-write for crash recovery.
- Understand tree maintenance (splitting and merging).

### 4: B+Tree Node and Insertion
- Design B+tree node format.
- Implement node operations (decode, lookup).
- Implement node update operations.
- Implement node splitting.
- Implement B+tree insertion algorithm.

### 5: B+Tree Deletion and Testing
- Implement high-level B+tree interfaces.
- Implement node merging.
- Implement B+tree deletion algorithm.
- Create a testing framework for B+tree.

### 6: Append-Only KV Store
- Create a KV store with copy-on-write B+tree.
- Implement two-phase update for atomicity and durability.
- Design file layout with a meta page.
- Implement disk page management.
- Handle errors and recovery.

### 7: Free List: Recycle & Reuse
- Implement a linked list on disk for free pages.
- Implement free list operations (push, pop).
- Integrate free list with KV store.
- Manage page reuse.

### 8: Tables on KV
- Encode rows as key-value pairs.
- Implement database schemas.
- Support different column types.
- Implement table operations (get, insert, update, delete).
- Create internal tables for metadata.

### 9: Range Queries
- Implement B+tree iterator.
- Create order-preserving encoding for different data types.
- Implement range query scanner.

### 10: Secondary Indexes
- Implement secondary indexes as extra keys.
- Select appropriate index for queries.
- Maintain secondary indexes during updates.

### 11: Atomic Transactions
- Implement transaction interfaces (begin, commit, rollback).
- Achieve atomicity via copy-on-write.
- Move tree operations to transactions.
- Implement transactional table operations.

### 12: Concurrency Control
- Implement snapshot isolation for readers.
- Handle conflicts for writers.
- Implement optimistic concurrency control.
- Use version numbers in the free list.

### 13: SQL Parser
- Define query language specification.
- Implement recursive descent parser.
- Parse SQL-like statements (SELECT, INSERT, UPDATE, DELETE).
- Parse expressions with operator precedence.

### 14: Query Language
- Implement query execution engine.
- Convert abstract syntax tree to query plan.
- Implement plan nodes for different operations.
- Implement expression evaluation.
- Support aggregations and grouping.
