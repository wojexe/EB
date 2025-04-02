package models

// Input form for creating/updating products (no ID required from client)
case class ProductFormInput(name: String, description: String, price: Double)
