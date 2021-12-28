# SimpleRESTApi
This is application for Book Keeping which implements fully-fledged REST API  #We are using Gorilla Mux 

In total, there are 5 end points.

1. /book/           - POST - Creates a book
2. /book/           - GET -  Fetches all books
3. /book/{bookId}   - GET - Fetches book by id
4. /book/{bookId}   - PUT - Updates book by id
5. /book/{bookId}   - DELETE - Deletes book by id


Book Model:

- name        string
- author      string
- publication string
