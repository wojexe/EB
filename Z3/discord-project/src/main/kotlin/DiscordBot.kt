package com.example

import dev.kord.common.entity.Snowflake
import dev.kord.core.Kord
import dev.kord.core.entity.channel.MessageChannel
import dev.kord.core.event.message.MessageCreateEvent
import dev.kord.core.on
import dev.kord.gateway.Intent
import dev.kord.gateway.PrivilegedIntent

class DiscordBot(private val token: String, private val channelId: String) {
    private lateinit var kord: Kord

    suspend fun start() {
        kord = Kord(token)

        kord.on<MessageCreateEvent> {
            if (message.author?.isBot != false) return@on

            onGreet()
            onRequestCategories()
            onRequestProducts()
        }

        kord.login {
            @OptIn(PrivilegedIntent::class)
            intents += Intent.MessageContent
        }
    }

    private suspend fun MessageCreateEvent.onGreet() {
        if (message.content == "!greet") {
            message.channel.createMessage("Hello, ${message.author?.username}! 🎀")
        }
    }

    private suspend fun MessageCreateEvent.onRequestCategories() {
        if (message.content == "!categories") {
            val categories = CATEGORIES.joinToString(", ") { it.name }
            message.channel.createMessage("Available categories: $categories")
        }
    }

    private suspend fun MessageCreateEvent.onRequestProducts() {
        val regex = Regex("^!products(?:\\s+(.+))?$")
        val matchResult = regex.find(message.content.trim()) ?: return

        // no category provided = empty string
        val categoryName = matchResult.groupValues[1]

        if (categoryName.isNotEmpty()) {
            // filter products by specified category
            val category = CATEGORIES.find { it.name.equals(categoryName, ignoreCase = true) }

            if (category != null) {
                val filteredProducts = PRODUCTS.filter { it.category == category }
                if (filteredProducts.isNotEmpty()) {
                    val products = filteredProducts.joinToString("\n") { "- ${it.name} for \$${it.price}" }
                    message.channel.createMessage("Products in ${category.name}:\n$products")
                } else {
                    message.channel.createMessage("No products found in category ${category.name}")
                }
            } else {
                message.channel.createMessage(
                    "Category '$categoryName' not found. Available categories: ${
                        CATEGORIES.joinToString(", ") { it.name }
                    }"
                )
            }
        } else {
            // show all products
            val products = PRODUCTS.joinToString("\n") { "- ${it.name} for \$${it.price}" }
            message.channel.createMessage("All products:\n$products")
        }
    }

    suspend fun sendMessage(content: String) {
        val channel = kord.getChannel(Snowflake(channelId)) as? MessageChannel
            ?: throw IllegalArgumentException("Channel not found or not a message channel")

        channel.createMessage(content)
    }

    suspend fun stop() {
        if (::kord.isInitialized) {
            kord.shutdown()
        }
    }
}