# for i in {1..5}; do python server.py server-$i 500$i &; done
import sys
from flask import Flask

app = Flask(sys.argv[1])

@app.route('/')
def index():
    return app.name

app.run(port=sys.argv[2])
