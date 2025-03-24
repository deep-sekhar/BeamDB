In this project, we will implement a database from scratch in Go, step by step.

## Project Structure (Final Layout)
```
/your-db-project
    /cmd
        main.go
    /internal
        /storage
            btree.go
            freelist.go
            kv_store.go
        /table
            table.go
            index.go
        /transaction
            tx.go
        /sql
            parser.go
            executor.go
        /util
            encoding.go
    /pkg
        /db
            db.go
    go.mod
    README.md
```

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
