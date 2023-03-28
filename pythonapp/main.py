#!/usr/bin/env python
# encoding: utf-8
import os
from flask import Flask, jsonify
from pymongo import MongoClient

app = Flask(__name__)

client = MongoClient(os.environ.get('CONNECTION_URI'))
db = client[os.environ.get('DB_NAME')]
collection = db[os.environ.get('COLLECTION_NAME')]

@app.route('/', methods=['GET'])
def index():
    return "Merhaba Python!"

@app.route('/staj', methods=['GET'])
def getCity():
    document = collection.aggregate([{ '$sample': { 'size': 1 } }]).next()
    document.pop('_id')
    return jsonify(document)

@app.errorhandler(404)
def page_not_found(e):
    return jsonify({'message': 'Not Found'}), 404

app.run(host="0.0.0.0", port=4444)