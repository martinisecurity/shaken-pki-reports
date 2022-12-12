# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 2277

Tested At: 12 Dec 22 23:34 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -50 day(s)\
Subject: CN=SHAKEN 2277, OU=SHAKEN, O=CentraCom, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/0ed28b24-35cd-4fbe-82b2-ae7b4e44f3d3/e096c22db9c707afa3ea6bcd8354432f.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6jCCApGgAwIBAgIQaNjxiziVzaw14Atbk4ZTHDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTYxNjAxMTBaFw0yMjEwMjMxNjAxMDlaMEgxCzAJBgNVBAYTAlVTMRIwEAYDVQQKEwlDZW50cmFDb20xDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDIyNzcwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQJZrDI2mc4EcJkmq5%2Fuju5y1QclW%2FFkJnydpumuWJw2oTJ%2FJx04LyGWQBNnVxZFXSs%2BKjfy%2FFyipkcVnbKGMYJo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFB0L90vCULdvldlolFPphh%2FK1pdzMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDIyNzcwCgYIKoZIzj0EAwIDRwAwRAIgQxEJrIguSfAB0IbBbqW40Egymb7UdgjzwXtD2QmTbxQCIH6jmg%2FW5xeX0M0RASFa3HlSTtbKULDtnwqOpgZAoK7p)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 12 Dec 22 23:45 UTC