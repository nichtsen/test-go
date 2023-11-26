stack and queque(channel)

b b+ red-black

thread vs Groutine

1. channel:
   1.1 defination
   Interactions between Groutines, either buffered or unbuffered.
   two side of channel : receiver and producer
   1.2 blocked
   In the case of unbuffered, the two side groutines of channel will not be blocked only if the buffer size is enough at present
   1.3 asnyc
   Asynchronous processing:receiver can put the channel aside and do something else until it need the message

#redis note

## history of handling large amount of io

Bio -> Nio -> 多路复用 (select/poll) -> e-poll
p.s aio?

// bio
fd1= socket()
bind(PORT)
listen()
while(1) {
cfd = accept(fd1)
clone(read(cfd)) //make thread to read in case of blocking the main thread
}

// nio
fd1= socket()
bind(PORT)
listen()
while(1) {
cfds = accept(fd1)
read(cfd1) //non-block
}

//select/poll
fd1= socket()
bind(PORT)
listen()
*cfds = select(fd1)
for cfd in range *cfds {
read(cfd1)  
}

//epoll (event notification & DMA<direct memory access>)
efd = epoll_create()
fd1= socket()
fd = epoll_ctl(efd, fd1)
bind(PORT)
listen()

while(1) {
*infds = epoll_wait()
for infd in range *infds {
swich infd {
case fd1:
cfd = accept(fd1)
epoll_ctl(efd, cfd)
default:
read(infd)
}
}
}

## local method

computing to data instead of data to computing

## thread trace

strace -ff -o ~/out ./redis-server

#nginx
##zero copy
sendfile(socket, file)

##linux kernal

GDT(global descriptor table) --> kernal space ; user space......
soft interrupt

2020 年 7 月 26 日 17 点 30 分
cache solution:
http server's native cache solution: e. g apache httpd.conf;
linux kernal?;
programming language's module:
java: jbossCache/treeCache;
golang: go-cache go2cache;
load balance;
database cluster:
LBC, HPC, HAC;
e .g of mysql: m-m-slave as a cluster, in addtion to a global router database \
 p.s: auto_increment is forbbiden;none pconnect(php);
to sperate image server;
库表散列;

interview map:
fundamental:
data structure: b-tree, b+tree, red-black tree;
algorithms:
sort: heapsort, selectionsort, insertionsort, quiksort, bubblesort, mergesort
map:
key value pair
key is unique;
hashtable;
network: socket, rpc
facility:
os: linux kernal;
linux kernal:
io
io:
nio;
bio;
aio;
database:
sql or nosql;
disk and memory;
table splite:
horizontal:
by id, hash(id) % n, datatime;
vertical:
extended table
database split:
service
consistent:
PCC(Pessimistic Concurrency Control);
OCC(Optimistic Concurrency Control):
CAS(Compare and Swap)
MCC/MVCC(Multiverion Concurrency Control);
compiler
concepts:
multi-process:
concurrency 同时处理:
context switch
parallelism 同一时刻执行:
lock versus Compare And Swap
unit: process, threads, coroutine：
process: unit of resource allocationg
thread: unit of software executing
programming language category: compiler, interpreter, virtual machine
golang:
goroutine:
Communicating Sequential Process(CSP);
synchronization:
low-level: sync.Mutex, sync.WaitGroup;
high-level:
channel:
select and timeout:
dafault: time.After()
dataProducer and dataReceiver:
close(channel ch): return zore value and false
inheritance: can not overide
middleware: message cache search filestore
单体应用:ERP CRM SCM HR PLM MES
rpc framework:
ailibaba: Dubbo/Dubbox;
Google gRPC;
Spring Boot/Spring Cloud;
Facebook: Thrift
Redis：
single process
data type： string hash list set zset
persistence:
RDB: copy on write
AOF: collection of commands
? io threads and thread pool
gin：
router：a prefix-tree via struct node
engine:
trees []methodtree:
struct methodTree{
method string
root \*node
}

## mysql

