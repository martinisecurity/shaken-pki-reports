# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 738J

Tested At: 26 Aug 24 17:45 UTC\
Initial Validity Period: 9 day(s)\
Remaining Validity Period: -599 day(s)\
Subject: CN=SHAKEN 738J, O=SkySwitch, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/738J/d69429ed-670a-46fb-8d2b-daca7e81bfaf.pem

[View certificate details](https://x509.io/?cert=MIICyTCCAm%2BgAwIBAgIQbGmA3goV7YG0o2AACndOHzAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjIxMjI2MjAzMDA0WhcNMjMwMTA0MjAzMDAzWjA3MQswCQYDVQQGEwJVUzESMBAGA1UEChMJU2t5U3dpdGNoMRQwEgYDVQQDEwtTSEFLRU4gNzM4SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABNE5uitNn9HNkucoPF%2F9OZB7LBGeepoaBkDZz296aaFUGIHIhHYHKQ03Z3lFKlT9XWtJkZsHVzodAcUX5vlg7kKjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQU81Nhmn5ZDSMMxwQSecWDy3HKohAwHwYDVR0jBBgwFoAUMPX18rfhSwsPEHV9SKSQ8T2LKCowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENzM4SjAKBggqhkjOPQQDAgNIADBFAiEA5CRKR66BPtE3vS%2BUqeWTIG5Sv5vZ4dNNh%2FzaSPtToukCIExlbfuHXu%2F2KK4QaFeRByPA9KLujBzBYthDYVAWvkFU)

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