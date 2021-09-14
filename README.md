# Stock Application

---

## Description
They are two APIs that interact together `stock-api` and` encryption-api` to obtain information about a stock and
then encrypt it using AES-256 in the encryption-api.

---

## Start up:

This project uses docker-compose (docker required) to start the two APIs,
you can run the following commands from the project root:

    docker-compose build
    docker-compose up -d

---

### Stock-API

It has two endpoints, a GET that is responsible for looking for the information of the last day for a stock and encrypts it obtaining the payload 
and a token to be able to decrypt (works like an OTP) and a POST to decrypt this payload together with the token.

#### Resources

**GET to encrypted stock**

Endpoint: /stock/:symbol

Example:

    curl -X GET http://localhost:8080/stock/IBM 
    
Response:

    {
        "token": "0e3b0862-4823-40ee-8d47-a259d732e012",
        "payload": "c0de60bbce8fe0f32e7d72d44ea5cda698f46ee235e332f0c9e60a923029c5624c542013c323dd5505a5bfd85826fcf4d474e6bed12d2285acb2ba7a0f987c78621e9c9d7da4614ad4052575014e7ae853ed9f705397c9af6c0ecb526ead11216cc7005efa4870626f89cff7555f93706d65075701fead9ea8d2ed4a40683a75e258804e769afb2a1420c3baf65c63a9"
    }    

**POST to decrypt the stock**

Endpoint: /stock/decrypt/:token

Example:

    curl -X POST \
      http://localhost:8080/stock/decrypt/0e3b0862-4823-40ee-8d47-a259d732e012 \
      -d '{
    	"payload": "c0de60bbce8fe0f32e7d72d44ea5cda698f46ee235e332f0c9e60a923029c5624c542013c323dd5505a5bfd85826fcf4d474e6bed12d2285acb2ba7a0f987c78621e9c9d7da4614ad4052575014e7ae853ed9f705397c9af6c0ecb526ead11216cc7005efa4870626f89cff7555f93706d65075701fead9ea8d2ed4a40683a75e258804e769afb2a1420c3baf65c63a9"
        }'
        
Response:
    
    {
        "date": "2021-03-29T00:00:00Z",
        "open": "135.9800",
        "high": "137.0700",
        "low": "135.5100",
        "close": "135.8600",
        "volume": "4622664"
    }