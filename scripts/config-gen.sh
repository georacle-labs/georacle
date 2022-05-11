#!/bin/bash

if [[ -z "${NETWORK}" ]]; then
	NETWORK=$1
fi

if [[ -z "${WS_URI}" ]]; then
	WS_URI=$2
fi

if [[ -z "${DB_URI}" ]]; then
	DB_URI=$3
fi

if [[ -z "${PASS}" ]]; then
	PASS=$4
fi

if [[ -z "${ADDR}" ]]; then
	PORT=$5
fi

if [[ -z "${PORT}" ]]; then
	PORT=$6
fi

echo -e "{\n
    \"network\":   \"${NETWORK}\",\n
    \"ws_uri\":    \"${WS_URI}\",\n
    \"db_uri\" :   \"${DB_URI}\",\n
    \"password\":  \"${PASS}\",\n
    \"addr\":      \"${ADDR}\",\n
    \"port\":      ${PORT}\n
}"
