# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 0172

Tested At: 04 Oct 24 15:48 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -595 day(s)\
Subject: CN=SHAKEN 0172, O=CallTower, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/0172/4636f98c-4300-41b2-87e8-79ada7138640.pem

[View certificate details](https://x509.io/?cert=MIICyDCCAm%2BgAwIBAgIQWS2DsKTy1ytd3PdjzFAelzAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjMwMjEwMDQ0MTQzWhcNMjMwMjE3MDQ0MTQyWjA3MQswCQYDVQQGEwJVUzESMBAGA1UEChMJQ2FsbFRvd2VyMRQwEgYDVQQDEwtTSEFLRU4gMDE3MjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABFnidE9VPfrVjg0g1kU3AhuD1Ng3q%2BoVQlBFkyA4Q3%2FnZJt%2BrdQCtrmG67H2PezH3jT0WiKHUYaTwGj9uWGh5lOjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQU6wdA%2FfqoHqh2mTTvzfM9Z3ieOVUwHwYDVR0jBBgwFoAUMPX18rfhSwsPEHV9SKSQ8T2LKCowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMDE3MjAKBggqhkjOPQQDAgNHADBEAiBuDTlS0etBbpmowfuaTuh%2FeP2D42vmNm3JqJ8uN8MJJAIgfWXb%2BJS0XaS%2BcP%2BUggxG9kbFQhtUZrLzlREZCnc%2B9TY%3D)

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