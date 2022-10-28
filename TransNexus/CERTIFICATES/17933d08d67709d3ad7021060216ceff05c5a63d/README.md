# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate 17933d08d67709d3ad7021060216ceff05c5a63d
Tested At: 2022-10-28 18:54:27 +0000 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 4 day(s)\
Subject: CN=SHAKEN 983J, OU=SHAKEN, O=ESI, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.clearip.com/59ebb7c1-25bd-4dfc-9794-fcb104b2f66a/ac538bf559c04fd9143ab4aa59589b1c.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC5DCCAougAwIBAgIQVLQ8tK8E%2F8KbMCe5QuciajAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjQyMDE0MDBaFw0yMjEwMzEyMDEzNTlaMEIxCzAJBgNVBAYTAlVTMQwwCgYDVQQKEwNFU0kxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDk4M0owWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATvTsOJVXG0J1I1Gb0bD%2Bgzh8phglTPaXNzcv%2FmrfZA0yWw7cDOc2muf6ChxA0NScDjLe2%2BFPMPD50JcaGPgcwFo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFMPok%2BZ1wbQQ10an%2B5%2BxF82lHSCLMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDk4M0owCgYIKoZIzj0EAwIDRwAwRAIgQb9bO1d6cjYC71vHShhJGFzffavNu1TpqLsJyEECX6sCIDMZERceKo4xIq8qAngv85xIhBZpIz0olpDPyVWdrGL6)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

Generated: 28/10/2022 at 18:55:01