- index: [clustered index and secondary index](https://dev.mysql.com/doc/refman/5.7/en/innodb-index-types.html),
  [usage of b+tree](https://www.callibrity.com/blog/database-index-usage-of-b-tree-in-the-practical-database-system#:~:text=InnoDB%20Usage%20of%20B%2B%20tree&text=As%20can%20be%20seen%2C%20each,shown%20in%20MyISAM%20storage%20engine.)

b+tree and range query(mysql)
b-tree and internal node that holds data(mongodb)

- covering index: avoid transmission from seconndary index to cluster index

- engine: innodb myIASM

  - innodb: raw data in leaf node
  - mylASM: address of the record in leaf node

- cache index: [buffer tool](https://dev.mysql.com/doc/refman/8.0/en/innodb-buffer-pool.html#:~:text=The%20buffer%20pool%20is%20an,memory%2C%20which%20speeds%20up%20processing.)

distributed storage:
distributed strategy:
partitioned by id range (in tikv as a region);
N mod hash(id)；
consistent hashing：
32bits, clockwise;
memcache
consensus between replica:
raft, proxy, PoW, PoS
容错单元
volatile:
viewable inter processes;
cup commands is serializable;

    go:
        ?multi-core cpu programming;
        singleton and init()
    api:
        soa
        restful
        web service
        rpc

dynamic programming  
#include<cstdio>
#include<iostream>
#include<algorithm>
using namespace std;
const int n = 100;
int A[n], dp[n];
int main(void){
int n;
scanf("%d", &n);
for(int i = 1; i <= n; ++i){
scanf("%d", &A[i]);
}
int ans = -1;
for(int i = 1; i <= n; ++i){
dp[i] = 1;
for(int j = 1; j < i; ++j){
if(A[i] >= A[j] && (dp[j] + 1 > dp[i])){
dp[i] = dp[j] + 1;
}
}
ans = max(ans, dp[i]);
}
printf("%d", ans);
return 0;
}

等额本息：
期末年金模型：
https://wenku.baidu.com/view/3495dfc773fe910ef12d2af90242a8956aecaa12.html?fr=search-4-income3
等额本金

zombie process:
if pid > 0 then father process
wait() waitpid()
signal(SIGCHLD, SIG_LGN)-> init
double fork() grandson process ->init

recursion and iteration:
recursion:
stack increases with deepth of recursion
iteration:
can implement a recursion like process by maintain a stack within a single function
depth-first search and breadth first search:
search is two dimensions path;
dfs:
implemented by stack no matter in form of recursion or iteration;
FILO
bfs:
implemented by queue;
FIFO
inter-process communication(IPC):
implemented by kernal:
shared memory;
message queue:
linked list;
semaphore;
signal
pipe:
non-name pipe:
a buffer in kernal space;
only relative process can share a common pipe
FIFO:
non-relative process can communicate with FIFO;
saved in disk;
socket;
c/s architechture:
partition:
horizontal;
vertical
distribution:
load balance:
allocation:
polling;
amount of request;
speed of request;
hashing;
health check:
heatbeat;
type:
replication:
load balance:
stateless web server
partition
load balance:
consistent hashing
Asynchronous:
message queque;
channel
cache:
local cache;
distributed cache;
CDN;
inverted proxy;
computaional process:
prefix notaion;
pretty-print;
happens after;
define is our language’s simplest means of abstraction
operator and operand:
analogy: function and parameters

?烟囱模型的应用间调用
如摩根大通等某些数字化转型的鉴定执行者，会引入占员工比例 15%甚至
20%的技术人员，并直接派驻到业务部门与之共同工作。

云计算 云平台

e programmer must seek both perfection
of part and adequacy of collection

# pacsal

using braces as block of statement instead of brackets

linear recursive

tree recursive

game server protocols:
tcp:
udp:
moba:
base on frames;
websocket:
upgrade from http;
ping pong to keep connection
tcp ack seq
little-hero:
how interceptor router a message to its handler;
3 goroutines per client:
read write task
client is a middleman between websocket connection and hub;
hub is responsible for register and unregister client

register machine:
speration of data-path and controller sequence.
a continuation register is a special register that save a label.
it is more ecomonical to sharing continuation register rather than creating a new register for every subroutine.
sharing could be problematic if a subroutine has its own subroutine, namely nested subroutine.
hence a stack structure is ideal for continuation register.
in the process of recursion, both data-path and controller sequence can be reused, only register need to be replicated via stack.

    null? predict empty list
    stack operations of multiple register, the strict order of save and restore
    metacircular scheme as a underlaying scheme for evaluation

    stop-and-copy gc: where the non-pair memory located, neither in working memory, nor in free memory
    how do gc judge if a node of memory is to be discarded or still reserved.

    a register machine is capble of evaluatine of a composite primimtive process that operated on composite data strcuture accommodated
    itself on the memory.


    TODO: label, stack and storage management

    both stack and register is a container of abstraction
    environment is a sequence of frame, whih each frame contain a table of variables and values.
    in a register macine language, 'perform' is an explicit action of option.

environment model:
interplay of eval and apply
the applying of compound procedure is evaluating of sequence
all definiations within a frame should occur simultaneously

x1 x2=f(x1) x3=f(x2) x4=f(x3)  
 x1 x2=f(x1) x3=f(x2)
x1 x2=f(x1)
x1

The body of a procedure body is a sequence of expressions to be evaluated, since these expressions are in
a state of 'to be finished', the procedure could be regarded as facility of delayed evalutation.

In a tail-recursion case, the last expression of procedure body would be evaluated without saving its environment
to stack as a prerequisite, the consequential recursive calling would cover its previous environment, thus will not request
incrementing storage.

the explicit-evaluator is a machine whose input is source language, explicit indicate that its way of control of flow
is based on label jumping

fp redex
tex expandable
scip subexpression

lambda calculus
diff between define and let

tail recursion and list args
optimization of explicit-evaluator, hint. memorization

dotted-tail noatation

three types of linkage
return next label

copy constructor

substiution model vs environment model

selector constructor predictor
extract make classify

a compiler implementation is separated from implemented language(object), from aspective of object language,
the implementation of its compiler is out of scope. in other words, its compiler is regarded as a facility.
such a separation give great flexibility to implementation of the compiler. so many possibility exists...
a software engineer hence shall never trapped in a single facility, he always prefer to work on a horizon
beyond any certain set of APIs. such a horizon liberate himself from a colored glasses that he may have wore
on with no intention.

#TODO: ast and interpretor

heap and memory

l-(l%r)

## tcp

[link](https://learn.microsoft.com/en-us/troubleshoot/windows-server/networking/three-way-handshake-via-tcpip)

- three way handshake
- control bits:
  URG: Urgent Pointer field significant
  ACK: Acknowledgment field significant
  PSH: Push Function
  RST: Reset the connection
  SYN: Synchronize sequence numbers
  FIN: No more data from sender
- full-duplex
- 3 packets for establishing a connection, both synchronizations need to be acknowledged
- 4 packets for endinng a connection, both side need a FIN to be acknowledged

## redis

- data types: string, hash, list, set, zset; geospatial, hyperloglog, bitmap
- persistence: rdb, aof

## golang

- schedule: GMP model:
  - local G queque
  - global G queque
- garbage collector: triple color mark and swap:
  - white: dead
  - black: active
  - gray: child node may be white
  - dfs

## leetcode::daul pointer

- two sum problem: two opposite sides and moving towards each other
- moving window: start on the side and moving towards the same diection

## leetcode::dfs and bfs search

- search tree
  - preorder
  - inorder
  - postorder
- search graph(2 dimentional array)
  - four directions
  - dfs: stack(FILO) or recursion
  - bfs: queque

## leetcode::lis and lcs

    in both longest increasing subsequence and longest commone subsequence,
    evey entry of a dp array, no matter it is 1 dimentional or 2,
    is the solution of the subproblem it represents.
    in the LIS problem, the transaction funciton is:
    dp[i] = max(dp[j]+1, dp[i]) (j<i)
    in the LCS problem:
    dp[i][j] = dp[i-1][j-1] + 1 (extensible);
    dp[i][j] = max(dp[i-1][j], dp[i][j-1]) (inheritance from previous)

# binary packing problem

- with input array num []int
  for every item in the pack, pick it or not, the solution is an array of binary solution []bool
- transfer expression:
  dp[i][j] = max(dp[i-1][j], dp[i-1][j-w]), w is the weight of ith item
