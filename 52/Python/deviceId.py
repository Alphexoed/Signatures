import os
import hmac
from hashlib import sha1


def generate_device_id() -> str:
    identifier = os.urandom(20)
    key = bytes.fromhex("AE49550458D8E7C51D566916B04888BFB8B3CA7D")
    mac = hmac.new(key, b"\x52" + identifier, sha1)
    return f"52{identifier.hex()}{mac.hexdigest()}"


print(generate_device_id())