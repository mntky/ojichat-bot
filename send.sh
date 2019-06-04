#!/bin/bash
# httpは各自のやつ 
curl https://hooks.slack.com/services/T63EC9465/BJ9MPN3HR/T4Vkw3vHKTpaVqTc7y4 -X POST -H "Content-Type: application/json" -d '{"username":"おぢさん", "text":"'$(ojichat $1)'"}'
