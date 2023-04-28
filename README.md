# webhook Usage

Usage :
// create a webhook for /api namespaces
 apiRouter.Post("/polka/webhooks", handlers.WebhookHandler(db, apiCfg))

 OR do a postman or curl POST query, NOTE ensure webserver is running at 8080 at /api/polka/webhooks
 curl POST http://localhost:8080/api/polka/webhooks

 Response:
 {
	is_chirpy_red: true
 }

When you receive a request to your webhook handler, you should expect an API key in this header format:
Authorization: ApiKey <key>

# what is a webhook
A webhook is just an event that’s sent to your server by an external service. There are just a couple of things to keep in mind when building a webhook handler:

The third-party system will probably retry requests multiple times, so your handler should be idempotent
Be extra careful to never “acknowledge” a webhook request unless you processed it successfully. By sending a 2XX code, you're telling the third-party system that you processed the request successfully, and they'll stop retrying it.
When you’re writing a server, you typically get to define the API. However, when you’re integrating a webhook from a service like PayPal, you’ll probably need to adhere to their API: they’ll tell you what shape the events will be sent in.
Webhooks are transmitted via HTTP or HTTPS, usually as a POST request over a specific URL. The POST data is interpreted by the receiving application’s API, which triggers the requested action and sends a message to the original application to confirm the task is complete. The data sent is commonly formatted using JSON or XML.

# Are webhooks and websockets the same thing?
Nope! A websocket is a persistent connection between a client and a server. Websockets are typically used for real-time communication, like chat apps. Webhooks are a one-way communication from a third-party service to your server.