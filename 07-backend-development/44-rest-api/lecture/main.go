package main

import (
	"fmt"
	"net/http"
	"strconv"
)

/*
============================================================
44 — REST API
============================================================

Topics
------------------------------------------------------------
1. REST
2. Resources
3. HTTP Methods
4. CRUD
5. Collection vs Individual Resource
6. Path Parameters
7. Query Parameters
8. Path vs Query Parameters
9. DTOs
10. Validation
11. Pagination
12. Filtering
13. Sorting
14. Combining Query Parameters
15. REST API Architecture
16. Complete Router
============================================================
*/

// ============================================================
// 1. REST
// ============================================================

/*
REST = Representational State Transfer

REST is an architectural style for designing APIs.

A REST API uses HTTP concepts to work with resources.

The basic mental model:

    URL
        ↓
    Resource

    HTTP Method
        ↓
    Operation

    Path Parameter
        ↓
    Specific Resource

    Query Parameter
        ↓
    Filtering / Pagination / Sorting


Example:

    GET /users/123

    GET
        → operation: read

    /users
        → resource: users

    123
        → specific user
*/

// ============================================================
// 2. RESOURCES
// ============================================================

/*
REST URLs should normally represent resources.

Good:

    /users
    /users/123

    /products
    /products/10

    /orders
    /orders/500


Avoid action-based URLs:

    /getUsers
    /createUser
    /deleteUser
    /updateUser


Why?

Because the HTTP method already represents the operation.

For example:

    GET /users

means:

    "Get users."


    POST /users

means:

    "Create a user."
*/

// ============================================================
// 3. HTTP METHODS
// ============================================================

/*
Common HTTP methods:

    GET
        Read resource(s)

    POST
        Create a resource

    PUT
        Replace a resource

    PATCH
        Partially update a resource

    DELETE
        Delete a resource
*/

// ============================================================
// 4. CRUD
// ============================================================

/*
CRUD:

    Create
    Read
    Update
    Delete


REST mapping:

    CREATE

        POST /users


    READ COLLECTION

        GET /users


    READ ONE

        GET /users/123


    UPDATE COMPLETELY

        PUT /users/123


    UPDATE PARTIALLY

        PATCH /users/123


    DELETE

        DELETE /users/123
*/

// ============================================================
// 5. COLLECTION VS INDIVIDUAL RESOURCE
// ============================================================

/*
These are different resources:

    /users

        Collection of users.


    /users/123

        One specific user.


Think:

    /users
        ↓
    "users collection"


    /users/123
        ↓
    "user 123"
*/

// ============================================================
// 6. PATH PARAMETERS
// ============================================================

/*
A path parameter identifies a specific resource.

Example:

    GET /users/123


    /users
        ↓
    resource


    123
        ↓
    resource identifier


Go 1.22+ supports path parameters in ServeMux:

    /users/{id}


We can retrieve it with:

    r.PathValue("id")
*/

func getUserHandler(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	fmt.Fprintf(w, "User ID: %s\n", id)
}

// ============================================================
// 7. MULTIPLE PATH PARAMETERS
// ============================================================

/*
A route can contain multiple path parameters.

Example:

    GET /users/10/orders/500


Route:

    /users/{userID}/orders/{orderID}
*/

func getUserOrderHandler(w http.ResponseWriter, r *http.Request) {

	userID := r.PathValue("userID")
	orderID := r.PathValue("orderID")

	fmt.Fprintf(w, "User ID: %s\n", userID)
	fmt.Fprintf(w, "Order ID: %s\n", orderID)
}

// ============================================================
// 8. QUERY PARAMETERS
// ============================================================

/*
Query parameters control how a collection is retrieved.

Example:

    GET /users?page=2&limit=20


Query string:

    page=2
    limit=20


In Go:

    query := r.URL.Query()

    page := query.Get("page")
    limit := query.Get("limit")
*/

func queryExampleHandler(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	page := query.Get("page")
	limit := query.Get("limit")
	search := query.Get("search")

	fmt.Fprintf(w, "Page: %s\n", page)
	fmt.Fprintf(w, "Limit: %s\n", limit)
	fmt.Fprintf(w, "Search: %s\n", search)
}

// ============================================================
// 9. PATH PARAMETER VS QUERY PARAMETER
// ============================================================

/*
This distinction is very important.


PATH PARAMETER
--------------

    GET /users/123

Means:

    "Give me user 123."


QUERY PARAMETER
---------------

    GET /users?role=admin

Means:

    "Give me users filtered by role."


Another example:

    GET /users?page=2&limit=20


The path:

    /users

identifies the collection.


The query:

    page=2&limit=20

controls how we retrieve it.


Mental model:

    Path parameter
        → WHICH resource?


    Query parameter
        → HOW should I retrieve the collection?
*/

// ============================================================
// 10. DTO
// ============================================================

/*
DTO = Data Transfer Object

A DTO represents data transferred between parts of a system.

For example, when creating a user, the client might send:

    name
    email


We can represent this with:

    CreateUserDTO
*/

