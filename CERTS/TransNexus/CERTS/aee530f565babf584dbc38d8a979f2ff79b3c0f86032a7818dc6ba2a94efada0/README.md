# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 120K

Tested At: 26 Aug 24 17:44 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -625 day(s)\
Subject: CN=SHAKEN 120K, O=VC3, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/bcb89748-38e2-45ae-8b54-7efd23525de2/3b48b222d81f28e3be5d312f80c29a3a.pem

[View certificate details](https://x509.io/?cert=MIICwzCCAmmgAwIBAgIQR%2FHrrypnmbpwuG7L0VXCJDAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjIxMjAyMjIyNzQ2WhcNMjIxMjA5MjIyNzQ1WjAxMQswCQYDVQQGEwJVUzEMMAoGA1UEChMDVkMzMRQwEgYDVQQDEwtTSEFLRU4gMTIwSzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABH%2FZnQVqtOxuMHUIQU%2FZylF33jhPX2wwU1I1UzY%2BV4Unz0XHXY1ELMu6eu%2FdA9VVD60QRYWZmyhxOVRbhmiOuNmjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUEc5Ij45kO7vcep7X4%2BBqpAx55%2BQwHwYDVR0jBBgwFoAUMPX18rfhSwsPEHV9SKSQ8T2LKCowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMTIwSzAKBggqhkjOPQQDAgNIADBFAiEArq4BFGN14G50qWhRLxwQTden%2FpROFFT9QWDhYg9kZYgCICpzoj9Ace0Jh0PD5G83IipqW1OTFjhqfnz0rw5NXqRC)

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