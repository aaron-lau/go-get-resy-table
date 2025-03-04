#!/bin/bash

# test.sh
curl -X POST http://localhost:8080/book \
  -H "Content-Type: application/json" \
  -H "X-Resy-Auth-Token: $RESY_AUTH_KEY" \
  -d '{
    "restaurant_name": "Test Restaurant",
    "date": "2024-02-14",
    "time": "19:00",
    "party_size": 2
  }'