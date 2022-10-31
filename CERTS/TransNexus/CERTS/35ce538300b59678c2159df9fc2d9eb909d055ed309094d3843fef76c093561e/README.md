# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 2277

Tested At: 31 Oct 22 19:20 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 7 day(s)\
Subject: CN=SHAKEN 2277, OU=SHAKEN, O=CentraCom, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/0ed28b24-35cd-4fbe-82b2-ae7b4e44f3d3/becf02556a3fe0add34e22f52442d2cf.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC7DCCApGgAwIBAgIQTTdCx5H7bFI76Srz6%2F3kPDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMzExNjAzMTRaFw0yMjExMDcxNjAzMTNaMEgxCzAJBgNVBAYTAlVTMRIwEAYDVQQKEwlDZW50cmFDb20xDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDIyNzcwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATTB5JQUxwtZTjjhxCidYMWa8BUNiK9Pj3GlcUibwRm7ZjHi7Jh8cUWj7hx7s%2FRod%2FALonk7HihTqIO%2BsDN9aHPo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFBIHXtvsUVoRySrI4fCte6QP7fL3MB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDIyNzcwCgYIKoZIzj0EAwIDSQAwRgIhAPTOCgtwhxVuhQ7pmJgIBHGnSHc6BQrjD4TuTMrDV9YEAiEAlocvWHpu8OOXz9driBX6WoBvVVx09%2F1vk3giqZilTuU%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 31/10/2022 at 19:21:49