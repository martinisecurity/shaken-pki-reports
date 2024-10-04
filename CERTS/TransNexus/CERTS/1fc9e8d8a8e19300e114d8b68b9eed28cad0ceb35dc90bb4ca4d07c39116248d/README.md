# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 749J

Tested At: 04 Oct 24 15:45 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -700 day(s)\
Subject: CN=SHAKEN 749J, OU=SHAKEN, O=iTalkGlobal, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/d7bfa171-ff4f-483d-849b-3f987c919d43/d231dd34f8310b359216e546f8265010.pem

[View certificate details](https://x509.io/?cert=MIIC7TCCApOgAwIBAgIQVNLtyqf4fIKQSODuBK6dnzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjcxOTA3NDJaFw0yMjExMDMxOTA3NDFaMEoxCzAJBgNVBAYTAlVTMRQwEgYDVQQKEwtpVGFsa0dsb2JhbDEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNzQ5SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABA8rfxCTaWnDvGAOxob1eh87KH5u6QBxGzCCO07emzPHE0zAFGnhpmIEtU%2FKLF0yDlgRAzMbyUHi5w1M9Bz6sgmjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUNsfBY7q3i9LnVcgrYftdU7%2FU54owHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENzQ5SjAKBggqhkjOPQQDAgNIADBFAiEAqjVf7Vy%2F8k8ynKnhRthwu4lq1tH5W7gGmqHLQ%2BQYrWYCIBDzO6ccfJWG8tv1MyryM2WjlnABSgJVsGpTTa310Tze)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 04 Oct 24 16:29 UTC