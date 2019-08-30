---
title: "Paxos"
excerpt: "Scalable API and Dynamic Programming"
header:
  teaser: assets/images/paxos_01.jpg
---

{% include toc %}

## Info

We consider two programming assignments on REST API design and on dynamic programming skills. The first problem revolves around designing a scalable REST API with ability to handle large volume of requests. The second problem poses a twist to classic dynamic programming problems such as [*Knapsack Problem*](https://en.wikipedia.org/wiki/Knapsack_problem) and [*Coin Change Problem*](https://en.wikipedia.org/wiki/Change-making_problem).

For clarity of presentation, the two problem statements are explained below, whereas the solutions are explained in their corresponding Github repositories.

## Problem Statement

### Problem 1: REST API

Publish a small service on the web that has two endpoints:
1. `/messages` takes a message (a string) as a `POST` and returns the SHA256 hash digest of that message (in hexadecimal format)
2. `/messages/<hash>` is a `GET` request that returns the original message. A request to a non-existent `<hash>` should return a `404` error.

**Examples**

+ Assume the service is published to http://localhost:8080/.
+ Post a message
    ```bash
    $ curl -X POST -H "Content-Type: application/json" -d '{"message": "foo"}' http://localhost:8080/messages
    {
        "digest": "2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae"
    }
    ```
+ Check that the result is correct on the command line
    ```bash
    $ echo -n "foo" | shasum -a 256 2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae -
    ```
+ Query the service for the original message
    ```bash
    $ curl http://localhost:8080/messages/2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae
    {
        "message": "foo"
    }
    ```
+ Query for a non-existent `<hash>`
    ```bash
    $ curl -i http://localhost:8080/messages/aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa 
    {
        "err_msg": "Message not found"
    }
    ```

**Additional Questions**

+ What would the bottleneck(s) be in your implementation as you acquire more users? How you might scale your microservice?
+ How did you deploy the application and why did you choose to deploy it this way? How would you improve this deployment process if you needed to maintain it long term?

### Problem 2: Find Items

You have been given a gift card that is about to expire and you want to buy gifts for 2 friends. You want to spend the whole gift card, or if that’s not an option as close to the balance as possible. You have a list of sorted prices for a popular store that you know they both like to shop at. Your challenge is to find two distinct items in the list whose sum is minimally under (or equal to) the gift card balance.

The file contains two columns:
1. A unique identifier of the item. You can assume there are no duplicates.
2. The value of that item in cents. It is always a positive integer that represents the price in cents (1000 = $10.00).

Write a program to find the best two items. It takes two inputs:
1. A filename with a list of sorted prices
2. The balance of your gift card

If no two items have a sum that is less than or equal to the balance on the gift card, print “Not possible”. You don’t have to return every possible pair that is under the balance, just one such pair.

Note: There may be many rows in the file, so be sure to optimize your solution to scale.

**Examples**

```bash
$ cat prices.txt
Candy Bar, 500
Paperback Book, 700
Detergent, 1000
Headphones, 1400
Earmuffs, 2000
Bluetooth Stereo, 6000
```
```bash
$ find-pair prices.txt 2500
Candy Bar 500, Earmuffs 2000
```
```bash
$ find-pair prices.txt 2300
Paperback Book 700, Headphones 1400
```
```bash
$ find-pair prices.txt 10000
Earmuffs 2000, Bluetooth Stereo 6000
```
```bash
$ find-pair prices.txt 1100
Not possible
```

**Additional Questions**

+ What is the big O notation for your program?
+ You are considering giving gifts to more people. Instead of choosing exactly 2 items, allow for 3 gifts.

## Solution

### Solution 1: REST API
[Github Repository](https://github.com/Adaickalavan/paxos/tree/master/RESTApi)

### Solution 2: Find Items
[Github Repository](https://github.com/Adaickalavan/paxos/tree/master/FindItems)
