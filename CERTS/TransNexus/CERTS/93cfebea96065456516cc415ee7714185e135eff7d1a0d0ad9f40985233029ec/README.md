# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 157C

Tested At: 31 Jan 23 21:39 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -107 day(s)\
Subject: CN=SHAKEN 157C, OU=SHAKEN, O=DigitalSpeed Communications, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/82d97ecb-9b38-4926-9062-8e21b8986930/28ec7e97780e7ae26bbe623cc083537c.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2FDCCAqOgAwIBAgIQdaI53SlqzpVwRdlOOTskcDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDkyMDUwNDdaFw0yMjEwMTYyMDUwNDZaMFoxCzAJBgNVBAYTAlVTMSQwIgYDVQQKExtEaWdpdGFsU3BlZWQgQ29tbXVuaWNhdGlvbnMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDE1N0MwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQiIM7JktGOqw%2B9g0V23rXN%2BMYme7q8FV%2BjUeWU8Sti9tbSJB9d82Vgs7FyLA0PzWhTvPl%2BnyTnOQDhcK9b6KbVo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFNA6BTX0lNpg34GFY4ZgDCNEynKBMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDE1N0MwCgYIKoZIzj0EAwIDRwAwRAIgOo%2Fwl6QBJYuPa%2F0GTzkD81hMMsH8jAURvHPqF3TdQGwCIAVAvQiJNJpRYcPPVyqNQtsufJ3Tfw92Wysyz1VwYs46)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 31 Jan 23 21:50 UTC