type CreateUserDTO struct {
	Name  string
	Email string
}

// ============================================================
// 11. REQUEST DTO VS DOMAIN MODEL
// ============================================================

/*
Don't automatically use your domain/entity model as
your HTTP request model.

For example:

    User

might contain:

    ID
    Name
    Email
    PasswordHash
    CreatedAt
    UpdatedAt


The client should NOT necessarily control all of these fields.


Instead:

    CreateUserDTO

contains only fields that the client is allowed to provide:

    Name
    Email
*/

type User struct {
	ID           string
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    string
	UpdatedAt    string
}

// ============================================================
// 12. VALIDATION
// ============================================================

/*
After receiving a request, validate it.

Example:

    Name cannot be empty.

    Email cannot be empty.

    Page must be >= 1.

    Limit must be between 1 and 100.


Typical flow:

    HTTP Request
          ↓
       Decode
          ↓
      Validate
          ↓
       Service
          ↓
      Repository
          ↓
       Database
*/

func validateCreateUser(dto CreateUserDTO) error {

	if dto.Name == "" {
		return fmt.Errorf("name is required")
	}

	if dto.Email == "" {
		return fmt.Errorf("email is required")
	}

	return nil
}

// ============================================================
// 13. PAGINATION
// ============================================================

/*
Suppose we have:

    1000 users


Returning all 1000 users isn't always desirable.

Instead:

    GET /users?page=2&limit=20


We divide users into pages.


Page 1:

    users 1 - 20


Page 2:

    users 21 - 40


Page 3:

    users 41 - 60
*/

// ============================================================
// 14. LIMIT AND OFFSET
// ============================================================

/*
Databases commonly use:

    LIMIT
    OFFSET


Example:

    page  = 2
    limit = 20


Formula:

    offset = (page - 1) * limit


Therefore:

    offset = (2 - 1) * 20
           = 20


Database:

    LIMIT 20
    OFFSET 20
*/

func calculateOffset(page, limit int) int {

	return (page - 1) * limit
}

// ============================================================
// 15. PAGINATION DEFAULTS
// ============================================================

/*
We should define defaults.

For example:

    page  = 1
    limit = 20


And protect the API:

    page < 1
        → page = 1


    limit < 1
        → limit = 20


    limit > 100
        → limit = 100
*/

func normalizePagination(page, limit int) (int, int) {

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 20
	}

	if limit > 100 {
		limit = 100
	}

	return page, limit
}

// ============================================================
// 16. PAGINATION QUERY DTO
// ============================================================

type PaginationQuery struct {
	Page   int
	Limit  int
	Offset int
}

// ============================================================
// 17. FILTERING
// ============================================================

/*
Query parameters can filter collections.

Examples:

    GET /users?role=admin


    GET /users?status=active


    GET /users?search=mahdi


Multiple filters:

    GET /users?role=admin&status=active


Search + filter:

    GET /users?search=mahdi&role=admin
*/

// ============================================================
// 18. SORTING
// ============================================================

/*
Sorting can also use query parameters.

Example:

    GET /users?sort=name


We can define:

    sort=name

as ascending.


And:

    sort=-name

as descending.


Example:

    GET /users?sort=-createdAt


means:

    createdAt DESC
*/

// ============================================================
// 19. LIST USERS QUERY
// ============================================================

/*
A realistic users endpoint might support:

    page
    limit
    search
    role
    status
    sort


Instead of passing these independently, group them.
*/

type ListUsersQuery struct {
	Page   int
	Limit  int
	Offset int

	Search string
	Role   string
	Status string
	Sort   string
}

// ============================================================
// 20. PARSE LIST USERS QUERY
// ============================================================

func parseListUsersQuery(r *http.Request) ListUsersQuery {

	query := r.URL.Query()

	page := 1
	limit := 20

	if value := query.Get("page"); value != "" {

		if parsed, err := strconv.Atoi(value); err == nil {
			page = parsed
		}
	}

	if value := query.Get("limit"); value != "" {

		if parsed, err := strconv.Atoi(value); err == nil {
			limit = parsed
		}
	}

	page, limit = normalizePagination(page, limit)

	offset := calculateOffset(page, limit)

	return ListUsersQuery{
		Page:   page,
		Limit:  limit,
		Offset: offset,

		Search: query.Get("search"),
		Role:   query.Get("role"),
		Status: query.Get("status"),
		Sort:   query.Get("sort"),
	}
}

// ============================================================
// 21. LIST USERS EXAMPLE
// ============================================================

func listUsersHandler(w http.ResponseWriter, r *http.Request) {

	params := parseListUsersQuery(r)

	fmt.Fprintf(w, "Page: %d\n", params.Page)
	fmt.Fprintf(w, "Limit: %d\n", params.Limit)
	fmt.Fprintf(w, "Offset: %d\n", params.Offset)

	fmt.Fprintf(w, "Search: %s\n", params.Search)
	fmt.Fprintf(w, "Role: %s\n", params.Role)
	fmt.Fprintf(w, "Status: %s\n", params.Status)
	fmt.Fprintf(w, "Sort: %s\n", params.Sort)
}

