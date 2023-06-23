---
title: Clean Architecture and transactional consistency
marp: true

---

# **Clean Architecture and transactional consistency**
## How to manage multiple data model objects in a single request

</br></br></br>
Duc Trinh
duc.trinh@hcl.com


---
# Table of content
1. The problem
2. Clean Architecture & Domain-driven design
3. DDD Aggregate
4. Unit-of-work pattern

---
# 1. The problem

---

# Write single object in a single transaction
<img src="./diagrams/single-tx-script.png" style="width:80%;"/>

---

# Write multiple objects with multiple transactions
<img src="./diagrams/multi-tx-script.png" style="width:80%;"/>

---

# Write multiple objects with multiple transactions
<img src="./diagrams/multi-tx-script-fail.png" style="width:75%;"/>
Transport object is not persisted, and violates the business requirement.

---

# 2. Clean Architect & Domain-driven design

---

# Clean Architecture
![image](./clean-architecture.jpg)

---


# Domain-driven design (DDD)
<img src="./ddd.png" style="width:80%"/>

---

# DDD Entity
<p align="center"><img src="./domain-entity-pattern.png"></img></p>

---

# DDD Aggregate & Aggregate root

"A DDD aggregate is a cluster of domain objects that can be treated as a single unit. An example may be an order and its line-items, these will be separate objects, but it's useful to treat the order (together with its line items) as a single aggregate.

An aggregate will have one of its component objects be the aggregate root. Any references from outside the aggregate should only go to the aggregate root. The root can thus ensure the integrity of the aggregate as a whole."

[Martin Fowler](https://martinfowler.com/bliki/DDD_Aggregate.html)

---

# DDD Aggregate & Aggregate root
"Those entities that need to be transactionally consistent are what forms an aggregate. Thinking about transaction operations is probably the best way to identify aggregates."

[Microsoft article](https://learn.microsoft.com/en-us/dotnet/architecture/microservices/microservice-ddd-cqrs-patterns/microservice-domain-model#the-aggregate-pattern)

---

# Repository
<p align="center"><img src="./repository-aggregate-database-table-relationships.png" style="width:55%;justify-content:center"></p>

For each aggregate or aggregate root, you should create one repository class
[Microsoft article](https://learn.microsoft.com/en-us/dotnet/architecture/microservices/microservice-ddd-cqrs-patterns/infrastructure-persistence-layer-design#define-one-repository-per-aggregate)

---

# Repository & Unit-of-work
<p align="center"><img src="./repo_unit_of_work.png" style="width:45%;justify-content:center"></p>

[Microsoft article](https://learn.microsoft.com/en-us/aspnet/mvc/overview/older-versions/getting-started-with-ef-5-using-mvc-4/implementing-the-repository-and-unit-of-work-patterns-in-an-asp-net-mvc-application#the-repository-and-unit-of-work-patterns)

---

# Code example

---

# Order entity as the Aggregate root
![image](./code_order_entity.png)

---

# Transport entity
![image](./code_transport_entity.png)

---

# Use case create order
<img src="./usecase_create_order.png" style="width:55%" />

---

# MySQL repository
![image](./mysql_repo_1.png)

---

# MySQL repository
![image](./mysql_repo_2.png)

---

# MongoDB repository
![image](./mongo_repo.png)

---

# Pros

- Data consistency
- More confident on maintance and develop features without breaking business
- Easy testing

---

# Cons
- Over-engineering
- Complex
- Rely on domain expert

---

# 4. Unit-of-work pattern

---

# Unit-of-work

"A Unit of Work keeps track of everything you do during a business transaction that can affect the database. When you're done, it figures out everything that needs to be done to alter the database as a result of your work."
[Martin Fowler](https://martinfowler.com/eaaCatalog/unitOfWork.html)

---

# Unit-of-work pattern
![image](./uow-caller-registration.png)

Caller registration

---

# Unit-or-work pattern
![image](./uow-object-registration.png)

Object registration

---

# Unit-of-work pattern
<img src="./unit_of_work_impl_1.png" style="width:65%"/>

---

# 4. Unit-of-work pattern
<img src="./unit_of_work_impl_2.png" style="width:65%"/>

---

# Thank you!