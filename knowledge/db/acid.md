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
