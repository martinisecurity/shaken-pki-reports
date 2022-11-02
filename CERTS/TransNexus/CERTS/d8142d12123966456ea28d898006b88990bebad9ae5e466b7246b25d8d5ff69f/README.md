# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 700H

Tested At: 02 Nov 22 15:13 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 4 day(s)\
Subject: CN=SHAKEN 700H, OU=SHAKEN, O=Metro Fibernet LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/22b6cee0-8559-4c73-8092-6eee861c4b49/ae1f22777aac8b17367e1d1b7899f0dd.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC9TCCApqgAwIBAgIQV2Ag0PJ0rB4ZskUN%2F%2Bz9yTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjkyMDE5MzdaFw0yMjExMDUyMDE5MzZaMFExCzAJBgNVBAYTAlVTMRswGQYDVQQKExJNZXRybyBGaWJlcm5ldCBMTEMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDcwMEgwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAR875C8sQXicGJ54tyLy1IPHnTdAxeu2EtvX92wRLpSQe%2B5vAE%2FjRb3QFkkZN9jqlE6ZU%2FEX8r1%2F0x5rx1oqMAbo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFAFQgE2yQ0ZeyjY1QvqIHwfs9UPGMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDcwMEgwCgYIKoZIzj0EAwIDSQAwRgIhAM9dzZ%2Br1N9%2BGqptQ7SIEO%2BRGJczUFBVz2HC4f4WeKClAiEAxTBHGL5YRIZr6IzxEH36oknrn3umbiMYrVVIgZ2aVdA%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 02 Nov 22 15:15 UTC