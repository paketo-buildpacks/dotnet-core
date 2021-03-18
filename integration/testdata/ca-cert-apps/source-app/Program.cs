using System.Security.Cryptography.X509Certificates;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Hosting;
using Microsoft.AspNetCore.Server.Kestrel.Https;

namespace ScottBrady.Pem.Kestrel
{
    public class Program
    {
        public static void Main(string[] args)
        {
            CreateHostBuilder(args).Build().Run();
        }

        public static IHostBuilder CreateHostBuilder(string[] args) =>
            Host.CreateDefaultBuilder(args)
                .ConfigureWebHostDefaults(webBuilder =>
                {
                    const string certPem = @"-----BEGIN CERTIFICATE-----
MIIFfTCCA2WgAwIBAgICBnowDQYJKoZIhvcNAQELBQAwTjEaMBgGA1UEChMRUGFr
ZXRvIEJ1aWxkcGFja3MxMDAuBgNVBAMTJ1Bha2V0byBCdWlsZHBhY2tzIENlcnRp
ZmljYXRlIEF1dGhvcml0eTAeFw0yMTAzMDIxNjU5MjdaFw0zMTAzMDIxNjU5Mjda
MEQxGjAYBgNVBAoTEVBha2V0byBCdWlsZHBhY2tzMSYwJAYDVQQDEx1QYWtldG8g
QnVpbGRwYWNrcyBDZXJ0aWZpY2F0ZTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCC
AgoCggIBAKDexTQRJcn/3xPtCWUhJv20s1pjbr1QxaJlOv6sI+xXBg0nMSn2i0d4
G7BC9yt704WJEy3rnKpA7WPejGYqi5dW711Z/gkie8MNL3zR1FrOX5/MgRgGm2S0
3HNdjxs84zeanhRc64Z80Qyu7dfi7i68mUzp8Z5rwU9Lf8HuouAG9H94Uv4v/x4A
gIZydiSLepG4Mp1QTznogVUMsSynwG6xWLPTsfOTA9fMidTSKSFcHeLl/qBnecWv
MxQdIm8hkSlq+kGu+GXuZQA3eRBufvEhX/swZGIwla8XjDQqIWubQlLvNHm9ETMm
Eru3MTX1sO71Xiv2rZhaZeg/hnUFDZrR7cvFZLbGdhkApWpmq/5lAzKKu+QRgbXN
+TfMfw0Y650b4Nq718NEmwFjYp1TXWrcBB+uG4qvrtYbBC9W/BN/EuNeBUIBZIA0
66UpoKCRXbd3OKZV0OPZs4apqqi+y+P54ye5tFEoASdAN6g70P4gIklUz6ftCj8T
fssPXRysqIwtZB0LTdEU0GZtfhGd32xF5RswJwxayfTOvd9c5yrqYs7iGg21MoOL
uxp+LU3Nb+Q96NMW+JuJ6XNSjTzpBTTUmFKeggAgizh7r+gHiX/sJT/g50cRztc6
euJLEl1u8cSqiX/p5a6Ecpieu7ORjJZqJ6LPUEv5KDDWeNvGhoGBAgMBAAGjbzBt
MA4GA1UdDwEB/wQEAwIHgDAdBgNVHSUEFjAUBggrBgEFBQcDAgYIKwYBBQUHAwEw
DgYDVR0OBAcEBQECAwQGMCwGA1UdEQQlMCOCCWxvY2FsaG9zdIcEfwAAAYcQAAAA
AAAAAAAAAAAAAAAAATANBgkqhkiG9w0BAQsFAAOCAgEAiexbcqQJxIjKFSBQSLpX
q21pzxoCr8qULMoJEBTiBz+Aqn23nE076PBiKAd+8jWP2wIiQXqhdDqE0QV/MHvE
v3Atabd7L5hkulUw13Gl7JtG750b2nC9EoJpimgicmh/upWtzBUcBNXUN0cY9N+I
T58CfG/MYLbGR07mSeRlxnQiTMSfJ5m1G57UmqOSBqhqXai6x5/MCQHl+MIglLqh
+bRzeS1ZlgFKFWaoa+lHVQt5qqDsaaQmtFdjskudEjezlsKXNmHPOFbOH9enOJ7Z
1EobCIvb9eLCjc8FL6hEGyVFACuxAkDds23idwIJXcOVPMs7JOPa/Lb5jBNbOYLc
hYj1Op0piHEPosedtighQc3sFNvroDRYm46zBRrwbLTPQlplOBkZEssUnoDCJQJq
TwgwucnRj0nnZtHfSsaHv5LcWjW+GI9ox2qpxx8XJgTnmr2LfiNJIBViGppqjEoZ
FgGzqcd3Gle0LAXcz8mzCxHsGb0AyyVK7+GziZxTc7tiBorpn9LOA0tBlrWTrN0u
ptIfJIsKo//HVwv/rhEfRapQYTMv4trF6jFF/Kh/eyfbfkndRFlu3ywdtf+3peKI
vYZ8GeSBM86UvRyMA9nWbj8Q7NEsmg9ZHlzzG7iC32U+Zp1yy/rDrKqQWKxNojrv
o7kYPHxnraAtNlmQEDyEQ9s=
-----END CERTIFICATE-----";
                    const string keyPem = @"-----BEGIN RSA PRIVATE KEY-----
MIIJKAIBAAKCAgEAoN7FNBElyf/fE+0JZSEm/bSzWmNuvVDFomU6/qwj7FcGDScx
KfaLR3gbsEL3K3vThYkTLeucqkDtY96MZiqLl1bvXVn+CSJ7ww0vfNHUWs5fn8yB
GAabZLTcc12PGzzjN5qeFFzrhnzRDK7t1+LuLryZTOnxnmvBT0t/we6i4Ab0f3hS
/i//HgCAhnJ2JIt6kbgynVBPOeiBVQyxLKfAbrFYs9Ox85MD18yJ1NIpIVwd4uX+
oGd5xa8zFB0ibyGRKWr6Qa74Ze5lADd5EG5+8SFf+zBkYjCVrxeMNCoha5tCUu80
eb0RMyYSu7cxNfWw7vVeK/atmFpl6D+GdQUNmtHty8VktsZ2GQClamar/mUDMoq7
5BGBtc35N8x/DRjrnRvg2rvXw0SbAWNinVNdatwEH64biq+u1hsEL1b8E38S414F
QgFkgDTrpSmgoJFdt3c4plXQ49mzhqmqqL7L4/njJ7m0USgBJ0A3qDvQ/iAiSVTP
p+0KPxN+yw9dHKyojC1kHQtN0RTQZm1+EZ3fbEXlGzAnDFrJ9M6931znKupizuIa
DbUyg4u7Gn4tTc1v5D3o0xb4m4npc1KNPOkFNNSYUp6CACCLOHuv6AeJf+wlP+Dn
RxHO1zp64ksSXW7xxKqJf+nlroRymJ67s5GMlmonos9QS/koMNZ428aGgYECAwEA
AQKCAgAUUxDnOyNjGgi9I72EIWQjuajPSrC7CnFtywxhEK6ZNYV2M/VqL9P4+5vD
8TH5NHPM8zyRGKt6dymG7J8gaU+plzo2uR/3V3v7cLcHNht2PYynS9cjifIoDxGr
Ia7q6g5rAAXo3LSFEU/4IkG6fNlK3lkf9o6oTUTnF8rUXaoGU9qgIDucEcRRrg6O
7fcvNtANiRAcAAGCd3WfoTLhSXBui8mBLsXU2EYsBZOEZ+j6ZgEAob5B1dD0wOXb
LLMlB0Cn0vQ7SDfp6Oyp0lhhUxSGsojF259TKIBA1uDH1mrShZMjl0Ux/EkoBS9o
uARnpNrt1eJH+6qDDSjC5wO91R2atpLe2I+igKQTzrTN/6V0JEsPvJu0Vg4hhT6o
CNuGB4q3HQ5nKsOCIcTUEsYXyGFTRzLU0IpfJwB5RwVuQ/wPjXHYSYFjSxjyF0L2
Y7L4Q0c1b+O21oRIYsAhnT+9Zn3JuzjXbrDoc5UwuvxYOie0ZaxRzquJhmatPiSI
RADFofCu9ORwkZT0jfB3GGmKVydgiEJJfkBGkxzZjyTcUBgFlXHVy7BN9dThaFot
7G7ZB8ukarMD6vxebZtRfXeMqWgEKw9Bd9Kkn2r0JrufGaVGZ2wSeS0SX2GLxBn+
QhBuHMzt4BTP8WdR0xMBz7sEjihoWvGDjKzWNdmjOa0ptlvvYQKCAQEAzVqSuE2S
yUl19dXCfJlCS8H6/Yv3RkzWF70ZnNxxUkWy2UvRFjSSt8q7aneBTIB7kvI1qY5C
I6KBAi/B/vmzctokmE9MDMb2UsapN9zaXuD7nFYxZ7ejK1600CniH1tU0e3w9BSQ
VeJTUaqHL4X2R4MEo7n49bLO2iJ2FG98GYmfVSPZJH35obeZU6jZko3xsb9Rez7O
iQ4VHBSKq5Zp0Gluo0sENZvO9dj1IHshf4UDaT/j6PTlzBVKLpb49cHqo6Q3ZfJi
E8VLTjak0/vQtxTHJeZ5rNyjEH9rwJCbWCVKQGNweqXtcXdLapbNYH2s7zJgGgjV
1iepE0djA6ZIBwKCAQEAyIuij3y0LnCEx8Mkr1idnxeJ7UwXxuzpAatCUCypCaew
JDxbupQynOxKMIDrmEQ2AmNpyjQ/2PvAdZ+UYVYdd0cpRgeu1FRpeZwHXRHP5dPG
xBfsFJZxMKpD9Ks8qPzLuVDlaRtzgsvVqGV0cvFKNxXC/dfXeHLH4b8SLlfvJVpC
gBeIOrzak1f3H0E5AxpUvxkeh1dliWdiqoxI6hIAQJPdejrT4fGA+fJGzNpSMJYo
HeQXjvbN64iAbaNW91KzN4akB7UTeV6skaeVKApX/1B5jr4oR2Qhr53k+lixmkqK
Vkc0GLOog1+SYAWZ+UCsh8MYtz72bLcJRIqoY7i4NwKCAQASNuduq4rKidaJsKUo
kht8Rr9xf9m2BZiz0FUWQcNXbdE4Tu5DzvP4k2XeQq1YUMklNpCl8nVRXdedjwjh
Cdrt5AV88QOo2nj8zJwz1UYVRlVq/4YwUeyKK0NUd3KUH3C9kiJeM/i9dW64fqyw
/Wvj2e4ua492k56fYJEirOTQCxHz5lMbVoUu3+8cqYxq4GZAwtyCVwbQx0v+CqBm
KdhG4SMsHrpH2wMXcWAEuoc1HWI5Eq6vehFr8bN7wG5itgmO7EDxrPcgE87jKBuk
peBUbOZhKTk/qO5Zx0OSeAEfZ2dXoLpYVqFiABfTY37iASO2r7wwcvosnaX0jM9u
gnA/AoIBAQC014Yd0hxBDGIQKV75Z/WrMvTDsax3S8fKI39HAR7lf/uMkYk+NMl3
THSSTI8m3cu+V5tcJcW1iz/AUcjiBV/I4bjMV71F52C9sv/I43kQDOrehZvz7L3h
XoitJ4Up9dxYHiThpUmClwDyO5rI0+FSzyLo+Sxqh0bLwRtKAy26ByyUiaPlI8wO
tnI4Ev6wV5w4PxSSgzMitsH4fUx7FwR3N1+vC0FqK/dcbSd/LxiSi7VdTwQXfWOv
k4YMWBDiMgc+eQGNmbIX7lG7ft04ICu+JfmXyM5VomvmC4IiZryxH6qjps3JwKii
3xoF0MdKRxHN0xaEmBhrbJrE3ix+0GH9AoIBAGOeMzXrwNxU81SXg8F8F/jCrplo
hZgDpv2rDwij5hCo6DXI54D6cNLf6utgvtd5khqWRujJA4i4u+MjyKTSpfD/431a
TvBUSa0rg51TSahUcjaW4um2YCfBogZKDmj9mylunA+hGiqukQLlZQE0lz9gpMH5
XUMdZIaKcGqHiE9nGGW0g/r5b1iF770lFLRN5Acc0XCAyFM8//Rg2qmW5z4fAb2q
8iXk+hllfsx9FI2jHa6s7OqqPzJ+w7o6CMYXKQTGKq8obXV1k95vQWIf67krkdzA
kupjszBrySfMWZdJYqwW0jiTTMyItD4L47nFWc/o7PEIvKKHO2OMUeztklg=
-----END RSA PRIVATE KEY-----";
                    
                    webBuilder.ConfigureKestrel(options =>
                    {
                        options.ConfigureHttpsDefaults(adapterOptions =>
                        {
                            adapterOptions.ServerCertificate = X509Certificate2.CreateFromPem(certPem, keyPem);
                            adapterOptions.ClientCertificateMode = ClientCertificateMode.RequireCertificate;
                        });
                    });
                    
                    webBuilder.UseStartup<Startup>();
                });
    }
}
