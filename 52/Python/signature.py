import hmac
import base64
import json
import time
from hashlib import sha1


def generate_signature(data: str) -> str:
    mac = hmac.new(bytes.fromhex("EAB4F1B9E3340CD1631EDE3B587CC3EBEDF1AFA9"), data.encode(), sha1)
    return base64.b64encode(b"\x52" + mac.digest()).decode("utf-8")


payload = json.dumps({
    "timestamp": int(time.time() * 1000)
})

print(generate_signature(payload))