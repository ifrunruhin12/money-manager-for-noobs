# Money Manager for Noobs

A personal money management + live simulation backend app built for learning and actually fixing my own bad money habits.

## Why this exists

I’m pretty bad at managing money, so instead of just suffering like a normal person, I decided to build a system for myself.
There are already tons of finance apps out there, but this one is **personalized, minimal, and built while I learn backend engineering properly**.

This project is also me leveling up in:

* Go (Golang)
* Backend architecture
* Authentication systems (JWT)
* Databases (PostgreSQL)
* Clean layered design (handler → service → repo)

## Features (so far)

### Authentication

* User registration & login
* Password hashing with bcrypt
* JWT-based auth system
* Protected routes with middleware

### Core Finance System

* User account with balance tracking
* Transaction system (planned/ongoing)
* Category-based spending structure:

  * Food
  * Transport
  * Health
  * Savings
  * Hobby
  * Big Buy
  * Extra Food

### System Design

* Clean layered architecture
* Repository pattern for DB access
* Service layer for business logic
* PostgreSQL with migrations
* Transaction-safe operations

## Tech Stack

* Go (Golang)
* PostgreSQL
* Gin (HTTP framework)
* JWT (auth)
* bcrypt (security)
* Docker (deployment/dev setup)

## Current Status

* Phase 0 (Authentication system) (mostly done)
* Backend API running with Docker
* Database migrations automated
* Ready to build Phase 1 features (transactions, analytics, etc.)

## Goal

Build a **fully personalized finance + simulation system** that helps me:

* Understand my spending behavior
* Track money properly
* Experiment with financial rules and simulations
* And obviously… become better at backend engineering

## Notes

This is not a production-grade fintech app (not yet).
It’s a learning-driven system that evolves as I level up.
