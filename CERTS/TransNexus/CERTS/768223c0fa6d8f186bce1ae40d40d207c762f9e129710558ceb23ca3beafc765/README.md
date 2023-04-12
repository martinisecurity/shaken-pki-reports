# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 700H

Tested At: 12 Apr 23 01:02 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -142 day(s)\
Subject: CN=SHAKEN 700H, OU=SHAKEN, O=Metro Fibernet LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/22b6cee0-8559-4c73-8092-6eee861c4b49/839958f07be8c8bd339603497b5802ad.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC9DCCApqgAwIBAgIQS0lRyH60yWlElMm4LlZDbDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMTMyMDIwMjlaFw0yMjExMjAyMDIwMjhaMFExCzAJBgNVBAYTAlVTMRswGQYDVQQKExJNZXRybyBGaWJlcm5ldCBMTEMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDcwMEgwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARy%2Bkzl2bluuxD57sEAtfmautrMoNAorVFS%2B3IYxzYSBLVMk7i7JNxq7kgMXU0OH6QStfv%2Bn%2FFT7LB%2FH3qHV3URo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFD9n0mZjOfOYDOVjtQvqi2xvIcccMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDcwMEgwCgYIKoZIzj0EAwIDSAAwRQIhALjIz1ka7DahHVmCNVsVcd%2FUEFfW3g7oFDsjrvKVfQ8GAiAhFfHuJiQI%2FtNT8Mzb1IAMYm4T5CDmCgs3t9pgFG%2Fu7Q%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 12 Apr 23 01:46 UTC