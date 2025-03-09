package com.example

import java.sql.DriverManager

fun main() {
    println("Hello from Kotlin!")

    // Przykład użycia JDBC SQLite
    try {
        Class.forName("org.sqlite.JDBC")
        val connection = DriverManager.getConnection("jdbc:sqlite::memory:")
        connection.use { conn ->
            println("SQLite connection established successfully!")

            val statement = conn.createStatement()
            statement.execute("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)")
            statement.execute("INSERT INTO users (name) VALUES ('Sample User')")

            val resultSet = statement.executeQuery("SELECT * FROM users")
            while (resultSet.next()) {
                println("User ID: ${resultSet.getInt("id")}, Name: ${resultSet.getString("name")}")
            }
        }
    } catch (e: Exception) {
        println("SQLite connection error: ${e.message}")
    }

    println("Application completed successfully!")
}
