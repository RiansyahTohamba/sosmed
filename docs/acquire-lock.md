acquire lock 

# command pada transaction
The command MULTI marks the beginning of a transaction, and the command
EXEC marks its end.

client.get(from, function(err, balance) { // 2
var multi = client.multi(); // 3
multi.decrby(from, value); // 4
multi.incrby(to, value); // 5

# acquire lock
teknik ini menjadi pre-requestid utk twitter case study.

WHY DOESN’T REDIS IMPLEMENT TYPICAL LOCKING? When accessing data for
writing ( SELECT FOR UPDATE in SQL ), relational databases will place a lock on
rows that are accessed until a transaction is completed with COMMIT or ROLL-
BACK . If any other client attempts to access data for writing on any of the same
rows, that client will be blocked until the first transaction is completed. 


# kenapa redis tidak implement mekanisme locking seperti rdbms lain?
pada tipikal rdbms, lock pada rows akan dilakukan pada query hingga transaction ditutup (either COMMIT or ROLLBACK).
proses lain tdk akan di izinkan untuk mengakses row tersebut, hingga transaction ditutup.
proses ini bagus sekali, tetapi akan terjadi long-wait-blocking jika transaction itu memang kompleks.
design dari redis adalah 'minimize wait time for clients', makanya redis tdk lock data during WATCH.
pada WATCH 




    This form of locking works well in practice (essentially all relational databases
    implement it), though it can result in long wait times for clients waiting to
    acquire locks on a number of rows if the lock holder is slow.



Because there’s potential for long wait times, and because the design of Redis
minimizes wait time for clients (except in the case of blocking LIST pops),
Redis doesn’t lock data during WATCH .
