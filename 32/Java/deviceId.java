package Generators;

import org.apache.commons.codec.digest.HmacUtils;
import java.math.BigInteger;
import java.security.NoSuchAlgorithmException;
import java.security.SecureRandom;
import java.util.Arrays;

public class DeviceIdGen {
    private static byte[] hexStringToByteArray(String s) {
        int len = s.length();
        byte[] data = new byte[len / 2];
        for (int i = 0; i < len; i += 2) {
            data[i / 2] = (byte) ((Character.digit(s.charAt(i), 16) << 4) + Character.digit(s.charAt(i+1), 16));
        }
        return data;
    }

    private static byte[] concat(byte[] a, byte[] b) {
        int lenA = a.length;
        int lenB = b.length;
        byte[] c = Arrays.copyOf(a, lenA + lenB);
        System.arraycopy(b, 0, c, lenA, lenB);
        return c;
    }

    public static void main(String[] args) throws NoSuchAlgorithmException {
        byte[] identifier = new byte[20];
        SecureRandom.getInstanceStrong().nextBytes(identifier);

        byte[] part = concat(hexStringToByteArray("32"), identifier);

        String hmac = HmacUtils.hmacSha1Hex(hexStringToByteArray("76b4a156aaccade137b8b1e77b435a81971fbd3e"), part);
        String identifierHex = new BigInteger(1, identifier).toString(16);
        String finall = "32" + identifierHex + hmac;

        System.out.println(finall.toUpperCase());
    }
}
