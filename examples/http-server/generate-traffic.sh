#!/bin/bash

echo "Generating traffic for the sample app. Use Ctrl-C to exit."

while true; do curl http://localhost:8082; curl localhost:8082/internal-err; curl localhost:8082/err; sleep 2; done