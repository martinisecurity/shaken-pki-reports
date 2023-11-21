# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 769J

Tested At: 21 Nov 23 18:48 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -422 day(s)\
Subject: CN=SHAKEN 769J, OU=SHAKEN, O=Affiliated Technology Solutions LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/471a46b3-3de4-4aa2-964b-2ff03ddbc8fe/851599d3ef3ba9a16aa77063a3240833.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIDBDCCAqugAwIBAgIQZ88g4bt717PDUzMo7KWpeDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTcyMDE3MzZaFw0yMjA5MjQyMDE3MzVaMGIxCzAJBgNVBAYTAlVTMSwwKgYDVQQKEyNBZmZpbGlhdGVkIFRlY2hub2xvZ3kgU29sdXRpb25zIExMQzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNzY5SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHs%2B6AOlcA73sFIHLjupkHMPRbq6H0QaD12CxM7n9l%2FWYAsiwbZj7jOPhdoBG%2FFylql%2BxMA8Dj774KGveFd36pujggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUay2FHtsrvw7ejkQ1KZJ7CeBYeVAwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENzY5SjAKBggqhkjOPQQDAgNHADBEAiAg7I8E2g4tDCCPBjpRNfnim%2FZ7Emo4EeG%2FfpRGFtQmwwIgLv79pE%2FZpXxj0UVwYpmuL%2BsEQ63lyrrK3zV%2FEXuixW0%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 19:18 UTC