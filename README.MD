Implement a **Basic URL Shortener** with Future Scalability Considerations in Golang.

The service should generate short aliases for long URLs and be capable of mapping those aliases back to the original URLs.

Functional Requirements
1. Implement two endpoints:
- encode - Accepts long URLs as input and generates short aliases for them.
- decode - Accepts short aliases as input and redirects users to the original URLs.
2. Error Handling: Make sure to implement proper error handling

### How to run my solution
1. Open a terminal
2. Run the command go
```bash
go run cmd/basicUrlShortener/main.go
```
This will start the server and it will begin listening for HTTP requests on port 8080.
3. Open a browser and go to the following URL
```bash
curl "http://localhost:8080/encode?url=http://example.com"
```
This will return a response with a short URL.
4. Open a browser and go to the following URL
```bash
curl "http://localhost:8080/decode/aHR0cDovL2V4YW1wbGUuY29t"
```
This will redirect to the original long URL (http://example.com in this case).

### Describe
Developing a basic URL shortening service in Golang involves creating two HTTP handlers: "Encode" for shortening URLs and "Decode" for resolving the original addresses. The "Encode" function encodes a long URL into a short alias using base64 encoding and stores the mapping in memory. "Decode" takes an alias, retrieves the URL from memory, and redirects the user. Error handling includes sending HTTP responses for invalid URLs or missing mappings in memory. Drawbacks include temporarily storing URLs in memory, which is inefficient for production use, and basic error handling that doesn't cover all possible scenarios.

### How I would improve it further
If this were a real project, key improvements would include using persistent storage for storing URL addresses, ensuring the creation of unique short URLs, adding rate limiting to prevent abuse, as well as implementing usage analytics for URL tracking. Other important aspects include error handling and logging, extensive testing, enhancing security and documentation, as well as planning deployment and scaling of the service.

### Math & Logic behind “short url” 
When you enter a long link into a URL shortening service, it creates a short link for it. This happens using a special math function that turns the long link into a short one and vice versa. For example, you can simply take the long link and "encode" it to get the short one. The problem is that this doesn't always make the link really shorter, and they can repeat. A more reliable way is to assign each link its unique number and use it to create a short link. This ensures that each short link is unique and linked to its long version. Once the short link is created, it is stored in the database along with the long version. When you use the short link, the service checks the database and redirects you to the corresponding long link.


