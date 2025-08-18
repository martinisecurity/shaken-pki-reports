# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 0172

Tested At: 18 Aug 25 20:15 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -884 day(s)\
Subject: CN=SHAKEN 0172, O=CallTower, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/0172/f51ee3a4-9314-492c-8bb9-634ead86a3bd.pem

[View certificate details](https://x509.io/?cert=MIICyTCCAm%2BgAwIBAgIQZmuddCfyQQavY%2BHniCvbgzAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjMwMzExMTQzMjQwWhcNMjMwMzE4MTQzMjM5WjA3MQswCQYDVQQGEwJVUzESMBAGA1UEChMJQ2FsbFRvd2VyMRQwEgYDVQQDEwtTSEFLRU4gMDE3MjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABBSyg1G5674BQ0oYgWy1TJoYOcVqG0Hv6DbMAmmb9Ylu1nPAV5Zwt7LPN0QSEMorjeM856oQAValgV%2FM9cdldZ%2BjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQU1gJes%2FOH71qEy71B9IyKFNbn1LUwHwYDVR0jBBgwFoAUMPX18rfhSwsPEHV9SKSQ8T2LKCowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMDE3MjAKBggqhkjOPQQDAgNIADBFAiBUWWPWcQLc0KWFFO1sD%2Fu8lCxZSGyV8HOnRJKavf7EVwIhAPrGtOh8MZiN15EAFxauZDl67EGqGqikmvRoHQIMV2Ko)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 18 Aug 25 21:13 UTC