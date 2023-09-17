# LoD Mini Workshop Golang

## Introduction

This is a mini workshop for learning the LoD (Law of Demeter) in Golang.
You will learn by doing the exercises in this workshop.

## Setup

1. Clone this repository

   ```sh
   https://github.com/NuttapolCha/lod-workshop-go.git
   ```

2. Checkout to main branch

    ```sh
    git checkout main
    ```

3. You will coding at `./workshop/workshop.go`

## Getting Started

You are given the pre-defined structs and the function signature. You job is to implement the *HandlePersonPurchaseProducts* function that satisfying requirements;

### Requirements

A person can purchase a quantity of products if

1. The person has enough money, i.e. person's money after purchasing product must be greater than or equal to zero.
2. The product is in stock and has enough quantity.

### Submission

You can run the test script to check your implementation.

```sh
make test
```

### What's next?

Share your code with your team, discussing on maintainability, readability. Thinking of scenarios where new requirements are added.

> Is your code easy to change? Will your new code duplicate existing code? How can you improve your code?

### New Requirements

Checkout to `with-promotion` branch. You are given a new requirement that a person can purchase a quantity of products with promotion.
The promotion contains only discount field which is the fixed discount for purchasing products.

```sh
git checkout with-promotion
```

Try to implement the *HandlePersonPurchaseProductsWithPromotion* then try to test if your code satisfy the requirements.

```sh
make test
```

## Appendix

### LoD

The Law of Demeter (LoD) or principle of least knowledge is a design guideline for developing software, particularly object-oriented programs. In its general form, the LoD is a specific case of loose coupling. The guideline was proposed at Northeastern University towards the end of 1987, and can be succinctly summarized in each of the following ways:

- Each unit should have only limited knowledge about other units: only units "closely" related to the current unit.
- Each unit should only talk to its friends; don't talk to strangers.
- Only talk to your immediate friends.

#### Bad Example

```go
type Money struct {
	Amount int
}

type Wallet struct {
	Cash Money
}

type Person struct {
	Name   string
	Wallet Wallet
}

func (p Person) CanAfford(amount int) bool {
    return p.Wallet.Cash.Amount >= amount
}
```

#### Good Example

```go
type Money struct {
	Amount int
}

func (money Money) CanSpendBy(amount int) bool {
    return money.Amount >= amount
}

type Wallet struct {
	Cash Money
}

func (w Wallet) CanSpendBy(amount int) bool {
    return w.Cash.CanSpendBy(amount)
}

type Person struct {
	Name   string
	Wallet Wallet
}

func (p Person) CanAfford(amount int) bool {
    return p.Wallet.CanSpendBy(amount)
}
```