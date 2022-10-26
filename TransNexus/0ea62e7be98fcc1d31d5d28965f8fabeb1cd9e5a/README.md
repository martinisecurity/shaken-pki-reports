# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate 0ea62e7be98fcc1d31d5d28965f8fabeb1cd9e5a
Tested At: 2022-10-26 20:59:30 +0000 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 6 day(s)\
Subject: CN=SHAKEN 606F, OU=SHAKEN, O=Global Data Systems Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.clearip.com/756cb700-f9d2-4a05-850e-c9dfe3e22de4/42afb34dfd1964ebae01739d9e7120bd.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC%2BzCCAqCgAwIBAgIQTGmwgdYJygtjU08tJFF1hzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjQyMTQxNTFaFw0yMjEwMzEyMTQxNTBaMFcxCzAJBgNVBAYTAlVTMSEwHwYDVQQKExhHbG9iYWwgRGF0YSBTeXN0ZW1zIEluYy4xDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDYwNkYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAR%2F0C5q1um2D3953q9goW5wwO82ba5DyrMEMa8c3vRygwPZPxDaJPNI%2FsmkFsm6ejcjgbDcl4R5D7LL6JyF%2BUrLo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFKkZAWblqnStaf1MFcVqhsaGJDNIMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDYwNkYwCgYIKoZIzj0EAwIDSQAwRgIhAP4eLYZXqSFaVV%2B%2FrcqFcTfg%2Fnu86%2FjgO2G9Ed9bWOL%2FAiEAoGvMVUMxwd4gGPBYcjNlEgYx5xKGbrvcuzhrv%2B%2FojWE%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

Generated: 26/10/2022 at 21:01:13