# REST vs GraphQL

A short, practical comparison and why teams often choose GraphQL.

## REST (quick)
- REST = Representational State Transfer: APIs around resources (e.g., users, posts, builds).
- Common methods: GET, POST, PUT, PATCH, DELETE. Responses are typically JSON.

Example:

GET /users/1

```json
{
  "id": 1,
  "name": "Parth",
  "email": "parth@email.com"
}
```

Problems with REST:
- Overfetching: server returns more fields than the client needs.
- Underfetching: multiple endpoints/requests needed to gather related data.

## GraphQL (quick)
- GraphQL is a query language + runtime that lets clients request exactly the shape of data they need.

Example query + response:

```graphql
query {
  user(id: 1) {
    name
    email
    posts {
      title
    }
  }
}
```

```json
{
  "data": {
    "user": {
      "name": "Parth",
      "email": "parth@email.com",
      "posts": [ { "title": "Post 1" }, { "title": "Post 2" } ]
    }
  }
}
```

Core GraphQL concepts (minimal):
- Schema: defines types and relationships (what fields exist).
- Query: the client request that specifies the desired shape of the response.
- Resolver: the function that retrieves the actual data for a field.

Minimal example (conceptual):

Schema (SDL):

```graphql
type User { id: ID!, name: String, email: String }
type Post { id: ID!, title: String }
type Query { user(id: ID!): User }
```

Query:

```graphql
{ user(id: 1) { name posts { title } } }
```

Resolver (pseudo):

```js
// Query.user resolver
function user(_, { id }) { return db.getUser(id); }
// User.posts resolver
function posts(user) { return db.getPostsByUser(user.id); }
```

N+1 problem (concise):
- If you request multiple builds and for each build ask for its triggering user, naive resolvers can issue one DB query per build (N+1 queries).
- Tools like DataLoader batch those lookups so you can fetch all needed users in a single batched query, reducing it to 1+1.

---

Keep it simple: use REST for straightforward resource endpoints; prefer GraphQL when clients need flexible, efficient data shaping and reduced round-trips.


