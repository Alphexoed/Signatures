package Generators;

import org.apache.commons.codec.digest.HmacUtils;
import java.util.Arrays;
import java.util.Base64;

public class NdcMsgSigGen {
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

    public static void main(String[] args) {
        String data = "{}";

        byte[] hmac = HmacUtils.hmacSha1(hexStringToByteArray("307c3c8cd389e69dc298d951341f88419a8377f4"), data.getBytes());
        String finall = Base64.getEncoder().encodeToString(concat(hexStringToByteArray("22"), hmac));

        System.out.println(finall);
    }
}

