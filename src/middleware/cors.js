module.exports = function allowCORS() {
    return (req, res, next) => {
        res.header("Access-Control-Allow-Origin", "*");
        res.header("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, HEAD, OPTIONS");
        res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, content-type, Accept");
        if (req.method.toLowerCase() === "options") {
            res.status(200);
            res.send("Ok");
        }
        else {
            next();
        }
    }
}