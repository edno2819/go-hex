# GO Hex

## Description:
A simple program implementation of RESTfull API based on principles of Hexagonal architecture (Ports and Adapters). This project is a part of my learning process in understanding [Hexagonal Architecture](https://alistair.cockburn.us/hexagonal-architecture/).

### At the end of this project, technology tools intended to be implemented:
Each external tool must have a corresponding adapter, connecting business functionality to technical implementation.

- **Redis**: Database for caching information.
- **PostgreSQL**: Primary database.
- **MongoDB**: A second option database to test drive principal.
- **Gin**: HTTP framework.
- ...

### Hexagonal Architecture Components:
#### Core
The core is where the heart of the application lives. It contains the essential business logic and rules of the application. This is where things like processing orders, managing user accounts, and performing all the tasks of the application is designed for. The core should be independent of any external technologies or frameworks, making it highly portable and reusable.

There are two other essential concepts in Hexagonal Architecture: "ports" and "adapters". These concepts dictates how the "core" interacts with the external components.

#### Ports
Think of ports as contracts or interfaces. They define how the application communicates with external systems or services. For example, a port could specify the rules for connecting to a database, interacting with another web services, or handling user interfaces. Ports belongs to the core, because the core defines which actions are required to achieve the business logic goals.

#### Adapters
Adapters are the ones who implement the contracts or interfaces defined by ports. Adapters are responsible for making sure the application can interact to databases, web services, or other things. They handle the technical details.

#### Driver Actors
Driver actors are the initiators of communication with the core. They reach out to the core to request specific services. Examples of driver actors can be HTTP request or command line interfaces (CLI).

#### Driven Actors
Driven actors are the ones that triggered by the core. If the core needs something from external services, it sends a request to the adapter, instructing it to perform a particular action. For example, if the core needs to store data in a Postgres database, it triggers communication with the Postgres client to execute an INSERT query. In this scenario, the core initiates the communication.