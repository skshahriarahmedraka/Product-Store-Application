
## Product Store application (rest API) backend using Golang(Gin).



Database Table Schema Design:

brands(id, name, status_id, created_at)
```
categories(id, name, parent_id, sequence, status_id, created_at)

suppliers(id, name, email, phone, status_id, is_verified_supplier, created_at)

products(id, name, description, specifications, brand_id, category _id, supplier_id, unit_price,
discount_price, tags, status_id)

product_stocks(id, product_id, stock quantity, updated_at)
```

### Task 1:
#### Create CRUD operations for brands, categories, suppliers, products, product_stocks. 


2. Create a product list API with Pagination. Please use the following filter criteria in this product list
API. The default product list format is unit price lower to high.

- Filter by product name

- Filter by product price range(min_price, max_price)
- Filter by multiple brands

- Filter by category

- Filter by supplier

- Filter by verified supplier

Note: Only Active products are shown in the list. And product stock quantity must be greater than 0.
3. Create a category tree based on category id and parent_id. Children layer must be in 3.
Note: Category tree order must be based on sequence column in database.
Response Example:

"data"

"id": 2
"category_name": "Mobile",
"children"

id": 2
"category_name": "10S"
"children

id": 3
"category_name": "Android"
‘children

"id": 4
"category_name": "Watch",
"children"

"id": §
"category_name": "Smart Watch
‘children

4. Let's consider a scenario where you are designing a music streaming service. The goal is to
implement an Least Recently Used (LRU) cache for storing recently played songs to optimize the
retrieval of frequently accessed songs. Here's the problem:

Problem: LRU Cache for Music Streaming

You are tasked with designing an LRU cache for a music streaming service. The service allows users to
play songs, and you want to optimize the retrieval of songs that are frequently played. The cache
should store information about the recently played songs, and when the cache reaches its capacity (for
example 100), it should remove the least recently played song to make room for new ones.

Note: Create a directory inside your golang project called problem_solving. Inside the
problem_solving directory create a file main.go for complete the LRU task.
Some Other Instructions:

- Dockerization is preferable but not mandatory
- Write documentation in your project in a README. md file

- Instruction about how to run the project

- All API Endpoints, Request Body, Query Params, and Response format
- Add .sql/mongo document file into the project directory