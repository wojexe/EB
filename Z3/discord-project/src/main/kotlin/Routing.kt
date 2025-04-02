package com.example

import io.ktor.http.*
import io.ktor.server.application.*
import io.ktor.server.request.*
import io.ktor.server.response.*
import io.ktor.server.routing.*

fun Application.configureRouting(discordBot: DiscordBot) {
    routing {
        get("/") {
            call.respondText("Hello World!")
        }
        
        post("/discord/send") {
            try {
                val messageRequest = call.receive<MessageRequest>()
                discordBot.sendMessage(messageRequest.content)
                call.respond(HttpStatusCode.OK, "Message sent successfully")
            } catch (e: Exception) {
                call.respond(HttpStatusCode.InternalServerError, "Failed to send message: ${e.message}")
            }
        }
    }
}
