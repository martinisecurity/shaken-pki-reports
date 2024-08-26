# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 7453

Tested At: 26 Aug 24 17:43 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -626 day(s)\
Subject: CN=SHAKEN 7453, O=TPx, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/2a26bd25-8973-44ba-9b94-a8e4716b3888/bb637b228cf588da70ac8c1a8f6d2af4.pem

[View certificate details](https://x509.io/?cert=MIICwzCCAmmgAwIBAgIQSSTrYgXmyMCsgYKEgKA0yjAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjIxMjAxMjAyMTQ4WhcNMjIxMjA4MjAyMTQ3WjAxMQswCQYDVQQGEwJVUzEMMAoGA1UEChMDVFB4MRQwEgYDVQQDEwtTSEFLRU4gNzQ1MzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABKSYuPEJVT02OfTIb9DBhvZIWlnTJdcuinhmhn3xlCAQJdzyKvbm2SLdZ7zMiu64s60jYDLUN3iSY4ee91Dp4xijggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQU1vApQ0oQ6oZ52uW2mIe4VLmzQ9swHwYDVR0jBBgwFoAUMPX18rfhSwsPEHV9SKSQ8T2LKCowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENzQ1MzAKBggqhkjOPQQDAgNIADBFAiEAzBdxJElG3OzruNL%2FLX1JqyX7SxmICjgLX9mDz4AzPm4CIEVBAhXw2kPjjkIG291ZDzZbpnFgAmz1GLIGrQ1%2FcNPE)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 26 Aug 24 18:03 UTC