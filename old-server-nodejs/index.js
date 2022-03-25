const express = require('express')
const path = require('path')
const fs = require('fs')
const app = express()

app.use(express.json({ limit: '5gb' }))

app.post('/', (req, res) => {
  console.log(req.body.hostname)
  fs.writeFile(
    path.join(__dirname, 'data', req.body.hostname + '.json'),
    JSON.stringify(req.body),
    err => {
      if (err) console.error(err)
    }
  )
  res.send()
})

app.listen(3000, () => console.log('listening on port 3000'))