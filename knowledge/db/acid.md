# What is ACID
A = Atomicity
C = Consistency
I = Isolation
D = durability

## Atomicity
**All Or Nothing**

Transaction must be **Fully Completed** or **Fully Rolled Back**. If one part fails, the whole transaction fails

**Example**

Transfering $100 from Foo to Bar
1. Debit Alice by 100
2. Credit Bob by 100
If credit to Bob fails, debit from alice must be rolledback

## Consistency
**Database must remain in valid state**

Transaction must transform database from one valid state to another.

What is invalid state? Invalid state means that the data breaks certain database rules

**Example**

**Rule**: Account balance cannot be negative

* Alice has $100 in her account
* Transaction A tries to send 150 dollar from Alice to Bob
* This should fail as the Alice balance would be negative


## Isolation
Concurrent transaction should not affect or interfere with each other. 

**example**

* Assume T1 and T2 running at the same time
* T1 substract $100 from Alice
* T2 reads Alice balance in the middle of T1 to do something
* **T2 might see inconsistent/partial data** such as Alice having 0 balance before T1 finishes
* T2 needs to wait for T1 to finish to continue as if T1 partial changes did not happen

## Durability
**Once transaction is committed, it stays**

Committed transactions should stay in persistent disk. In the event of a crash, the data should remain

**Example**

Transaction A successfully commit, then the system crash. When it reboots, the recorded value should not disappear

## Understanding them
**Isolation vs Consistency**
* Consistency focus on the validity of the final result. Does it follow the rules?
* Isolation focuses on the process of the transaction. Check if there is any conflict such as seeing uncommitted changes from another transactions

# ACID relationship to Isolation level
ACID is the rulebook of how transaction should be. While Isolation level affect the degree of **I** in ACID

| Property     | Meaning                                             | Ensures                                     | Can You Configure It?     |
|--------------|-----------------------------------------------------|---------------------------------------------|----------------------------|
| **A - Atomicity**   | All steps in a transaction succeed or none do      | No partial changes; rollback on failure     | ❌ No (handled by DB engine) |
| **C - Consistency** | Data remains valid according to rules/constraints  | No corruption; integrity is preserved       | ⚠️ Partially (via schema, triggers) |
| **I - Isolation**   | Transactions don’t affect each other’s execution   | No dirty/partial reads                      | ✅ Yes (via isolation levels) |
| **D - Durability**  | Once committed, data stays even after failure      | Crash-safe persistence                      | ❌ No (handled by DB and storage) |

What are the isolation levels and the tradeoff between them?

| Isolation Level        | What It Allows                                     | Speed   | Safe?   |
|------------------------|----------------------------------------------------|---------|---------|
| **Read Uncommitted**   | Can see **uncommitted** changes (dirty reads)     | 🚀 Fast | ❌ Risky |
| **Read Committed**     | Only sees **committed** data                       | ✅ Good | Standard |
| **Repeatable Read**    | Prevents changes to rows during a transaction     | 🔒 Safer| Slower  |
| **Serializable**       | Transactions run as if executed one at a time     | 🛡 Safest | 🐢 Slowest |

Starting from Read Committed, the transaction fulfils **Isolation** part of ACID

So,
* ACID says: Isolation must be enforced to prevent dirty, non-repeatable, or phantom reads.
* Isolation levels control how much isolation you want, depending on your app’s needs.

Higher isolation → More safety, less concurrency.

Lower isolation → More speed, but more risk of conflicts.

# Isolation Level and problem it tries to solve

Concurrency Problem
* **Dirty Read**: Reading data **not yet** committed by another transaction
* **Non-repeatable read**: Re-reading data returns different values
* **Phantom read**: Row appears / disappears when re-running the query

| Isolation Level     | Prevents                 | Allows                         | Main Problem It Solves       |
|---------------------|--------------------------|---------------------------------|-------------------------------|
| Read Uncommitted    | Nothing                  | Dirty Reads, Non-repeatable, Phantom | None                         |
| Read Committed      | Dirty Reads              | Non-repeatable, Phantom Reads  | Dirty Read                    |
| Repeatable Read     | Dirty + Non-repeatable   | Phantom Reads                  | Non-repeatable Read           |
| Serializable        | All 3 issues             | None                           | Phantom Read + Full Isolation |


## Read Uncommitted - Lowest Isolation
Problem it has: **Dirty Read**

Example: T1 updates Alice’s balance to $0 but hasn’t committed. T2 reads $0. T1 then rolls back. Now T2 has used wrong data.

**Why it is bad?**: The transaction basing its decision on unconfirmed data

## Read Committed - Default Isolation in Most DB
This solves **Dirty Read**

Other problem it has: **Non-repeatable Read**

```
T1 reads Alice’s balance = $500. T2 then updates it to $400 and commits.
T1 reads Alice’s balance again → now it’s $400.
➡️ Same query, different result.
```

**Why it is bad?**: Inconsistency during the same transaction

## Repeatable Read
This solves **Non-repeatable Read**

Other Problem it has: **Phantom Read**
```
T1 runs SELECT * FROM orders WHERE amount > 100.
T2 inserts a new matching row and commits.
T1 runs the same query again and sees new rows.
```

**Why it’s bad**: The result set changed even though T1 never modified anything.

## Serializable
This solves **Phantom Read**

Another example of Phantom Read issue
```
T1 checks if username is taken, DB returns everything is ok. T2 adds the the same username. T1 proceeds as if the username is not taken.
```
This breaks business logic

# Solving concurrency
If we do not want serializable due to performance reason, we can use Repeatable Read with UPDATE LOCK to lock the rows being queried. This prevents any insert or delete

Or we can just use a GAP Lock to lock the row in between them.




