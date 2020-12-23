# Shopee Crawler API
[![goversion](https://img.shields.io/badge/Go-v1.14.4-blue)](https://golang.org/)
[![mysqlversion](https://img.shields.io/badge/mysql-v8.0.22-blue)](https://mysql.com/)

It is an API helps you crawl data from https://shopee.vn/ about **brands** & **products** and save them to database which you can access later by also this one.
# Table of contents
* [Installation](#installation)
* [Usage](#usage)
* [Test](#test)
* [Route Details](#route-details)
    * [Crawl Brands](#crawl-brands)
    * [Crawl Products](#crawl-products)
    * [Crawl Products of a Brand](#crawl-products-of-a-brand)
    * [Endpoint `brands`](#endpoint-brands)
        * [Create a Brand (POST)](#create-a-brand-post)
        * [Get All Brands (GET)](#get-all-brands-get)
        * [Get Detail a Brand (GET)](#get-detail-a-brand-get)
        * [Update a Brand (PUT)](#update-a-brand-put)
        * [Delete a Brand (DELETE)](#delete-a-brand-delete)
        * [Get All Products of a Brand (GET)](#get-all-products-of-a-brand-get)
    * [Endpoint `products`](#endpoint-products)
        * [Create a Product (POST)](#create-a-product-post)
        * [Get All Products (GET)](#get-all-products-get)
        * [Get Detail a Product (GET)](#get-detail-a-product-get)
        * [Update a Product (PUT)](#update-a-product-put)
        * [Delete a Product (DELETE)](#delete-a-product-delete)
# Installation
Fistly, we have to care about database migration. Supposed you're a MySQL user then:
```
migrate -database "mysql://username:password@tcp(yourhost:port)/databasename" -path migrations up
```
Folder `migrations` was prepared for purposes that you can easily `up` and `down` version of your database.  
Lastly, you need to config database in `env\database.env`. Here is an example:
```
DATABASE_CONFIGURED = "YES"
DATABASE_USERNAME = "MY_USERNAME"
DATABASE_PASSWORD = "MY_PASSWORD"
DATABASE_HOST = "LOCALHOST"
DATABASE_PORT = "3306"
DATABASE_MAINDATABASENAME = "MAIN"
DATABASE_TESTDATABASENAME = "TEST"
```
# Usage
Run your local server:
```
go run .
```
Result should be appear:
```
yyyy/mm/dd hh:mm:ss Hosting: Local server: localhost:8080
```
Running port is changeable in `env/server.env`.
```
SERVER_PORT = "8080"
```
# Test
Use the following commands:
```
go test github.com/dee-ex/shopee_crawler_api/modules/brands
go test github.com/dee-ex/shopee_crawler_api/modules/products
```
# Route Details
## Crawl Brands
Crawl all brands from https://shopee.vn/mall/brands and return as JSON.
```
/jobs/trigger/crawl_brands
```
## Crawl Products
Crawl all products of all brands from https://shopee.vn/each_brand_username and return as JSON (maximum return is `1000`).  
This process can be adjusted by query parameters.
```
/jobs/trigger/crawl_products?limit=X&from=Y&to=Z
```
## Crawl Products of a Brand
Crawl all products of a specific brand username.
```
/jobs/trigger/crawl_products/{brand_username}
```
## Endpoint `brands`
### Create a Brand (POST)
```
/brands
```
Create a new brand from data containing in body of request, save it to database and return as JSON.
```
{
    "Shopid": 999999,
    "Username": "testusername",
    "BrandName": "testbrandname",
    "Logo": "testlogo"
}
```
### Get All Brands (GET)
```
/brands
```
Get all brands and return as JSON.
### Get Detail a Brand (GET)
```
/brands/{brand_id}
```
Get detail a brand by ID and return as JSON.
### Update a Brand (PUT)
```
/brands/{brand_id}
```
Update a brand by ID. Updated data containing in body of request. Arguments are not required.
```
{
    "BrandName": "updatebrandname",
    "Logo": "testlogo"
}
```
### Delete a Brand (DELETE)
```
/brands/{brand_id}
```
Delete a brand by ID.
### Get All Products of a Brand (GET)
```
/brands/{brand_id}/products
```
Get all products of a specific brand and return as JSON.
## Endpoint `products`
### Create a Product (POST)
```
/products
```
Create a new product from data containing in body of request, save it to database and return as JSON.
```
{
    "Shopid": 999999
    "Itemid": 999998
    "PriceMax": 999997
    "PriceMin": 999996
    "Name": "testname"
    "Images": "testimage"
    "HistoricalSold": 999995
    "Rating": "testrating"
}
```
### Get All Products (GET)
```
/products
```
Get all products and return as JSON.
### Get Detail a Product (GET)
```
/products/{product_id}
```
Get detail a product by ID and return as JSON.
### Update a Product (PUT)
```
/products/{product_id}
```
Update a product by ID. Updated data containing in body of request. Arguments are not required.
```
{
    "Shopid": 999994
    "PriceMax": 999993
    "PriceMin": 999992
    "Name": "updatename"
    "Images": "updateimages"
    "HistoricalSold": 999991
    "Rating": "updaterating"
}
```
### Delete a Product (DELETE)
```
/products/{product_id}
```
Delete a product by ID.
