const dotenv = require('dotenv')
const path = require('path')
const express = require('express')
const app = express()
const bodyParser = require('body-parser')
dotenv.config({"path": path.resolve(__dirname, `../config/${process.env.ENVIRONMENT}.env`)})

const emailService = require('./services/emailService')
const messageRoutes = require('./routes/messageRouter')
const corsHandler = require('./middleware/cors')

app.use(bodyParser.json(), corsHandler())

const messageRouter = messageRoutes(express.Router(), emailService)

app.use("/api", messageRouter);

if(!process.env.port) {
    throw(Error('Port environment variable is required.'))
}
app.listen(process.env.port, () => console.log(`Example app listening on port ${process.env.port}!`))

