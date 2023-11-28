# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 345J

Tested At: 28 Nov 23 16:02 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -393 day(s)\
Subject: CN=SHAKEN 345J, OU=SHAKEN, O=Ooma Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/345J/31d44cfd-02e8-4a99-830b-c56a58de1835.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6zCCApCgAwIBAgIQVAZVDoxUtd2Z8ozIQnEHbDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDEwNDA0MTZaFw0yMjEwMzEwNDA0MTVaMEcxCzAJBgNVBAYTAlVTMREwDwYDVQQKEwhPb21hIEluYzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMzQ1SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABFiy9OeR8egLCKfIVyPRs1UftGsoCE6RyNxRwdQ9PHKJx9vKHpsNoNVvKOUaiZh1tKGrHd4YO46TlkXrCHN62kejggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUIJY67HjeQJEZlhimvicpeVqGbicwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMzQ1SjAKBggqhkjOPQQDAgNJADBGAiEAnk%2F1XxBSYAOt4HIbL3C6E0O0BhwfdsfzPQa6iWPbDuACIQDqXEd6VJp%2F4Es8R4gCIZ9SzzJZhxh8Y0Ge%2FbLECZAclA%3D%3D)

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


Generated: 28 Nov 23 16:15 UTC