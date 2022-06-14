#!/bin/bash

export $(cat .env)

export TIME_ZONE=Asia/Taipei

export HTTP_LISTEN_ADDR=0.0.0.0
export HTTP_LISTEN_PORT=8080
export ALLOW_ORIGINS=http://localhost:3000,https://localhost:3000

export BLOCKCHAIN_CURRENCY=bitcoin,ethereum,solana,green-satoshi-token,stepn
export BLOCKCHAIN_API_URL=https://api.coingecko.com/api/v3/coins

export LINEBOT_CHANNEL_SECRET=
export LINEBOT_CHANNEL_ACCESS_TOKEN=