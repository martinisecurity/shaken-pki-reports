# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 606F

Tested At: 02 Nov 22 21:15 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 5 day(s)\
Subject: CN=SHAKEN 606F, OU=SHAKEN, O=Global Data Systems Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/756cb700-f9d2-4a05-850e-c9dfe3e22de4/dccffbc5c2aa54f5f8526ca234e29830.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BjCCAqCgAwIBAgIQdAU4alSbF%2Fh%2BELAqghh1WjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMzAyMTQxNTBaFw0yMjExMDYyMTQxNDlaMFcxCzAJBgNVBAYTAlVTMSEwHwYDVQQKExhHbG9iYWwgRGF0YSBTeXN0ZW1zIEluYy4xDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDYwNkYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQsVzWypgBjrH0AM5lpD9RSr4VtVS8J%2F%2FMI%2FckhpWgDpbFKpUu0eMigA4H8pKlAogxGO00O5P42aZWsPg78H0XQo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFOt4N3e7tQ6oua%2FCzaofiN03BJ1CMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDYwNkYwCgYIKoZIzj0EAwIDSAAwRQIgL3Wj76contGrydp3Rp0zlRRdV34W3KfO7xLzv7kKCJoCIQCgDnxxdmQZ6iN1WWOrBB1zsDlW%2FsEzdYU4WniqQwvU%2BA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 02 Nov 22 21:24 UTC