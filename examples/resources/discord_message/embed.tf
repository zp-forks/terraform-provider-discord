resource "discord_message" "hello_world" {
  channel_id = var.channel_id

  embed {
    title = "Hello World"

    footer {
      text = "I'm awesome"
    }

    fields {
      name   = "foo"
      value  = "bar"
      inline = true
    }

    fields {
      name   = "bar"
      value  = "baz"
      inline = false
    }
  }
}
