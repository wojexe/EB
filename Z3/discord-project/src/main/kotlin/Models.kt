package com.example

import kotlinx.serialization.Serializable

@Serializable
data class Category(val name: String)

@Serializable
data class Product(
    val name: String,
    val price: Double,
    val category: Category
)

@Serializable
data class MessageRequest(
    val content: String
)

val CATEGORIES = listOf(
    Category("Electronics"),
    Category("Books"),
    Category("Clothing"),
    Category("Food & Beverages"),
    Category("Home & Garden")
)

val PRODUCTS = listOf(
    // Electronics
    Product("Smartphone", 899.99, CATEGORIES[0]),
    Product("Laptop", 1299.99, CATEGORIES[0]),
    Product("Headphones", 199.99, CATEGORIES[0]),

    // Books
    Product("Novel", 15.99, CATEGORIES[1]),
    Product("Textbook", 79.99, CATEGORIES[1]),
    Product("Comic Book", 9.99, CATEGORIES[1]),

    // Clothing
    Product("T-shirt", 19.99, CATEGORIES[2]),
    Product("Jeans", 49.99, CATEGORIES[2]),
    Product("Jacket", 99.99, CATEGORIES[2]),

    // Food & Beverages
    Product("Coffee", 4.99, CATEGORIES[3]),
    Product("Chocolate Bar", 2.99, CATEGORIES[3]),
    Product("Pasta", 3.49, CATEGORIES[3]),

    // Home & Garden
    Product("Plant", 24.99, CATEGORIES[4]),
    Product("Lamp", 39.99, CATEGORIES[4]),
    Product("Cushion", 12.99, CATEGORIES[4])
)