const express = require("express");
const app = express();
const port = process.env.PORT || 9000;
const bodyParser = require("body-parser");
const notifier = require("node-notifier");
const path = require("path");

app.use(bodyParser.json());

app.get("/health", (req, res) => res.status(200).send());
app.post("/notify", (req, res) => {
    notify(req.body, reply => res.send(reply));
})

app.listen(port, () => {
    console.log(`Server is running on ${port}`)
});


const notify = ({ title, message }, cb) => {
    notifier.notify({
        title: title || "Unknown title",
        icon: path.join(__dirname, "chris-icon.jpg"),
        message: message || "Unknown message",
        sound: true,
        wait: true,
        reply: true,
        closeLabel: "Completed",
        timeout: 15,
    },
        (err, resp, reply) => {
            cb(reply);
        }
    )
}