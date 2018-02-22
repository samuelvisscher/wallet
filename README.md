# Initial Kitty Offering Node

Kitties for sale.

## Run (test mode)

Run a node with test data. 100 kitties.

```bash
iko \
    -master-public-key 03429869e7e018840dbf5f94369fa6f2ee4b380745a722a84171757a25ac1bb753 \
    -memory \
    -test-mode \
    -test-secret-key 190030fed87872ff67015974d4c1432910724d0c0d4bfbd29d3b593dba936155 \
    -test-injection-count 100
```

RESTful API will be served on port `:8080`.

**Get Kitty of ID:**

Request:

```text
GET http://127.0.0.1:8080/api/kitty/9
```

Response:

```json
{
    "data": {
        "kitty_id": 9,
        "address": "2fzr9thfdgHCWe8Hp9btr3nNEVTaAmkDk7",
        "transactions": [
            "TO_BE_IMPLEMENTED"
        ]
    }
}
```

**Get Transaction of Hash:**

Request:

```text
GET http://127.0.0.1:8080/api/tx/b4380dc5125320efb24abebda81e0522cec35bca95cf1e9b4e8b4d6a4fda1634
```

Response:

```json
{
    "data": {
        "meta": {
            "hash": "b4380dc5125320efb24abebda81e0522cec35bca95cf1e9b4e8b4d6a4fda1634",
            "raw": "690f7c314facfc570bdae9888c91527a50f6d21129a23498d29bce8f51e9309008000000000000006e2d890c4b9d15150800000000000000000427fcd0f0b9461c5c516cd66a4b5ac413978272000427fcd0f0b9461c5c516cd66a4b5ac413978272832fc2618823f9d2d4427ce439ca0cccce05796a1ca4b5dfe0f1636bd0e3be5844295390909abd6dd1b6056ad36281bff2e819874455818fb8902244fa21637101"
        },
        "transaction": {
            "prev_hash": "690f7c314facfc570bdae9888c91527a50f6d21129a23498d29bce8f51e93090",
            "seq": 8,
            "time": 1519293394965835118,
            "kitty_id": 8,
            "from": "2fzr9thfdgHCWe8Hp9btr3nNEVTaAmkDk7",
            "to": "2fzr9thfdgHCWe8Hp9btr3nNEVTaAmkDk7",
            "sig": "832fc2618823f9d2d4427ce439ca0cccce05796a1ca4b5dfe0f1636bd0e3be5844295390909abd6dd1b6056ad36281bff2e819874455818fb8902244fa21637101"
        }
    }
}
```

**Get Transaction of Sequence:**

Request:

```text
GET http://127.0.0.1:8080/api/tx_seq/7
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
GET http://127.0.0.1:8080/api/address/2fzr9thfdgHCWe8Hp9btr3nNEVTaAmkDk7
```

Response:

```json
{
    "data": {
        "address": "2fzr9thfdgHCWe8Hp9btr3nNEVTaAmkDk7",
        "kitties": [
            7,
            18,
            29,
            42,
            2,
            3,
            35,
            48,
            76,
            91,
            0,
            31,
            72,
            83,
            93,
            94,
            28,
            44,
            24,
            27,
            80,
            86,
            92,
            98,
            21,
            23,
            15,
            41,
            51,
            57,
            63,
            66,
            4,
            6,
            81,
            68,
            73,
            58,
            97,
            1,
            52,
            61,
            70,
            71,
            79,
            90,
            10,
            46,
            38,
            40,
            53,
            62,
            78,
            82,
            11,
            19,
            85,
            88,
            39,
            54,
            59,
            67,
            74,
            77,
            13,
            36,
            95,
            55,
            87,
            8,
            34,
            60,
            65,
            89,
            33,
            50,
            25,
            69,
            75,
            84,
            17,
            20,
            43,
            47,
            56,
            9,
            37,
            26,
            49,
            12,
            16,
            64,
            30,
            32,
            22,
            45,
            96,
            99,
            5,
            14
        ],
        "transactions": [
            "TO_BE_IMPLEMENTED"
        ]
    }
}
```