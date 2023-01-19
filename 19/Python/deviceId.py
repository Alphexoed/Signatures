import os
import hmac
from hashlib import sha1


def generate_device_id() -> str:
    identifier = os.urandom(20)
    key = bytes.fromhex("E7309ECC0953C6FA60005B2765F99DBBC965C8E9")
    mac = hmac.new(key, b"\x19" + identifier, sha1)
    return f"19{identifier.hex()}{mac.hexdigest()}"


print(generate_device_id())