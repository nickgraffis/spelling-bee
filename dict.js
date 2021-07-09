const dic = require("./dictionary.json")
const fs = require("fs")

function words() {
  const words = Object.keys(dic)
  fs.writeFileSync("./dict.txt", words.join("\n"))
}

words()