from flask import Flask, request
import json
import os

if not os.path.exists('data'):
    os.mkdir('data')

app = Flask(__name__)

@app.post('/')
def index():
    print(request.json['hostname'])
    with open(os.path.join(
        'data',
        request.json['hostname'] + '.json'
    ), 'w') as f:
        f.write(json.dumps(request.json))
        f.close()

    return ''