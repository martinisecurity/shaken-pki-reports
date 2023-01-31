# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 807J

Tested At: 31 Jan 23 17:08 UTC\
Initial Validity Period: 60 day(s)\
Remaining Validity Period: -31 day(s)\
Subject: CN=SHAKEN 807J, OU=SHAKEN, O=SipPhony LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/807J/a97ca418-f914-4add-ad4b-7caf93cf0f2b.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7jCCApSgAwIBAgIQZ84d%2BncLgQhYXxd8q%2FWGoTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMDExMTE0MzBaFw0yMjEyMzExMTE0MjlaMEsxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwxTaXBQaG9ueSBMTEMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDgwN0owWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAS49j9VALOvhZpj%2BsDhSZe6mfy3EesgMC7yVtiSohvM1MZwFvjvxh7qRfgz69H8rcoVOu4TKRoq9lKGjsAcFcsgo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFKiLf%2BHC4hSD6KzyUBC2GyVFT%2B96MB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDgwN0owCgYIKoZIzj0EAwIDSAAwRQIhAOSBtiFjPaeGk7mgfFsQFFJWzgwwM%2BnONZBd1CnwW%2F6tAiBK6YDlJV1PYktqJKyoI2y7pvE%2F7ZFu8C5wowjwc87Oiw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 31 Jan 23 17:51 UTC