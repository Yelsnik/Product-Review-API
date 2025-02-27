# Product Review API

This platform allows users to enter reviews for products that they choose. Users can search for products on the platform and then enter a review for it. The API then ranks the reviews entered by the users for a particular product.

## Table of Contents
- Technologies Used
- Services
- API endpoints
- Getting Started

### Technologies Used

This API is built with languages like Go, NestJS and Python. Aside from these programming languages, it uses other technologies like gRPC and Redis. There are two main services; the nlp-service built with Python and the review service built with Go. Both services use gRPC for communication. The api gateway uses NestJS to route traffic to various services.

### Services
The two main services are;
1. Review service. This service handles the logic for the product review. It has the logic to add reviews, get reviews, get products for reviews and even see details of the product.

2. NLP service. This service contains the logic to analyse the reviews coming from the user. It is built with Python and it utilises some of Python's NLP libraries.

3. API gateway. This service more or less sets up HTTP endpoint for the gRPC services and routes traffic to the appropriate service.

### API endpoints

- Add Reviews

```
POST {url}/v1/reviews/:productid
```
request body:
```
reviews: string
```
- Get products

```
GET {url}/v1/reviews?page=1&country=us
```
query params:
```
page: string
country: string
```
- Get reviews for a product

```
GET {url}/v1/reviews/:productid
```
- Get product details

```
GET {url}/v1/reviews/product-details?productId=iwiosuhdj&country=us
```
query params:
```
productId: string
country: string
```

More endpoints to come...

### Getting started

To run the review service:

```
cd review-service
```
then

build and run:
```
docker build -t review-service \ 
docker run -p 6000:6000 review-service
```

To run the nlp service:

```
cd nlp-service
```
then

build and run:
```
docker build -t nlp-service \ 
docker run -p 6000:6000 nlp-service
```

To run all services together

from the root directory, run
```
docker compose up --build
```

