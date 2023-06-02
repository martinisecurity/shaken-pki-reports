# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 793J

Tested At: 02 Jun 23 01:02 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -207 day(s)\
Subject: CN=SHAKEN 793J, OU=SHAKEN, O=UniTel Voice, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/45c216c2-fbb9-4f33-a3cf-6a64aad19357/25951ca4954a54c8d336007f41644089.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7zCCApSgAwIBAgIQZ4SEv2%2BWtGjkJRBHegu%2BcTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMzAxNTQ0MzFaFw0yMjExMDYxNTQ0MzBaMEsxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwxVbmlUZWwgVm9pY2UxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDc5M0owWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAS%2FPMJxxxuzg%2Fj4lwMYQqvaMkjeR4x8sW8hLciH0TPh6v4y0p8QraH3mGQwQCaSfkhQJaBYTlKO3cTmM4CeeSIso4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFAmtXW8%2FAc5pgBJ4fY9MdXE7L4ZHMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDc5M0owCgYIKoZIzj0EAwIDSQAwRgIhAOktxIa8E2YBjRPqNCn30nt9D3lcmmZ24jitc9F71xGdAiEAqWMflUUwKLLGA7Kfa%2Fk6jGi1nxWtieqKCGD%2BAHQeJ%2Bw%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 02 Jun 23 01:12 UTC