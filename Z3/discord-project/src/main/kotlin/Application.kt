package com.example

import io.ktor.server.application.*
import io.ktor.server.netty.*
import kotlinx.coroutines.DelicateCoroutinesApi
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch
import kotlinx.coroutines.runBlocking

fun main(args: Array<String>): Unit = EngineMain.main(args)

@OptIn(DelicateCoroutinesApi::class)
fun Application.module() {
    val discordToken = System.getenv("DISCORD_TOKEN")
    val discordChannelId = System.getenv("DISCORD_CHANNEL_ID")

    val discordBot = DiscordBot(discordToken, discordChannelId)

    configureSerialization()
    configureRouting(discordBot)

    GlobalScope.launch {
        discordBot.start()
    }

    // Handle application shutdown
    environment.monitor.subscribe(ApplicationStopping) {
        runBlocking {
            discordBot.stop()
        }
    }
}