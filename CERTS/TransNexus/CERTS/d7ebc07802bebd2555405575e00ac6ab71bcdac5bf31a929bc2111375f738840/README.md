# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 0172

Tested At: 26 Aug 24 18:11 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -543 day(s)\
Subject: CN=SHAKEN 0172, O=CallTower, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/0172/0e8cc042-8f05-4a57-a8ad-1e26da4244d5.pem

[View certificate details](https://x509.io/?cert=MIICyTCCAm%2BgAwIBAgIQS6WqGtptKDCICs1lyamkGzAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjMwMjIyMTkxMjA5WhcNMjMwMzAxMTkxMjA4WjA3MQswCQYDVQQGEwJVUzESMBAGA1UEChMJQ2FsbFRvd2VyMRQwEgYDVQQDEwtTSEFLRU4gMDE3MjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJvjXNa%2Fhj%2FWZcKscTiZo3RhCQHnxkjNW9U%2FrYgrEUixpI3HYmg4vz88wEiDt869yRMKT0biVyKPMEv5VDKJUYKjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUviVyWLjvsQ5MjSyNfeCl%2Fyc9o6kwHwYDVR0jBBgwFoAUMPX18rfhSwsPEHV9SKSQ8T2LKCowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMDE3MjAKBggqhkjOPQQDAgNIADBFAiAc%2BhR54sH5CdeioCees%2FfLaIqLAZKArfe3a6GjS7mw5wIhANxdyFOO1QpiOoJ7G%2F7BUpk6oUJxfmPqJv4QyUarwFOC)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 26 Aug 24 18:49 UTC