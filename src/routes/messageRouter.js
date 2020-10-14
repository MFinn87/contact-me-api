module.exports = function(router, emailService) {
    function sendResponse(res, status, message) {
        res.status(status)
        res.send(message)
    }
    
    router.post("/message", function(req, res) {
        try {
            const name = req.body.name
            const phone = req.body.phone
            const senderEmail = req.body.email
            const message = req.body.message

            if(!name) {
                return sendResponse(res, 400, "Name must be included in the body.")
            }

            if(!senderEmail) {
                return sendResponse(res, 400, "Email address must be included in the body.")
            }
            
            if(!message) {
                return sendResponse(res, 400, "Message must be included in the body.")
            }
            
            return emailService(name, phone, senderEmail, message)
            .then((resp) => {
                sendResponse(res, 200, 'OK');
                console.log(resp)
            })
            .catch((error) => {
                sendResponse(res, 400, error.message);
            })
        }
        catch (error) {
            sendResponse(res, 400, error.message);
        }
       
    })

    return router
}