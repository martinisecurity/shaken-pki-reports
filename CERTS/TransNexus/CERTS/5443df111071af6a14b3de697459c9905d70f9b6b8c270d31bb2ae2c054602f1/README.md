# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 747J

Tested At: 04 Oct 24 15:42 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -729 day(s)\
Subject: CN=SHAKEN 747J, OU=SHAKEN, O=Magic Apple, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/9e31f6fe-cfd3-49cc-b9fc-22963012a8d7/6a30b4d524fb8c97cfbc2c6131c2b4ef.pem

[View certificate details](https://x509.io/?cert=MIIC7TCCApOgAwIBAgIQX2P4traerpoc2%2FrPVTslrzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MjkwNDUzNTZaFw0yMjEwMDYwNDUzNTVaMEoxCzAJBgNVBAYTAlVTMRQwEgYDVQQKEwtNYWdpYyBBcHBsZTEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNzQ3SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABB0dCTtHcP%2BwnprkDExU2X4wuL0GiTWjCcFDw%2BpsyE221z5rhVXgvgXlcFuN5D%2FkKe9qOlZJTLKggWn6uUeRJMqjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUNECSCrm98eOBMoucuG%2F%2BdiJJJHcwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENzQ3SjAKBggqhkjOPQQDAgNIADBFAiEAu3dYTN8abhF78E27p7zr4uOoY%2BA0b0TW%2BvZGr9MN%2FjQCIBjGiCwA4KakWYbqMamN%2B1jJDWUZJoIa7bwVYnIbRIuS)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 04 Oct 24 16:29 UTC