# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 0172

Tested At: 22 Aug 24 15:27 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -602 day(s)\
Subject: CN=SHAKEN 0172, O=CallTower, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/0172/6124abb6-773c-4564-b0c7-ecc3b36b60f9.pem

[View certificate details](https://x509.io/?cert=MIICyTCCAm%2BgAwIBAgIQZj3fiAHdtby9bnZ%2FPH%2Fi8TAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjIxMjIxMTgzMzA3WhcNMjIxMjI4MTgzMzA2WjA3MQswCQYDVQQGEwJVUzESMBAGA1UEChMJQ2FsbFRvd2VyMRQwEgYDVQQDEwtTSEFLRU4gMDE3MjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABK4I%2BTOHPBXo8jjdqJRdo83L8e7V5AQNz93R4YrGsXYgR1SPBGBavH2TtSDf%2BmKKyEjPyATUrThGo9heJXazOZejggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUDymbc4dYj3avXH3XYzWL9la1gzYwHwYDVR0jBBgwFoAUMPX18rfhSwsPEHV9SKSQ8T2LKCowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMDE3MjAKBggqhkjOPQQDAgNIADBFAiEAv5b8wbd2ixDCU4xSQ7uf47pObYyaiBs1FPC2yEqhZD4CICMlaST%2BCd2fKtqzI5ENqrV6USI4xC%2FYk8WDJ%2BXxpmox)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 22 Aug 24 16:06 UTC