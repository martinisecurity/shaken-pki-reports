# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 115K

Tested At: 31 Oct 22 18:32 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 9 day(s)\
Subject: CN=SHAKEN 115K, OU=SHAKEN, O=RenterNET, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/115K/fc9eb1ec-7be6-4ebc-b2ac-d1727482970c.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC6zCCApGgAwIBAgIQaP6ZGSSxoGcGoz085IvXSDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTAxODMwMjhaFw0yMjExMDkxODMwMjdaMEgxCzAJBgNVBAYTAlVTMRIwEAYDVQQKEwlSZW50ZXJORVQxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDExNUswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQhRWsmKAFL1Lj91OwSWTU3H3Oh1fiDWZoBl2qSM%2FZS%2BlpJu9b%2BZXuHgMMF1fr6hmIKQuRiS%2FMAoKSEnwXYdOM8o4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFE%2FoM1S2bFsRSu2jBqebMlvV%2BOCRMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDExNUswCgYIKoZIzj0EAwIDSAAwRQIgVvIkQKS%2BR4gccsuJZ4LS6qr%2BEZKlMLudCJFQxu1p9l4CIQCDw4j5rYUnkLnheOPMaWynz58ROK9vG2AUNesZXH1Q3Q%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 31/10/2022 at 18:34:12