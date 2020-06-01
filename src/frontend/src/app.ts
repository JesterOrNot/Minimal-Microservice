import * as express from "express";
import * as path from "path";

const app = express();

app.use(express.static('/static'));

app.get('/', (req, res) => {
    res.sendFile(__dirname + "/static/index.html");
})

app.get('/index.js', (req, res) => {
    res.sendFile(__dirname + "/static/index.js")
})

app.listen(8080);

console.log("Server running on port 3000\nlocalhost:3000");
