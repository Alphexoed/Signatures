import hmac
import base64
import json
import time
from hashlib import sha1


def generate_signature(data: str) -> str:
    mac = hmac.new(bytes.fromhex("DFA5ED192DDA6E88A12FE12130DC6206B1251E44"), data.encode(), sha1)
    return base64.b64encode(b"\x19" + mac.digest()).decode("utf-8")


payload = json.dumps({
    "timestamp": int(time.time() * 1000)
})

print(generate_signature(payload))