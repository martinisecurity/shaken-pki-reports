# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 2277

Tested At: 12 Feb 24 16:27 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -477 day(s)\
Subject: CN=SHAKEN 2277, OU=SHAKEN, O=CentraCom, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/0ed28b24-35cd-4fbe-82b2-ae7b4e44f3d3/e096c22db9c707afa3ea6bcd8354432f.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6jCCApGgAwIBAgIQaNjxiziVzaw14Atbk4ZTHDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTYxNjAxMTBaFw0yMjEwMjMxNjAxMDlaMEgxCzAJBgNVBAYTAlVTMRIwEAYDVQQKEwlDZW50cmFDb20xDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDIyNzcwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQJZrDI2mc4EcJkmq5%2Fuju5y1QclW%2FFkJnydpumuWJw2oTJ%2FJx04LyGWQBNnVxZFXSs%2BKjfy%2FFyipkcVnbKGMYJo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFB0L90vCULdvldlolFPphh%2FK1pdzMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDIyNzcwCgYIKoZIzj0EAwIDRwAwRAIgQxEJrIguSfAB0IbBbqW40Egymb7UdgjzwXtD2QmTbxQCIH6jmg%2FW5xeX0M0RASFa3HlSTtbKULDtnwqOpgZAoK7p)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 12 Feb 24 17:02 UTC