/*
Example request:

    GET /users?page=2&limit=10&search=mahdi&role=admin&sort=-createdAt


Result:

    Page: 2
    Limit: 10
    Offset: 10

    Search: mahdi
    Role: admin
    Status:
    Sort: -createdAt
*/

// ============================================================
// 22. CRUD HANDLERS
// ============================================================

func getUsersHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Get users")
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Create user")
}

func replaceUserHandler(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	fmt.Fprintf(w, "Replace user: %s\n", id)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	fmt.Fprintf(w, "Partially update user: %s\n", id)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	fmt.Fprintf(w, "Delete user: %s\n", id)
}

// ============================================================
// 23. COMPLETE REST ROUTER
// ============================================================

func createRouter() *http.ServeMux {
	mux := http.NewServeMux()

	/*
		------------------------------------------------------------
		COLLECTION
		------------------------------------------------------------
	*/

	mux.HandleFunc(
		"GET /users",
		getUsersHandler,
	)

	mux.HandleFunc(
		"POST /users",
		createUserHandler,
	)

	/*
		------------------------------------------------------------
		INDIVIDUAL RESOURCE
		------------------------------------------------------------
	*/

	mux.HandleFunc(
		"GET /users/{id}",
		getUserHandler,
	)

	mux.HandleFunc(
		"PUT /users/{id}",
		replaceUserHandler,
	)

	mux.HandleFunc(
		"PATCH /users/{id}",
		updateUserHandler,
	)

	mux.HandleFunc(
		"DELETE /users/{id}",
		deleteUserHandler,
	)

	/*
		------------------------------------------------------------
		USER ORDERS
		------------------------------------------------------------
	*/

	mux.HandleFunc(
		"GET /users/{userID}/orders/{orderID}",
		getUserOrderHandler,
	)

	return mux
}

// ============================================================
// 24. COMPLETE REQUEST EXAMPLES
// ============================================================

/*
CREATE

    POST /users


READ COLLECTION

    GET /users


READ ONE

    GET /users/123


REPLACE

    PUT /users/123


PARTIAL UPDATE

    PATCH /users/123


DELETE

    DELETE /users/123


PAGINATION

    GET /users?page=2&limit=20


FILTER

    GET /users?role=admin


SEARCH

    GET /users?search=mahdi


SORT

    GET /users?sort=name


DESCENDING SORT

    GET /users?sort=-createdAt


COMBINED

    GET /users
        ?page=2
        &limit=20
        &search=mahdi
        &role=admin
        &status=active
        &sort=-createdAt
*/

// ============================================================
// 25. REST API ARCHITECTURE
// ============================================================

/*
A clean backend usually separates responsibilities:

                    HTTP
                      │
                      ▼
                   Handler
                      │
          ┌───────────┴───────────┐
          │                       │
     Path Params            Query Params
          │                       │
          └───────────┬───────────┘
                      │
                      ▼
                     DTO
                      │
                      ▼
                  Validation
                      │
                      ▼
                   Service
                      │
                      ▼
                 Repository
                      │
                      ▼
                   Database


Handler:

    - HTTP concerns
    - path parameters
    - query parameters
    - request body
    - HTTP status
    - HTTP response


Service:

    - business logic
    - business rules
    - use cases


Repository:

    - database access
    - persistence
    - SQL
*/

// ============================================================
// 26. IMPORTANT DESIGN RULE
// ============================================================

/*
Don't put everything in the HTTP handler.

Bad:

    func getUsersHandler(...) {

        // parse query

        // validate

        // business logic

        // SQL query

        // filtering

        // sorting

        // pagination

        // database

        // response
    }


Better:

    Handler
       ↓
    Service
       ↓
    Repository
       ↓
    Database
*/

// ============================================================
// 27. MAIN
// ============================================================

func main() {

	mux := createRouter()

	/*
		For demonstration, the list-users example is available
		separately because /users is already used by the REST
		collection routes.
	*/

	mux.HandleFunc(
		"GET /list-users",
		listUsersHandler,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server running on http://localhost:8080")

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Server error:", err)
	}
}

/*
============================================================
REST API MENTAL MODEL
============================================================

Question 1:

    What resource?

Answer:

    URL


Question 2:

    What operation?

Answer:

    HTTP Method


Question 3:

    Which specific resource?

Answer:

    Path Parameter


Question 4:

    How should the collection be retrieved?

Answer:

    Query Parameters


Question 5:

    What data crosses the HTTP boundary?

Answer:

    DTO


============================================================
EXAMPLE
============================================================

    PATCH /users/123?notify=true


    PATCH
        ↓
    Operation


    /users
        ↓
    Resource


    123
        ↓
    Specific Resource


    notify=true
        ↓
    Query Parameter


    Request Body
        ↓
    Request DTO


============================================================
END OF 44 — REST API
============================================================
*/
