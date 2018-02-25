# Kittycash Wallet

Where one claims ownership of dem' kitties.

## Run IKO node (test mode)

Run a node with test data. 100 kitties.

```bash
iko \
    -master-public-key 03429869e7e018840dbf5f94369fa6f2ee4b380745a722a84171757a25ac1bb753 \
    -memory \
    -test-mode \
    -test-secret-key 190030fed87872ff67015974d4c1432910724d0c0d4bfbd29d3b593dba936155 \
    -test-injection-count 10
```

RESTful API will be served on port `:8080`.

**Get Kitty of ID:**

Request:

```text
GET http://127.0.0.1:8080/api/iko/kitty/9
```

Response:

```json
{
    "data": {
        "kitty_id": 9,
        "address": "2fzr9thfdgHCWe8Hp9btr3nNEVTaAmkDk7",
        "transactions": [
            "bf04e7d9a39a94acb2b60810d77ebafea566b3ee065b109238e2b9765673c40b"
        ]
    }
}
```

**Get Transaction of Hash:**

Request:

```text
GET http://127.0.0.1:8080/api/iko/tx/f6f70fbd908a2cdbbd948f2e30317970f76c1ef32ccded970062290cbb455190
```

Response:

```json
{
    "data": {
        "meta": {
            "hash": "f6f70fbd908a2cdbbd948f2e30317970f76c1ef32ccded970062290cbb455190",
            "raw": "a6ebfd0fcf5f4772c446b2c4021928ba314d0a486a62480b3cd461b8b16c13440800000000000000c0bc46a6927516150800000000000000000427fcd0f0b9461c5c516cd66a4b5ac413978272000427fcd0f0b9461c5c516cd66a4b5ac4139782729d476717ae39c118e13fc2e44f9225a6bbbbcb6780d4460348b634fa89a21ebc5e23eb56246e33368df802aa2b0e89fdc4a7752a51307c04225c80952a9b087801"
        },
        "transaction": {
            "prev_hash": "a6ebfd0fcf5f4772c446b2c4021928ba314d0a486a62480b3cd461b8b16c1344",
            "seq": 8,
            "time": 1519531196999449792,
            "kitty_id": 8,
            "from": "2fzr9thfdgHCWe8Hp9btr3nNEVTaAmkDk7",
            "to": "2fzr9thfdgHCWe8Hp9btr3nNEVTaAmkDk7",
            "sig": "9d476717ae39c118e13fc2e44f9225a6bbbbcb6780d4460348b634fa89a21ebc5e23eb56246e33368df802aa2b0e89fdc4a7752a51307c04225c80952a9b087801"
        }
    }
}
```

**Get Transaction of Sequence:**

Request:

```text
GET http://127.0.0.1:8080/api/iko/tx_seq/7
```

Response:

```json
{
    "data": {
        "meta": {
            "hash": "690f7c314facfc570bdae9888c91527a50f6d21129a23498d29bce8f51e93090",
            "raw": "87901ffad3f378bdd75a5ba9e435c66468414ef4bcf76e4134ac208650283c8d07000000000000005c95640c4b9d15150700000000000000000427fcd0f0b9461c5c516cd66a4b5ac413978272000427fcd0f0b9461c5c516cd66a4b5ac4139782725ac63892ee8e25f0a6c27e6f065ae755c621d51bb0ac288f537d7bc362dd09a13e42a93b2fb48060c9ee01e4f187177bc8654e2ea10b970d84e10c40a4d62d6f01"
        },
        "transaction": {
            "prev_hash": "87901ffad3f378bdd75a5ba9e435c66468414ef4bcf76e4134ac208650283c8d",
            "seq": 7,
            "time": 1519293394963436892,
            "kitty_id": 7,
            "from": "2fzr9thfdgHCWe8Hp9btr3nNEVTaAmkDk7",
            "to": "2fzr9thfdgHCWe8Hp9btr3nNEVTaAmkDk7",
            "sig": "5ac63892ee8e25f0a6c27e6f065ae755c621d51bb0ac288f537d7bc362dd09a13e42a93b2fb48060c9ee01e4f187177bc8654e2ea10b970d84e10c40a4d62d6f01"
        }
    }
}
```

**Get Address:**

Request:

```text
GET http://127.0.0.1:8080/api/iko/address/2fzr9thfdgHCWe8Hp9btr3nNEVTaAmkDk7
```

Response:

```json
{
    "data": {
        "address": "2fzr9thfdgHCWe8Hp9btr3nNEVTaAmkDk7",
        "kitties": [
            0,
            1,
            2,
            3,
            4,
            5,
            6,
            7,
            8,
            9
        ],
        "transactions": [
            "9d30e6cd189dc3faf4afde0d48fa0d90d062f9cef211ccf9ce2e4fd01520cf18",
            "ab4f013da9fdc5a890330fb1ae89f2d884ab3cee3af21510865a4b8afd923bff",
            "de557e38f393545843e6a7213fdbd86e89d52b4ab45589fdd7120f993751bcee",
            "f521f44e0e959ef3c799da9d05ab3bd15a930b2354f0fd7a4267ebf243a6479f",
            "a020ba08fe38205e69c8bfb78dc9e279dc6476b72109d85dc291903241af8793",
            "13e8eb5d4d5448e345dc8e64a6ec971231b6617406ee16af65e08dcae8da7531",
            "1d88a6c15e40b40d846fe00fae0bd893fd3fd1e42af023349190c262d364b871",
            "a6ebfd0fcf5f4772c446b2c4021928ba314d0a486a62480b3cd461b8b16c1344",
            "f6f70fbd908a2cdbbd948f2e30317970f76c1ef32ccded970062290cbb455190",
            "bf04e7d9a39a94acb2b60810d77ebafea566b3ee065b109238e2b9765673c40b"
        ]
    }
}
```