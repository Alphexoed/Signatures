using System.Security.Cryptography;
using System.Text;

static T[] CombineTwoArrays<T>(T[] a1, T[] a2)
{
    T[] arrayCombined = new T[a1.Length + a2.Length];
    Array.Copy(a1, 0, arrayCombined, 0, a1.Length);
    Array.Copy(a2, 0, arrayCombined, a1.Length, a2.Length);
    return arrayCombined;
}


static byte[] StringToByteArray(string hex)
{
    return Enumerable.Range(0, hex.Length)
                     .Where(x => x % 2 == 0)
                     .Select(x => Convert.ToByte(hex.Substring(x, 2), 16))
                     .ToArray();
}

static string signature(string data)
{
    string prefix = "52";
    string key = "EAB4F1B9E3340CD1631EDE3B587CC3EBEDF1AFA9";

    HMACSHA1 hmac = new HMACSHA1(StringToByteArray(key));
    byte[] buffer = Encoding.Default.GetBytes(data);
    byte[] result = hmac.ComputeHash(buffer);
    return Convert.ToBase64String(CombineTwoArrays(StringToByteArray(prefix), result));
}

string sig = signature("{}");
Console.WriteLine(sig);