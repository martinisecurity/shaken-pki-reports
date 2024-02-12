# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 625J

Tested At: 12 Feb 24 16:30 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -475 day(s)\
Subject: CN=SHAKEN 625J, OU=SHAKEN, O=Victory Telecom Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/b7343686-5ed8-402c-89a3-8bf1a3d48975/1c896d70e07d655d9b044626a079c628.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC9jCCApygAwIBAgIQdxmCKq0ZVG%2B525PqYKC2xDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTcyMDIyMTdaFw0yMjEwMjQyMDIyMTZaMFMxCzAJBgNVBAYTAlVTMR0wGwYDVQQKExRWaWN0b3J5IFRlbGVjb20gSW5jLjEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNjI1SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHzmaB3NncYawS2Xx9kgkjvNIS1hcyD2%2F7KoOiUOtfdRF4sP7Px0kRBPhgnYmHCMB9l7gPsjq5r1XB4ocxo5eoqjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUyZpGlqEAIillPlVX32CtriWa1VQwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENjI1SjAKBggqhkjOPQQDAgNIADBFAiEAsrdRgpCLn2h7bcJjswDFtBI0AtqX%2F%2FWoFHfL3c%2BoCE0CIHwVZ3q7Fwe5qmrW3n8zPM%2BkBknz8A9wUNG5XhGH%2FsZ3)

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