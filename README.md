# webhook
/*
Usage :
// create a webhook for /api namespaces
 apiRouter.Post("/polka/webhooks", handlers.WebhookHandler(db, apiCfg))

 OR do a postman or curl POST query, NOTE ensure webserver is running at 8080 at /api/polka/webhooks
 curl POST http://localhost:8080/api/polka/webhooks

 Response:
 {
	is_chirpy_red: true
 }
*/